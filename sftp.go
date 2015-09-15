package main

import (
  "github.com/pkg/sftp";
  "golang.org/x/crypto/ssh";
  "fmt";
  "log"
)

struct sftpConn {
  client *sftp.Client
  username string
  password string
  addr string
}

// connect connects to the sftp server based on user credentials
func (s *sftpConn) connect() error {
  ssh, err := sshConnect(s.addr, s.name, s.auth)
  if err != nil {
    return err
  }

  s.client, err := sftp.NewClient(ssh)
  if err != nil {
    return err
  }
}

func (s *sftpConn) readDir(dir string) os.FileInfo, error
{
  return s.ReadDir(dir);
}

// sshConnect initialises the SSH connection to the remote server.
func sshConnect(addr string, name string, pass string)(*ssh.Client, error){
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
