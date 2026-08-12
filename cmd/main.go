package main

import (
  "fmt"
  "github.com/s0h311/project-matrix/pkg/sandbox"
)

func main() {
  // LIST
  sandboxes, _ := sandbox.List()

  fmt.Printf("%s\n\n", sandboxes)

  // CREATE
  err := sandbox.Create(&sandbox.CreateOptions{
    Name:       "testing-1",
    Agent:      "claude",
    Workspaces: []string{"."},
  })

  if err != nil {
    fmt.Printf("%s\n\n", err)
  }

  // LIST
  sandboxes, _ = sandbox.List()
  fmt.Printf("%s\n\n", sandboxes)

  // EXEC
  workDir := "/home"

  execRes1, _ := sandbox.Exec(&sandbox.ExecOptions{
    Sandbox:          "testing-1",
    Cmd:              "env",
    Env:              &[]string{"HELLO=was geht", "MOIN=okay"},
    WorkingDirectory: &workDir,
  })

  fmt.Printf("%s\n\n", execRes1)

  execRes2, _ := sandbox.Exec(&sandbox.ExecOptions{
    Sandbox:          "testing-1",
    Cmd:              "pwd",
    WorkingDirectory: &workDir,
  })

  fmt.Printf("%s\n\n", execRes2)

  // RM
  err = sandbox.Rm(&[]string{"testing-1"})

  if err != nil {
    fmt.Printf("%s\n\n", err)
  }

  // LIST
  sandboxes, _ = sandbox.List()
  fmt.Printf("%s\n\n", sandboxes)
}
