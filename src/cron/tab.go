//
// Copyright © 2019 Ikey Doherty <contactikeydoherty@gmail.com>
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
//

package cron

import (
	"fmt"
	"sync/atomic"
)

// A Tab is used internally to store cron events, which are then
// available for query and scheduling.
//
// As a design decision, the Tab is not responsible for the parsing
// of crontab files, as it would force the scheduler component to become
// I/O bound. Instead it becomes the responsibility of the calling program
// to feed parsed events.
type Tab struct {
	events []*Event

	tid int64
}

// NewTab will construct a new Tab type and prepare it for
// use within the calling process.
func NewTab() *Tab {
	return &Tab{
		events: nil,
	}
}

// PushEvent will attempt to insert the event into the crontab for
// future scheduling
func (t *Tab) PushEvent(e *Event) {
	e.setTID(t.nextTID())
	fmt.Println("Not yet implemented")
}

func (t *Tab) nextTID() int64 {
	return atomic.AddInt64(&t.tid, 1)
}