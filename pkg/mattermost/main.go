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

package mattermost

import (
	"fmt"
	"net/url"
	"net/http"
	"crypto/tls"
)


type Mattermost struct {
	url		string
	client	http.Client
	Receive	chan []string
	payload	string
}

func New(payload string) *Mattermost {
	return &Mattermost{ Receive: make(chan []string), payload: payload }
}

func (m *Mattermost) Connect(mattermost_url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ InsecureSkipVerify: true },
		DisableCompression: true,
	}

	mattermost_client := http.Client{ Transport: tr }

	m.url = mattermost_url
	m.client = mattermost_client
}

// TODO: implement exit code
func (m *Mattermost) LoopOnEvents() {
	go func() {
		for {
			var data = <-m.Receive

			payload := translateString(m.payload, data)

			fmt.Println("Will be sent:", payload)
			resp, err := m.client.PostForm(m.url, url.Values{"payload": {payload}})
			if err != nil {
				fmt.Println(" [!] Error sending msg to mattermost:", err.Error())
				fmt.Println(" [!] Where trying to send:", payload)
			}
			if resp != nil {
				resp.Body.Close()
			}
		}
	}()
}
