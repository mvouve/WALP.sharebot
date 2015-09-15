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
  "fmt"
)

func main(){
  if len(os.Args) == 4 {
    connect(os.Args[1], os.Args[2], os.Args[3])
  } else {
    fmt.Println("Invalid: Correct usage: " + os.Args[0] + " [address:portno] [username] [password]")
  }
}
