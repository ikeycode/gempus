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
	"time"
)

// EventTimeFlag provides a strongly typed bitmask for event timing flags
type EventTimeFlag uint

const (
	// EventPerMinute will be set if repeating-per-minute is required
	EventPerMinute EventTimeFlag = 1 << iota

	// EventPerHour will be set if repeating-per-hour is required
	EventPerHour EventTimeFlag = 1 << iota

	// EventPerWeek will be set if repeating-per-week is required
	EventPerWeek EventTimeFlag = 1 << iota

	// EventPerMonth will be set if repeating-per-month is required
	EventPerMonth EventTimeFlag = 1 << iota

	// EventPerYear will be set if repeating-per-year is required
	EventPerYear EventTimeFlag = 1 << iota

	// EventRepeats is set whenever a repeating flag is used
	EventRepeats EventTimeFlag = 1 << iota
)

// Event provides a simplistic interface to implement events
type Event interface {
	ID() string                         // ID should return a display ID
	TID() int64                         // Time ID
	Execute(now time.Time) (int, error) // Attempt to execute event
	setTID(int64)                       // Internal method to set the TID

	Timing() *EventTiming // Timing returns the event timing structure
}

// EventTiming allows us to determine if and when an event runs
type EventTiming struct {
	Hour   int // -1 Indicates running every hour
	Minute int // -1 Indicates running every minute

	flags EventTimeFlag // Event timing flags

	tm time.Time

	// Day uint8
	// Month uint8
}

// EventQueue wraps the Event type to provide time-based sorting
// of cron events. Honestly, a less lame solution is to use the
// container/heap interface and some channels with a single writer.
type EventQueue []Event

// ShouldRun will determine if we actually need to be run.
// 'now' should be a current-time value
func (t *EventTiming) ShouldRun(now time.Time) bool {
	return t.tm.Before(now)
}

// Repeats will return true if the event repeats
func (t *EventTiming) Repeats() bool {
	return t.HasFlag(EventRepeats)
}

// NewEventTiming will set up a basic EventTiming object
func NewEventTiming(hour, minute int) *EventTiming {
	timing := &EventTiming{
		Hour:   hour,
		Minute: minute,
		flags:  0,
	}

	// Repeats hourly
	if hour < 0 {
		timing.flags |= EventPerHour | EventRepeats
	}

	// Repeats every minute
	if minute < 0 {
		timing.flags |= EventPerMinute | EventRepeats
	}

	return timing
}

// HasFlag returns true if the flag is set
func (t *EventTiming) HasFlag(mask EventTimeFlag) bool {
	return t.flags&mask == mask
}

// NextTimestamp sets the timestamp for the next time the
// event should run.
func (t *EventTiming) NextTimestamp(now time.Time, ran bool) {
	// Run every hour
	tm := now

	// Run every minute of every hour
	if t.HasFlag(EventPerHour | EventPerMinute | EventRepeats) {
		// FIXME
		// Ugly if-check required due to test. Basically they expect
		// cron job to run from minute-0 as opposed to minute-after
		// invocation. Will remove in future!
		if ran {
			tm = tm.Add(time.Minute)
		}
		goto compl
	}

	if t.HasFlag(EventPerHour) {
		// Every hour
		if ran || t.Minute < tm.Minute() {
			tm = tm.Add(time.Hour)
		}
	} else {
		// Specified hour
		hour := t.Hour - tm.Hour()
		tm = tm.Add(time.Duration(hour) * time.Hour)
	}

	if t.HasFlag(EventPerMinute) {
		// Add one minute from now.
		if ran {
			tm = tm.Add(time.Minute)
		}
	} else {
		// Set exact minute
		minute := t.Minute - tm.Minute()
		tm = tm.Add(time.Duration(minute) * time.Minute)
	}

	// Now check if this time is back in time..
	if tm.Before(now) {
		tm = tm.Add(time.Hour * time.Duration(24))
	}

	// Reset the minutes to start of the hour
	if now.Before(tm) && t.HasFlag(EventPerMinute) {
		tm = tm.Add(-time.Duration(tm.Minute()) * time.Minute)
	}

compl:
	t.tm = tm
}

func (eq EventQueue) Less(i, j int) bool {
	a := eq[i].Timing()
	b := eq[j].Timing()
	return a.tm.Before(b.tm)
}

func (eq EventQueue) Len() int {
	return len(eq)
}

func (eq EventQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
}
