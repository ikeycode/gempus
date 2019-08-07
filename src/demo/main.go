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

package main

import (
	"bufio"
	"cron"
	"fmt"
	"os"
	"time"
)

func main() {
	tab := cron.NewTab()

	tm := time.Now()
	var err error

	if len(os.Args) > 1 {
		tm, err = time.Parse("15:04", os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	// Loop stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		event, err := cron.NewEventSimpleFormat(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Cheekily force the timestamp change now
		event.Timing().NextTimestamp(tm, false)
		tab.PushEvent(event)
		event.Execute(tm)
	}
}
