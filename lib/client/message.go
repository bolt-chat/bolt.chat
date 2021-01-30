// boltchat
// Copyright (C) 2021  The boltchat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Client) SendMessage(m *protocol.Message) error {
	util.WriteJson(c.Conn, *events.NewMessageEvent(m))
	return nil
}
