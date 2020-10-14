// Copyright 2020 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileutil

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSON reads a JSON file into an interface.
func ReadJSON(s string, i interface{}) error {
	bb, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bb, &i); err != nil {
		return err
	}

	return nil
}

// WriteJSON writes a JSON file from an interface.
func WriteJSON(s string, i interface{}) error {
	bb, err := json.Marshal(i)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(s, bb, 0644); err != nil {
		return err
	}

	return nil
}
