/*
  Copyright 2015 W. Max Lees

  This file is part of jarvisos.

  Jarvisos is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  Jarvisos is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with jarvisos.  If not, see <http://www.gnu.org/licenses/>.

  File: demoapp.go
  Author: W. Max Lees <max.lees@gmail.com>
  Date: 06.14.2015
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jarvisos/app"
	"strconv"
)

type DemoApp struct {
	name string
	info app.Info
}

func (d *DemoApp) Call(call string, result *[]byte) error {
	str := []byte("{ \"result\": \"Success\" }")
	result = &str

	return nil
}

func (d *DemoApp) Who(full bool, result *[]byte) error {
	if full {
		info, err := json.Marshal(d.info)
		if err != nil {
			fmt.Printf("Error marshalling app information: %v\n", err)
			return err
		}
		result = &info
	} else {
		info, err := json.Marshal(d.info.Port)
		if err != nil {
			fmt.Printf("Error marshalling app address: %v\n", err)
			return err
		}
		result = &info
	}

	return nil
}

func main() {
	demoApp := DemoApp{}

	// Get the command line flags
	port := flag.Int("p", 0, "Port for the program to listen on")
	flag.Parse()

	// Maybe generate this with a command line arg
	demoApp.info.Port = strconv.Itoa(*port)

	err := app.Run(&demoApp, "localhost:"+demoApp.info.Port)
	if err != nil {
		fmt.Printf("Error running application: %v\n", err)
	}
}
