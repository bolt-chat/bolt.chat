// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
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

package events

// MotdType is the event type used for the Message-of-the-Day (MOTD).
const MotdType Type = "motd"

// MotdEvent is the event used for the Message-of-the-Day (MOTD).
type MotdEvent struct {
	BaseEvent
	Motd string
}

// NewMotdEvent constructs a new MotdEvent
func NewMotdEvent(motd string) *MotdEvent {
	return &MotdEvent{
		BaseEvent: *NewBaseEvent(MotdType),
		Motd:      motd,
	}
}
