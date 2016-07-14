package raft

import (
	"io"
)

type JoinCommand interface {
	Command
	NodeName() string
}

//Join Command
type DefaultJoinCommand struct {
	Name             string `json:"name"`
	ConnectionString string `json:"connectionString"`
}

type LeaveCommand interface {
	Command
	NodeName() string
}

//Leave Command
type DefaultLeaveCommand struct {
	Name string `json:"name"`
}

//NOP command
type NOPCommand struct {
}

func (c *DefaultJoinCommand) CommandName() string {
	return "raft:join"
}

func (c *DefaultJoinCommand) Apply(server Server) (interface{}, error) {
	err := server.AddPeer(c.Name, c.ConnectionString)
	return []byte("join"), err
}

func (c *DefaultJoinCommand) NodeName() string {
	return c.Name
}

// The name of the Leave command in the log
func (c *DefaultLeaveCommand) CommandName() string {
	return "raft:leave"
}

func (c *DefaultLeaveCommand) Apply(server Server) (interface{}, error) {
	err := server.RemovePeer(c.Name)

	return []byte("leave"), err
}
func (c *DefaultLeaveCommand) NodeName() string {
	return c.Name
}

// The name of the NOP command in the log
func (c NOPCommand) CommandName() string {
	return "raft:nop"
}

func (c NOPCommand) Apply(server Server) (interface{}, error) {
	return nil, nil
}

func (c NOPCommand) Encode(w io.Writer) error {
	return nil
}

func (c NOPCommand) Decode(r io.Reader) error {
	return nil
}
