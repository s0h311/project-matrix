package sandbox

import (
  "encoding/json"
  "os"
  "os/exec"
)

type Sandbox struct {
  Id   string `json:"id"`
  Name string `json:"name"`
  // claude | codex | gemini
  Agent string `json:"agent"`
  // stopped | running
  Status     string   `json:"status"`
  Workspaces []string `json:"workspaces"`
}

type LsResult struct {
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

func List() ([]Sandbox, error) {
  cmd := []string{"ls", "--json"}

  result, err := execSbxCmd(cmd)

  if err != nil {
    return nil, err
  }

  var lsResult LsResult

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

  _, err := execSbxCmd(cmd)

  if err != nil {
    return err
  }

  return nil
}

func Rm(sandboxes *[]string) error {
  cmd := append([]string{"rm", "--force"}, *sandboxes...)

  _, err := execSbxCmd(cmd)

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

  return execSbxCmd(cmd)
}

func execSbxCmd(cmd []string) ([]byte, error) {
  _cmd := exec.Command("sbx", cmd...)
  _cmd.Stderr = os.Stderr
  result, err := _cmd.Output()

  if err != nil {
    return nil, err
  }

  return result, nil
}
