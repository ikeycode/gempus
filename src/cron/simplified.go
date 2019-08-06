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
func NewEventSimpleFormat(line string) (*Event, error) {
	return nil, nil
}
