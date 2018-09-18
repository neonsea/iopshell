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

package setting

import (
	"fmt"

	"github.com/chzyer/readline"
	"gitlab.com/c-/iopshell/internal/cmd"
	"gitlab.com/c-/iopshell/internal/connection"
	"gitlab.com/c-/iopshell/internal/textmutate"
)

var (
	// ConnChannel listens for connection requests
	ConnChannel = make(chan []string)
	// ScriptIn listens for script run requests
	ScriptIn = make(chan []string)
	// ScriptRet houses return codes from scripts
	ScriptRet = make(chan error)
	// In receives incoming messages
	In = make(chan interface{})
	// Out houses outgoing messages
	Out = make(chan interface{})
)

// Opt houses some options
type Opt struct {
	Description string
	Val         interface{}
}

// ShellVars house some important structs, so they can be accessed elsewhere
type ShellVars struct {
	Conn      *connection.Connection
	Completer readline.PrefixCompleter
	Instance  *readline.Instance
	Opts      map[string]*Opt
}

// Set sets opt to val
func (s *ShellVars) Set(opt string, val interface{}) bool {
	if o, ok := s.Opts[opt]; ok {
		o.Val = val
		if opt == "verbose" { // there's got to be a better way to do this
			textmutate.Verbose = val.(bool)
		}
		return true
	}
	return false
}

// Get returns the specified option's pointer, or false if it doesn't exist
func (s *ShellVars) Get(opt string) (*Opt, bool) {
	if o, ok := s.Opts[opt]; ok {
		return o, true
	}
	return &Opt{}, false
}

// GetS returns the specified option's value as a string, or false if it doesn't exist
func (s *ShellVars) GetS(opt string) (string, bool) {
	if o, ok := s.Opts[opt]; ok {
		val, err := o.Val.(string)
		return val, err
	}
	return "", false
}

// GetB returns the specified option's value as a bool, or false if it doesn't exist
func (s *ShellVars) GetB(opt string) (bool, bool) {
	if o, ok := s.Opts[opt]; ok {
		val, err := o.Val.(bool)
		return val, err
	}
	return false, false
}

// GetI returns the specified option's value as an int, or false if it doesn't exist
func (s *ShellVars) GetI(opt string) (int, bool) {
	if o, ok := s.Opts[opt]; ok {
		val, err := o.Val.(int)
		return val, err
	}
	return 0, false
}

// UpdatePrompt refreshes the prompt and sets it according to current status
func (s *ShellVars) UpdatePrompt() {
	var prompt string
	if s.Conn.Ws == nil {
		// Not connected
		prompt = "\033[91miop\033[0;1m$\033[0m "
	} else {
		if s.Conn.User == "" {
			// Connected but not authenticated
			prompt = "\033[32miop\033[0;1m$\033[0m "
		} else {
			// Connected and authenticated
			prompt = fmt.Sprintf("\033[32miop\033[0m %s\033[0;1m$\033[0m ", s.Conn.User)
		}
	}
	if s.Instance != nil {
		s.Instance.SetPrompt(prompt)
		s.Instance.Refresh()
	}
}

// UpdateCompleter adds commands registered with .Register() to the autocompleter
func (s *ShellVars) UpdateCompleter(cmdlist map[string]cmd.Command) {
	s.Completer = *readline.NewPrefixCompleter()
	s.Completer.SetChildren(*new([]readline.PrefixCompleterInterface))

	commands := make([]string, len(cmdlist))
	i := 0
	for c := range cmdlist {
		commands[i] = c
		i++
	}
	for _, c := range commands {
		s.Completer.Children = append(s.Completer.Children, readline.PcItem(c))
	}
}

// Init values for Vars
func genVars() *ShellVars {
	var vars ShellVars
	vars.Opts = make(map[string]*Opt)
	vars.Opts["host"] = &Opt{
		Description: "Host to connect to by default",
		Val:         "192.168.1.1",
	}
	vars.Opts["user"] = &Opt{
		Description: "User to authenticate as by default",
		Val:         "user",
	}
	vars.Opts["pass"] = &Opt{
		Description: "Password to authenticate with by default",
		Val:         "user",
	}
	vars.Opts["verbose"] = &Opt{
		Description: "Print verbose messages",
		Val:         textmutate.Verbose,
	}
	return &vars
}

// Vars is an instance of ShellVars
var Vars = *genVars()
