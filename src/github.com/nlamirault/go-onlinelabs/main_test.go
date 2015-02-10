// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	//"fmt"
	"strings"
	"testing"

	"github.com/codegangsta/cli"
)

func checkCommand(t *testing.T, name string, command cli.Command) {
	if command.Name != name {
		t.Errorf("Invalid command name: %s", command.Name)
	}
}

func TestCLICommands(t *testing.T) {
	app := makeApp()
	if len(app.Commands) != 13 {
		t.Errorf("Invalid CLI number of commands")
	}
	// for _, command := range app.Commands {
	// 	fmt.Printf("command : %v", command)
	// }
	checkCommand(t, "listImages", app.Commands[0])
	checkCommand(t, "listServers", app.Commands[1])
	checkCommand(t, "getServer", app.Commands[2])
	checkCommand(t, "deleteServer", app.Commands[3])
	checkCommand(t, "actionServer", app.Commands[4])
	checkCommand(t, "getUser", app.Commands[5])
	checkCommand(t, "getOrganizations", app.Commands[6])
	checkCommand(t, "getTokens", app.Commands[7])
	checkCommand(t, "getToken", app.Commands[8])
	checkCommand(t, "listVolumes", app.Commands[9])
	checkCommand(t, "getVolume", app.Commands[10])
	checkCommand(t, "deleteVolume", app.Commands[11])
	checkCommand(t, "createVolume", app.Commands[12])
}

func checkGlobalArgument(flags []cli.Flag, name string) int {
	for i, flag := range flags {
		//fmt.Printf("Flag: %v\n", flag.String())
		if strings.HasPrefix(flag.String(), name) {
			return i
		}
	}
	return -1
}

func TestOnlinelabsUserIDArgument(t *testing.T) {
	app := makeApp()
	if checkGlobalArgument(app.Flags, "--onlinelabs-userid") == -1 {
		t.Errorf("no userid flag")
	}
}

func TestOnlinelabsTokenArgument(t *testing.T) {
	app := makeApp()
	if checkGlobalArgument(app.Flags, "--onlinelabs-token") == -1 {
		t.Errorf("No token flag")
	}
}

func TestOnlinelabsOrganizationArgument(t *testing.T) {
	app := makeApp()
	if checkGlobalArgument(app.Flags, "--onlinelabs-organization") == -1 {
		t.Errorf("No organization flag")
	}
}
