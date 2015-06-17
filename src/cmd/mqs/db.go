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
)

type dbHandler interface {
	Execute(statement string) (sql.Result, error)
	Query(statement string) (*sql.Rows, error)
	KillQuery(qid query_id) error
}
