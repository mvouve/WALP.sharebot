package main

import (
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	//  "fmt";
	//  "log"
)

type sftpConn struct {
	client   *sftp.Client
	ssh      *ssh.Client
	username string
	password string
	addr     string
}

// connect connects to the sftp server based on user credentials
func (s *sftpConn) connect() error {
	ssh, err := sshConnect(s.addr, s.username, s.password)
	if err != nil {
		return err
	}

	c, err := sftp.NewClient(ssh)
	if err != nil {
		return err
	}

	s.ssh = ssh
	s.client = c

	return nil
}

func (s *sftpConn) readDir(dir string) ([]os.FileInfo, error) {
	return s.readDir(dir)
}

// sshConnect initialises the SSH connection to the remote server.
func sshConnect(addr string, name string, pass string) (*ssh.Client, error) {
	// configure connection
	config := &ssh.ClientConfig{
		User: name,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		}}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *sftpConn) close() {
	s.client.Close()
	s.ssh.Close()
}

func (s *sftpConn) get(path string) (*os.File, error) {
	return os.Open(path)
}
