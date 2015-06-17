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
	"database/sql"
	"log"
)

type MySQLQueryWatcher struct {
	db      *MySQLHandler
	queries QueryMap
}

func (w *MySQLQueryWatcher) rowToQuery(rows *sql.Rows) *Query {
	q := new(Query)

	if err := rows.Scan(&q.id, &q.user, &q.time, &q.info); err != nil {
		log.Println(err)
		return nil
	}
	return q
}

func (w *MySQLQueryWatcher) Refresh() {
	w.queries = make(QueryMap)
	rows, _ := w.db.Query(SQL_SHOW_PROCESS_LIST)
	if rows == nil {
		return
	}
	for rows.Next() {
		// Convert row to Query object
		q := w.rowToQuery(rows)
		if q != nil {
			q.dbh = w.db
			w.queries[q.id] = q
		}
	}
}

func (w *MySQLQueryWatcher) QueryMap() QueryMap {
	return w.queries
}

func NewMySQLQueryWatcher(dbh *MySQLHandler) *MySQLQueryWatcher {
	h := new(MySQLQueryWatcher)
	h.db = dbh
	return h
}
