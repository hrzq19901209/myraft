package raft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

var commandTypes map[string]Command

func init() {
	commanTypes = map[string]Command{}
}

type Command interface {
	CommandName() string
}

type CommandApply interface {
	Apply(Context) (interface{}, error)
}

type deprecatedCommandApply interface {
	Apply(Server) (interface{}, error)
}

type CommandEncoder interface {
	Encode(w io.Writer) error
	Decode(r io.Reader) error
}

func newCommand(name string, data []byte) (Command, error) {
	command := commandTypes[name]
	if command == nil {
		return nil, fmt.Errorf("raft.Command:Unregistered command type:%s", name)
	}

	v := reflect.New(reflect.Indirect(reflect.ValueOf(command)).Type()).Interface()
	copy, ok := v.(Command)
	if !ok {
		panic(fmt.Sprintf("raft:Unable to copy command: %s (%v)", command.CommandName(), reflect.ValueOf(v).Kind().String()))
	}
	if data != nil {
		if encoder, ok := copy.(CommandEncoder); ok {
			if err := encoder.Decode(bytes.NewReader(data)); err != nil {
				return nil, err
			}
		} else {
			if err := json.NewDecoder(bytes.NewReader(data)).Decode(copy); err != nil {
				return nil, err
			}
		}
	}
	return copy, nil
}

func RegisterCommand(command Command) {
	if command == nil {
		panic(fmt.Sprintf("raft: Cannot register nil"))
	} else if commandTypes[command.CommandName()] != nil {
		panic(fmt.Sprintf("raft: Duplicate registration: %s", command.CommandName()))
	}
	commandTypes[command.CommandName()] = command
}
