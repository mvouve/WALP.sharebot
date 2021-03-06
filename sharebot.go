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
	"container/list"
	"log"
	"os"
)

// file system to move files to.
type writeFileSystem interface {
	// connect initialises a connection to a some sort of file share
	connect()
	// find returns true if a file matches the description provided
	exists(fname string, fsize int, fdate int) bool
	// addFile adds a file to the file system based on name

	// close closes the connection and cleans up the fs
	close()
}

// file system to connect to
type readFileSystem interface {
	connect() error
	// readDir reads the files from that connection and returns os.FileInfo
	//readDir(dir string) (os.FileInfo, error)
	// getFile gets a file from the file system based off file name
	//get(path string) (*os.File, error)
	// close closes and cleans up the connection
	close()
}

// notification interface
type notifier interface {
	connect()
	alert(alert string)
}

// main starts the program.
//
// if a contributer would wish to create a new readFileSystem/writeFileSystem/notifier
// they could add it here. where the construct is initialised.
func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Invalid: Correct usage: " + os.Args[0] + " [address:portno] [username] [password]")
	}

	readFS := &sftpConn{username: os.Args[1], password: os.Args[2], addr: os.Args[3]}
	err := readFS.connect()
	if err != nil {
		log.Fatalf("Error faild to initialise read file system: %v", err.Error())
	}
	defer readFS.close()

	writeFS := &googleDrive{}
	err = writeFS.connect()
	if err != nil {
		log.Fatalf("Error faild to initialise write file system: %v", err.Error())
	}
	defer writeFS.close()

	notifier := &slack{}
	err = notifier.connect()
	if err != nil {
		log.Fatalf("Error faild to initialise notifier: %v", err.Error())
	}
	defer notifier.close()

	files, _ := readFS.readDir(os.Args[5])
	newfiles := list.New()

	for _, i := range files {
		if !writeFS.exists(i.Name(), i.Size()) {
			f, err := readFS.get(i.Name())
			if err != nil {
				log.Println("File disappeared between match and fetch")
			}
			writeFS.add(f)
			newfiles.PushBack(i.Name())
		}
	}
}
