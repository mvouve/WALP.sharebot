/*
THIS FILE IS PART OF SHAREBOT

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/**
This program is designed to interface with Slack and ShareOut at
BCIT, Although it could be used to interface with a number of different file
shares, or other chat interfaces.

@author Marc Vouve (mvouve@gmail.com)
@date Sept 15, 2015
*/

package main

import (
  "os";
  "fmt";
  "github.com/pkg/sftp";
)

// file system to move files to.
type writeFileSystem interface
{
  // connect initialises a connection to a some sort of file share
  connect(addr string, name string, auth string)
  // find returns true if a file matches the description provided
  find(fname string, fsize int, fdate int) bool
  // addFile adds a file to the file system based on name

  // close closes the connection and cleans up the fs
}

// file system to connect to
type readFileSystem interface
{
  connect(addr string, name string, auth string) error
  // readDir reads the files from that connection and returns os.FileInfo
  readDir(dir string) (os.FileInfo, error)
  // getFile gets a file from the file system based off file name
  getFile(name string)

  // close closes and cleans up the connection
}
// notification interface
type notifier interface
{
  connect()
  alert(alert string)
}

// main starts the program.
//
// if a contributer would wish to create a new readFileSystem/writeFileSystem/notifier
// they could add it here. where the construct is initialised.
func main(){
  if len(os.Args) != 5 {  // req 4 args addr, user, pass
    fmt.Println("Invalid: Correct usage: " + os.Args[0] + " [address:portno] [username] [password]")
  }

  readFS:= &sftpConn{username: os.Arg[1], password: os.Arg[2], addr: os.Arg[3]}
  err := readFileSystem.connect()
  if err != nil
  {
    log.Fatalf("Error faild to initialise read file system: %v", err.Error())
  }
  defer readFileSystem.close()

  writeFS := &googleDrive{}
  err := writeFS.connect()
  if err != nil {
    log.Fatalf("Error faild to initialise write file system: %v", err.Error())
  }
  defer writeFS.close()

  notifier := &slack{}
  err := notifier.connect()
  if err != nil {
    log.Fatalf("Error faild to initialise notifier: %v", err.Error())
  }
  defer notifier.close()

  files := readFS.readDir(os.Args[5])
  newfiles = new(List)

  for _ , i :=  range files {
    if !writeFS.exists(i.Name()) {
      writeFS.add(readFS.get(i.Name()))
      newfiles.pushBack(i.Name())
    }
  }
}
