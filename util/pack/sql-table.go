// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pack

import "github.com/abcum/surreal/sql"

type extSqlTable struct{}

func (x extSqlTable) ReadExt(dst interface{}, src []byte) {
	dst.(*sql.Table).TB = string(src)
	return
}

func (x extSqlTable) WriteExt(src interface{}) (dst []byte) {
	switch obj := src.(type) {
	case sql.Table:
		dst = []byte(obj.TB)
	case *sql.Table:
		dst = []byte(obj.TB)
	}
	return
}