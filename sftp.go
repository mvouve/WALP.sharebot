package main

import (
  "github.com/pkg/sftp";
  "golang.org/x/crypto/ssh";
  "fmt";
  "log"
)

func connect(addr string, name string, pass string)(*sftp.Client){
  ssh := ssh_connect(addr, name, pass)
  conn := new_sftp(ssh)
  fmt.Printf(addr + name + pass)

  return conn
}

func ssh_connect(addr string, name string, pass string)(*ssh.Client){
  // configure connection
  config := &ssh.ClientConfig{
    User: name,
    Auth: []ssh.AuthMethod{
    ssh.Password(pass),
    }}

    client, err := ssh.Dial("tcp", addr, config)
    if err != nil {
      log.Fatal("Failed to Dial: " + err.Error())
    }

    return client
}

func new_sftp(ssh *ssh.Client)(*sftp.Client){
  ret, err := sftp.NewClient(ssh)
  if err != nil{
    log.Fatal("Failed to initialise SFTP error: " + err.Error())
  }

  return ret
}
