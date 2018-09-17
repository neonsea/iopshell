/*
   Copyright (c) 2018 Rasmus Moorats (neonsea)

   This file is part of iopshell.

   iopshell is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   iopshell is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with iopshell. If not, see <https://www.gnu.org/licenses/>.
*/

package commands

import (
	"fmt"

	"gitlab.com/c-/iopshell/internal/cmd"
	"gitlab.com/c-/iopshell/internal/setting"
)

var auth = cmd.Command{
	Name:        "auth",
	UsageText:   "auth [user pass]",
	Description: "Authenticates as [user] with [pass]. If none are specified, uses the values from settings.",
	Action:      authRun,
	MinArg:      1,
	MaxArg:      3,
}

func authRun(param []string) {
	var user, pass string
	switch len(param) {
	case 0:
		user = setting.Vars.Opts.User
		pass = setting.Vars.Opts.Pass
	case 1:
		fmt.Println("Both user and pass need to specified, or none at all.")
		return
	case 2:
		user = param[0]
		pass = param[1]
	}

	// Authenticating is just another call
	setting.Vars.Conn.Call("session", "login", map[string]interface{}{"username": user,
		"password": pass})
}

func init() {
	auth.Register()
}
