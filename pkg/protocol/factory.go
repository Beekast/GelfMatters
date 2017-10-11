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

const (
	GELF = 1
)


type Protocol interface {
	Parse(stream []byte) error
	Extract() []string
}

func New(proto int, fields []string) Protocol {
	switch proto {
	case GELF:
		return &Gelf{ fields: fields }
	default:
		panic("Protocol not supported")
	}
}
