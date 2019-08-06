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
	"strconv"
	"strings"
	"time"
)

// SimpleEvent extends the Event struct to add custom values
type SimpleEvent struct {
	tid     int64
	id      string
	command string // Really this could become a userdata function.

	timing *EventTiming // Maintain timing information
}

// setTID will be called internally by the managing Tab once it can
// construct a timing ID
func (e *SimpleEvent) setTID(tid int64) {
	e.tid = tid
}

func (e *SimpleEvent) TID() int64 {
	return e.tid
}

// Timing will return the EventTiming for this SimpleEvent
func (e *SimpleEvent) Timing() *EventTiming {
	return e.timing
}

// Execute currently does nothing
func (e *SimpleEvent) Execute() (int, error) {
	t := e.Timing()
	day := "today"
	if t.tm.Day() > time.Now().Day() {
		day = "tomorrow"
	}
	fmt.Printf("%02d:%02d %s - %s\n", t.tm.Hour(), t.tm.Minute(), day, e.command)
	return 255, nil
}

// ID will return the ID for the event
func (e *SimpleEvent) ID() string {
	return e.id
}

// NewEventSimpleFormat will attempt to construct an event for the
// simplified crontab format. Effectively, this supports only 3 fields:
//
// MINUTE   HOUR      COMMAND
// 0-59/*   0-59/*    /some/command
//
// Minute and Hour may be either numerical, or the '*' operator to match
// and execute every minute of the specified hour, once an hour every day,
// or some combination of both.
//
// Note the simplified event does not have any notion of days/months nor
// does it support time ranges.
func NewEventSimpleFormat(line string) (Event, error) {
	var err error
	var min int64
	var hour int64

	tokens := strings.Fields(line)
	if len(tokens) < 3 {
		return nil, fmt.Errorf("invalid number of tokens")
	}

	// Command to run
	command := strings.Join(tokens[2:], "")

	// Minutes
	if tokens[0] == "*" {
		min = -1
	} else {
		min, err = strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			return nil, err
		}
	}

	// Hours
	if tokens[1] == "*" {
		hour = -1
	} else {
		hour, err = strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return NewEventSimpleFormatValues(int(hour), int(min), command), nil
}

// NewEventSimpleFormatValues will return a new event with the given
// values for timing.
func NewEventSimpleFormatValues(hour, minute int, command string) Event {
	event := &SimpleEvent{
		id:      fmt.Sprintf("run %v @ M(%v) H(%v)", command, minute, hour),
		command: command,
		timing:  NewEventTiming(hour, minute),
	}
	event.timing.NextTimestamp(time.Now(), false)
	return event
}
