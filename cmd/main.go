package main

import (
  "fmt"
  "github.com/s0h311/project-matrix/pkg/sandbox"
)

func main() {
  //// LIST
  //sandboxes, _ := sandbox.Ls()
  //
  //fmt.Printf("%s\n\n", sandboxes)
  //
  //// CREATE
  //err := sandbox.Create(&sandbox.CreateOptions{
  //  Name:       "testing-1",
  //  Agent:      "claude",
  //  Workspaces: []string{"."},
  //})
  //
  //if err != nil {
  //  fmt.Printf("%s\n\n", err)
  //}
  //
  //// LIST
  //sandboxes, _ = sandbox.Ls()
  //fmt.Printf("%s\n\n", sandboxes)
  //
  //// EXEC
  //workDir := "/home"
  //
  //execRes1, _ := sandbox.Exec(&sandbox.ExecOptions{
  //  Sandbox:          "testing-1",
  //  Cmd:              "env",
  //  Env:              &[]string{"HELLO=was geht", "MOIN=okay"},
  //  WorkingDirectory: &workDir,
  //})
  //
  //fmt.Printf("%s\n\n", execRes1)
  //
  //execRes2, _ := sandbox.Exec(&sandbox.ExecOptions{
  //  Sandbox:          "testing-1",
  //  Cmd:              "pwd",
  //  WorkingDirectory: &workDir,
  //})
  //
  //fmt.Printf("%s\n\n", execRes2)
  //
  //// RM
  //err = sandbox.Rm(&[]string{"testing-1"})
  //
  //if err != nil {
  //  fmt.Printf("%s\n\n", err)
  //}
  //
  //// LIST
  //sandboxes, _ = sandbox.Ls()
  //fmt.Printf("%s\n\n", sandboxes)

  sandbox1 := "matrix-tabley"
  force := true

  err1 := sandbox.SecretSet(&sandbox.SecretSetOptions{
    Service: "openai",
    Token:   "abc",
    Sandbox: &sandbox1,
    Force:   &force,
  })

  if err1 != nil {
    panic(err1)
  }

  secrets, err := sandbox.SecretLs()

  if err != nil {
    panic(err)
  }

  fmt.Println(secrets)

  err2 := sandbox.SecretRm(&sandbox.SecretRmOptions{
    Service: "openai",
    Sandbox: &sandbox1,
  })

  if err2 != nil {
    panic(err2)
  }

  secrets2, err3 := sandbox.SecretLs()

  if err3 != nil {
    panic(err3)
  }

  fmt.Println(secrets2)
}
