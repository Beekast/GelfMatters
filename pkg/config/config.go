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

package config

import (
	"os"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)


// YAML format
//
// server:
//   bind: 0.0.0.0
//   port: 12301
// mattermost:
//   url: https://mattermost.or.slack/hooks/my_token
//   payload: "{\"username\": \"graylog\", \"text\": \"env: {}\nlevel: {}\nservice: {}\nmsg: {}\"}"
// gelf:
//   - _k8s_namespace
//   - _hornet_level
//   - _hornet_name
//   - _hornet_msg
//
type Conf struct {
	Server struct {
		Bind	string
		Port	string
	}
	Mattermost struct {
		Url		string
		Payload	string
	}
	Gelf []string `yaml:",flow"`
}

func Import() *Conf {
	f, err := ioutil.ReadFile("gelfmatters.conf")
	if err != nil {
		fmt.Println(" [!] Error at reading config file:", err.Error())
		os.Exit(1)
	}

	var c Conf
	yaml.Unmarshal(f, &c)

	return &c
}
