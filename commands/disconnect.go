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
	"gitlab.com/neonsea/iopshell/internal/cmd"
	"gitlab.com/neonsea/iopshell/internal/setting"
)

var Disconnect = cmd.Command{
	Name:        "disconnect",
	UsageText:   "disconnect",
	Description: "Disconnects from the currently connected host",
	Action:      disconnect,
	MaxArg:      1,
}

func disconnect(param []string) {
	setting.Cmd <- []string{"disconnect"}
}

func init() {
	Disconnect.Register()
}