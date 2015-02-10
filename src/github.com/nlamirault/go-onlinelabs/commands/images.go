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
	//"log"
	//"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/nlamirault/go-onlinelabs/api"
	// "github.com/nlamirault/go-onlinelabs/logging"
)

var commandListImages = cli.Command{
	Name:        "listImages",
	Usage:       "List availables images",
	Description: ``,
	Action:      doListImages,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

var commandGetImage = cli.Command{
	Name:        "getImage",
	Usage:       "Retrieve a image",
	Description: ``,
	Action:      doGetImage,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "imageid",
			Usage: "Image unique identifier",
			Value: "",
		},
	},
}

func doGetImage(c *cli.Context) {
	log.Infof("Getting image %s", c.String("imageid"))
	client := getClient(c)
	b, err := client.GetImage(c.String("imageid"))
	if err != nil {
		log.Errorf("Retrieving image: %v", err)
	}
	response, err := api.GetImageFromJSON(b)
	if err != nil {
		log.Errorf("Failed response %v", err)
		return
	}
	log.Infof("Image: ")
	response.Image.Display()
}

func doListImages(c *cli.Context) {
	log.Infof("List images")
	client := getClient(c)
	b, err := client.GetImages()
	if err != nil {
		log.Errorf("Retrieving imagess %v", err)
		return
	}
	response, err := api.GetImagesFromJSON(b)
	if err != nil {
		log.Errorf("Reading images %v", err)
		return
	}
	log.Infof("Images: ")
	for _, image := range response.Images {
		log.Infof("----------------------------------------------")
		image.Display()
	}
}
