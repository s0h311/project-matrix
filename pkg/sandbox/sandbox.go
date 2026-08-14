package sandbox

import (
  "encoding/json"
  "errors"
  "os"
  "os/exec"
  "regexp"
  "strings"
)

// Service github | anthropic | openai | google | ...
type Service = string

type Sandbox struct {
  Id   string `json:"id"`
  Name string `json:"name"`
  // claude | codex | gemini
  Agent string `json:"agent"`
  // stopped | running
  Status     string   `json:"status"`
  Workspaces []string `json:"workspaces"`
}

type lsResult struct {
  Sandboxes []Sandbox `json:"sandboxes"`
}

type CreateOptions struct {
  Name       string
  Agent      string
  Template   *string
  Workspaces []string
}

type ExecOptions struct {
  Sandbox          string
  Cmd              string
  Args             *[]string
  WorkingDirectory *string
  Env              *[]string
  EnvFiles         *[]string
}

type Secret struct {
  // global
  Scope string
  // service
  Type string
  Name Service
  // stored | oauth_configured
  Secret string
}

type SecretSetOptions struct {
  Service Service
  Token   string
  Force   *bool
  Sandbox *string
}

type SecretRmOptions struct {
  Service Service
  Sandbox *string
}

type PolicyLsOptions struct {
  // allow | deny
  Decision        *string
  IncludeInactive *bool
  // all | network | filesystem
  Type *string
}

type Policy struct {
  Id           string   `json:"id"`
  Name         string   `json:"name"`
  PolicyId     string   `json:"policy_id"`
  Scope        string   `json:"scope"`
  AppliesTo    string   `json:"applies_to"`
  ResourceType string   `json:"resource_type"`
  Decision     string   `json:"decision"`
  Resources    []string `json:"Resources"`
  Origin       string   `json:"origin"`
  Layer        string   `json:"layer"`
  Status       string   `json:"status"`
  Editable     bool     `json:"editable"`
}

type policyLsResult struct {
  Rules []Policy
}

func Ls() ([]Sandbox, error) {
  cmd := []string{"ls", "--json"}

  result, err := execSbxCmd(&cmd)

  if err != nil {
    return nil, err
  }

  var lsResult lsResult

  err = json.Unmarshal(result, &lsResult)

  if err != nil {
    return nil, err
  }

  return lsResult.Sandboxes, nil
}

func Create(createOptions *CreateOptions) error {
  cmd := []string{"create", "--name", createOptions.Name}

  if createOptions.Template != nil {
    cmd = append(cmd, "--template", *createOptions.Template)
  }

  cmd = append(cmd, createOptions.Agent)
  cmd = append(cmd, createOptions.Workspaces...)

  _, err := execSbxCmd(&cmd)

  if err != nil {
    return err
  }

  return nil
}

func Rm(sandboxes *[]string) error {
  cmd := append([]string{"rm", "--force"}, *sandboxes...)

  _, err := execSbxCmd(&cmd)

  if err != nil {
    return err
  }

  return nil
}

func Exec(execOptions *ExecOptions) ([]byte, error) {
  cmd := []string{"exec"}

  if execOptions.Env != nil {
    for _, env := range *execOptions.Env {
      cmd = append(cmd, "--env", env)
    }
  }

  if execOptions.EnvFiles != nil {
    for _, envFile := range *execOptions.EnvFiles {
      cmd = append(cmd, "--env-file", envFile)
    }
  }

  if execOptions.WorkingDirectory != nil {
    cmd = append(cmd, "--workdir", *execOptions.WorkingDirectory)
  }

  cmd = append(cmd, "--", execOptions.Sandbox, execOptions.Cmd)

  if execOptions.Args != nil {
    cmd = append(cmd, *execOptions.Args...)
  }

  return execSbxCmd(&cmd)
}

func SecretLs() ([]Secret, error) {
  cmd := []string{"secret", "ls"}

  rawResult, err := execSbxCmd(&cmd)

  if err != nil {
    return nil, err
  }

  result := strings.Trim(string(rawResult), "\n")

  result = regexp.MustCompile(`\s{2,}|\n`).
    ReplaceAllString(result, ";")

  result = regexp.MustCompile(`[()]`).
    ReplaceAllString(result, "")

  result = regexp.MustCompile(`\s`).
    ReplaceAllString(result, "_")

  fields := strings.Split(result, ";")

  if len(fields)%4 != 0 {
    return nil, errors.New("'sbx secret ls' output unsupported format")
  }

  var secrets []Secret

  for i := 4; i < len(fields); i = i + 4 {
    secrets = append(secrets, Secret{
      Scope:  fields[i],
      Type:   fields[i+1],
      Name:   fields[i+2],
      Secret: fields[i+3],
    })
  }

  return secrets, nil
}

func SecretSet(secretSetOptions *SecretSetOptions) error {
  cmd := []string{"secret", "set", secretSetOptions.Service, "--token", secretSetOptions.Token}

  if secretSetOptions.Force != nil && *secretSetOptions.Force == true {
    cmd = append(cmd, "--force")
  }

  if secretSetOptions.Sandbox != nil {
    cmd = append(cmd, "--sandbox", *secretSetOptions.Sandbox)
  }

  _, err := execSbxCmd(&cmd)

  return err
}

func SecretRm(secretRmOptions *SecretRmOptions) error {
  cmd := []string{"secret", "rm", secretRmOptions.Service, "--force"}

  if secretRmOptions.Sandbox != nil {
    cmd = append(cmd, "--sandbox", *secretRmOptions.Sandbox)
  }

  _, err := execSbxCmd(&cmd)

  return err
}

func PolicyLs(policyLsOptions *PolicyLsOptions) ([]Policy, error) {
  cmd := []string{"policy", "ls", "--json"}

  if policyLsOptions.Decision != nil {
    cmd = append(cmd, "--decision", *policyLsOptions.Decision)
  }

  if policyLsOptions.IncludeInactive != nil && *policyLsOptions.IncludeInactive {
    cmd = append(cmd, "--include-inactive")
  }

  if policyLsOptions.Type != nil {
    cmd = append(cmd, "--type", *policyLsOptions.Type)
  }

  rawResult, err := execSbxCmd(&cmd)

  if err != nil {
    return nil, err
  }

  var result policyLsResult

  err = json.Unmarshal(rawResult, &result)

  if err != nil {
    return nil, err
  }

  return result.Rules, nil
}

func PolicyAllowNetwork() {}
func PolicyCheckNetwork() {}
func PolicyDenyNetwork()  {}
func PolicyRmNetwork()    {}

func execSbxCmd(cmd *[]string) ([]byte, error) {
  _cmd := exec.Command("sbx", *cmd...)
  _cmd.Stderr = os.Stderr
  result, err := _cmd.Output()

  if err != nil {
    return nil, err
  }

  return result, nil
}
