// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

package commands

import (
	//"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/nlamirault/go-onlinelabs/api"
)

var commandListServers = cli.Command{
	Name:        "listServers",
	Usage:       "List all servers associate with your account",
	Description: ``,
	Action:      doListServers,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

var commandGetServer = cli.Command{
	Name:        "getServer",
	Usage:       "Retrieve a server",
	Description: ``,
	Action:      doGetServer,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "serverid",
			Usage: "Server unique identifier",
			Value: "",
		},
	},
}

var commandDeleteServer = cli.Command{
	Name:        "deleteServer",
	Usage:       "Delete a server",
	Description: ``,
	Action:      doDeleteServer,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "serverid",
			Usage: "Server unique identifier",
			Value: "",
		},
	},
}

var commandActionServer = cli.Command{
	Name:        "actionServer",
	Usage:       "Execute an action on a server",
	Description: "Execute an action on a server",
	Action:      doActionServer,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "serverid",
			Usage: "Server unique identifier",
			Value: "",
		},
		cli.StringFlag{
			Name:  "action",
			Usage: "the action to perform",
			Value: "",
		},
	},
}

func doListServers(c *cli.Context) {
	log.Infof("List servers")
	client := getClient(c)
	b, err := client.GetServers()
	if err != nil {
		log.Errorf("Retrieving servers %v", err)
		return
	}
	response, err := api.GetServersFromJSON(b)
	if err != nil {
		log.Errorf("Reading servers %v", err)
		return
	}
	log.Infof("Servers: ")
	for _, server := range response.Servers {
		log.Infof("----------------------------------------------")
		log.Infof("Id   : %s", server.ID)
		log.Infof("Name : %s", server.Name)
		log.Infof("Date : %s", server.ModificationDate)
		log.Infof("IP   : %s", server.PublicIP.Address)
		log.Infof("Tags : %s", server.Tags)
	}
}

func doGetServer(c *cli.Context) {
	log.Infof("Getting server %s", c.String("serverid"))
	client := getClient(c)
	b, err := client.GetServer(c.String("serverid"))
	if err != nil {
		log.Errorf("Retrieving server: %v", err)
	}
	response, err := api.GetServerFromJSON(b)
	if err != nil {
		log.Errorf("Failed response %v", err)
		return
	}
	log.Infof("Server: ")
	log.Infof("Id    : %s", response.Server.ID)
	log.Infof("Name  : %s", response.Server.Name)
	log.Infof("State : %s", response.Server.State)
	log.Infof("Date  : %s", response.Server.ModificationDate)
	log.Infof("IP    : %s", response.Server.PublicIP.Address)
	log.Infof("Tags  : %s", response.Server.Tags)
}

func doDeleteServer(c *cli.Context) {
	log.Infof("Remove server %s", c.String("serverid"))
	client := getClient(c)
	b, err := client.DeleteServer(c.String("serverid"))
	if err != nil {
		log.Errorf("Retrieving server: %v", err)
	}
	log.Infof("Server deleted: %s", string(b))
}

func doActionServer(c *cli.Context) {
	log.Infof("Perform action %s on server %s", c.String("action"), c.String("serverid"))
	client := getClient(c)
	b, err := client.PerformServerAction(
		c.String("serverid"), c.String("action"))
	if err != nil {
		log.Errorf("Failed: %v", err)
		return
	}
	response, err := api.GetTaskFromJSON(b)
	if err != nil {
		log.Errorf("Failed response %v", err)
		return
	}
	log.Infof("Task: ")
	log.Infof("Id          : %s", response.Task.ID)
	log.Infof("Status      : %s", response.Task.Status)
	log.Infof("Description : %s", response.Task.Description)

}
