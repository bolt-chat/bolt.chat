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

package events

import "time"

// Type represents an event type identifier.
type Type string

// EventMeta represents the metadata that is
// accompanied with each event.
type EventMeta struct {
	// The event identifier/type.
	Type Type `json:"t"`
	// The event creation date (client-side, untrustworthy).
	CreatedAt int64 `json:"c"`
}

// Event represents a server event.
type Event struct {
	Meta *EventMeta  `json:"e"`
	Data interface{} `json:"d"`
}

// NewEvent TODO
func NewEvent(t Type, data interface{}) *Event {
	return &Event{
		Meta: &EventMeta{
			Type:      t,
			CreatedAt: time.Now().Unix(),
		},
		Data: data,
	}
}
