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
	"time"
)

// An Event represents a specific job to be executed within the crontab
// process(es). It provides tracking data to allow repeat events as well
// as the actual execution data.
type Event struct {
	id       string     // String ID for display purposes. Set by a parser.
	tid      int64      // Internal timing ID
	runFirst *time.Time // When we should first run.
}

// NewEvent will construct a new Event type
func NewEvent() *Event {
	return &Event{
		id:       "",
		tid:      0,
		runFirst: nil,
	}
}

// setTID will be called internally by the managing Tab once it can
// construct a timing ID
func (e *Event) setTID(tid int64) {
	fmt.Printf("TID now: %v\n", tid)
	e.tid = tid
}
