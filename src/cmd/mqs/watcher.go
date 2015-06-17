// Copyright © 2015 Stephane Bunel - https://bitbucket.org/StephaneBunel
//
// Licensed under the Simple Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://opensource.org/licenses/Simple-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"strconv"
	"strings"
)

type Watcher interface {
	Refresh()
	QueryMap()
}

type Query struct {
	dbh  dbHandler
	id   query_id
	user string
	time int
	info string
	tags map[string]string
}

type query_id int
type QueryList []query_id
type QueryMap map[query_id]*Query

func (query *Query) HandleTags() {
	for tag, value := range query.tags {
		switch tag {

		case "mqs-log-longer":
			timeout, err := strconv.Atoi(value)
			if err != nil {
				log.Println(err)
				continue
			}
			if query.time > timeout {
				log.Printf("Process %v: %s\n", query.id, query.info)
			}

		case "mqs-timeout":
			timeout, err := strconv.Atoi(value)
			if err != nil {
				log.Println(err)
				continue
			}
			if query.time > timeout {
				query.dbh.KillQuery(query.id)
				log.Printf("Process %v: Query terminated (killed) after %v seconds: %s\n", query.id, timeout, query.info)
			}
		}
	}
}

func (q *Query) ParseTags(delim string) *Query {
	q.tags = make(map[string]string)

	//
	if !strings.HasPrefix(q.info, "/*") {
		return q
	}

	// Find the end of the comment
	end := strings.Index(q.info, "*/")
	if end == -1 {
		return q
	}

	// Split comment as an array of string splitted by spaces
	args := QuottedFields(strings.TrimSpace(q.info[2:end]))

	//
	for _, arg := range args {
		if arg[0] == delim[0] {
			// Asume this is a tag
			if p := strings.Index(arg, "="); p != -1 {
				q.tags[arg[1:p]] = arg[p+1 : len(arg)]
			}
		}
	}

	return q
}
