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

	_ "github.com/go-sql-driver/mysql"
)

const (
	SQL_SHOW_PROCESS_LIST = "SELECT id, user, FLOOR(time_ms/1000), info FROM INFORMATION_SCHEMA.PROCESSLIST WHERE info is not null"
	SQL_KILL_QUERY        = "KILL QUERY ?"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Execute(statement string) (sql.Result, error) {
	return handler.Conn.Exec(statement)
}

func (handler *MySQLHandler) Query(statement string) (*sql.Rows, error) {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func (handler *MySQLHandler) Connexion() *sql.DB {
	return handler.Conn
}

func (handler *MySQLHandler) KillQuery(qid query_id) error {
	_, err := handler.Conn.Exec(SQL_KILL_QUERY, qid)
	if err != nil {
		log.Println(err)
	}
	return err
}

func NewMySQLHandler(dsn string) *MySQLHandler {
	h := new(MySQLHandler)
	conn, _ := sql.Open("mysql", dsn)
	h.Conn = conn
	return h
}
