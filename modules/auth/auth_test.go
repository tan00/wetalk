// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package auth

import (
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/beego/wetalk/modules/utils"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=wetalk sslmode=disable", 30)
	orm.RunSyncdb("default", true, false)
}

func TestVerifyPassword(t *testing.T) {
	pwd := "111111"
	salt := "0101010101"
	encodPwd := utils.EncodePassword(pwd, salt)
	t.Log("get encode passwd=", encodPwd)
	dbEncodPwd := salt + "$" + encodPwd

	type args struct {
		rawPwd     string
		encodedPwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case1", args{pwd, dbEncodPwd}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.args.rawPwd, tt.args.encodedPwd); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
