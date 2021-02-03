// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connect

import (
	"fmt"

	"github.com/bolt-chat/client/cli/cmd"
)

var ConnectCommand = &cmd.Command{
	Name:    "connect",
	Desc:    "Connects to a Bolt instance.",
	Usage:   "<host> [identity]",
	Handler: connectHandler,
}

func connectHandler(args []string) {
	fmt.Println("Hello, world!", args)
}
