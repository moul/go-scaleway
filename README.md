# go-onlinelabs

[![License GPL 3][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]

A CLI for the Online labs cloud. See [https://doc.cloud.online.net/api]
for documentation.

## API

### Tokens

Action               | Implemented
---------------------|-----------------------------
Create a token       | [ ]
List all tokens      | [ ]
Retrieve a token     | [ ]
Update a token       | [ ]
Remove a token       | [ ]

### Organizations

Action               | Implementation
---------------------|------------------------------
List organizations   | [x]

### Users

Action               | Implementation
---------------------|------------------------------
List informations    | [x]

### Servers

Action               | Implementation
---------------------|------------------------------
List servers         | [x]
Create a server      | [x]
Retrieve a server    | [x]
Update a server      | [ ]
Remove a server      | [ ]
List all actions     | [ ]
Execute an action    | [x]

### Volumes

Action                     | Implementation
---------------------------|------------------------------
Volumes                    | [ ]
Create a new volume        | [ ]
Retrieves all volumes      | [ ]
Retrieves informations     | [ ]
Delete a volume            | [ ]

### Snapshots

Action                    | Implementation
--------------------------|------------------------------
Snapshots                 | [ ]
Create a snapshot         | [ ]
List all snapshots        | [ ]
Snapshot                  | [ ]
Retrieve a snapshot       | [ ]
Update a snapshot         | [ ]
Remove a snapshot         | [ ]

### Images

Action                         | Implementation
-------------------------------|------------------------------
Images                         | [ ]
Create a new image             | [ ]
List all images                | [ ]
Operation on a single image    | [ ]
Retrieves an image             | [ ]
Update an image                | [ ]
Delete an image                | [ ]

### IPs

Action                         | Implementation
-------------------------------|------------------------------
Create a new IP                | [ ]
Retrieves all IPs addresses    | [ ]
Retrieve an IP address         | [ ]
Attach an IP address           | [ ]
Remove an IP address           | [ ]

### Metadata

Action                         | Implementation
-------------------------------|------------------------------
Pimouss metadata               | [ ]


## Development

* Install requirements :

        $ make init

* Initialize dependencies :

        $ make deps

* Make the binary:

        $ make build

* It make the build inside of a Docker container and the compiled binaries will
  appear in the project directory on the host. By default, it will run a build
  which cross-compiles binaries for a variety of architectures and operating
  systems.

        $ make build-all
        $ make build-linux


## License

go-onlinelabs is free software: you can redistribute it and/or modify it under the
terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

go-onlinelabs is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU General Public License for more details.

See [LICENSE][] for the complete license.


## Changelog

A changelog is available [here](ChangeLog.md).


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-GPL_3-green.svg?style=flat
[LICENSE]: https://github.com/nlamirault/go-onlinelabs/blob/master/LICENSE
[travis]: https://travis-ci.org/nlamirault/go-onlinelabs
[badge-travis]: http://img.shields.io/travis/nlamirault/go-onlinelabs.svg?style=flat
[badge-drone]: https://drone.io/github.com/nlamirault/go-onlinelabs/status.png
[drone]: https://drone.io/github.com/nlamirault/go-onlinelabs/latest
