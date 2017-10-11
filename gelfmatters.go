/*
Copyright 2017 Beekast.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"gitlab.beekast.info/infra/gelfmatters/pkg/mattermost"
	"gitlab.beekast.info/infra/gelfmatters/pkg/server"
	"gitlab.beekast.info/infra/gelfmatters/pkg/config"
	"gitlab.beekast.info/infra/gelfmatters/pkg/protocol"
)


func main() {
	// Import configuration from the yaml gelfmatters.conf file
	conf := config.Import()
	var gelf protocol.Protocol = protocol.New(protocol.GELF, conf.Gelf)

	// Handle mattermost connection
	m := mattermost.New(conf.Mattermost.Payload)
	m.Connect(conf.Mattermost.Url)
	m.LoopOnEvents()

	// Initialise server and wait for connections
	s := server.New(conf.Server.Bind, conf.Server.Port, gelf)
	s.SetSender(m.Receive)
	s.AcceptConnections()
}
