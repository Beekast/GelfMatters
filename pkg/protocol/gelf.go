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

package protocol

import (
	"errors"
	"fmt"
	"strconv"
	"html"
	"encoding/json"
)


type Gelf struct {
	data	map[string]interface{}
	fields	[]string
}

func (g *Gelf) Parse(stream []byte) error {
	var f interface{}

	if err := json.Unmarshal(stream, &f); err != nil {
		fmt.Println(" [!] Error decoding json:", err.Error())
		return errors.New("Error decoding json")
	}

	g.data = f.(map[string]interface{})

	return nil
}

func (g *Gelf) Extract() []string {
	var result = make([]string, len(g.fields))

	for i, field := range g.fields {
		if value, ok := g.data[field]; ok {
			switch t := value.(type) {
				case string:
					result[i] = html.EscapeString(t)
				case float64:
					result[i] = strconv.Itoa(int(t))
			}
		} else {
			result[i] = ""
		}
	}

	return result
}
