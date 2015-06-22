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

// bin/mqs -dsn="stephane:totosql@/mysql"
// MariaDB/Mysql user MUST have the PROCESS and SUPER privileges.

import (
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DSN_DEFAULT = "root:@/mysql"
	DSN_ENVVAR  = "MQS_DSN"
)

func main() {
	var dsn, dsn_cli string
	var dsn_envvar string = os.Getenv(DSN_ENVVAR)

	flag.StringVar(&dsn_cli, "dsn", DSN_DEFAULT, "DSN to connect with MySQL")
	flag.Parse()

	dsn = dsn_cli
	if dsn_envvar != "" && dsn_cli == DSN_DEFAULT {
		dsn = dsn_envvar
	}

	// Create database handler and watcher
	dbh := NewMySQLHandler(dsn)
	Watcher := NewMySQLQueryWatcher(dbh)

	log.Println("Sniper in position.")
	for {
		Watcher.Refresh()
		for _, query := range Watcher.QueryMap() {
			query.ParseTags("-")
			query.HandleTags()
		}
		time.Sleep(1 * time.Second)
	}
}
