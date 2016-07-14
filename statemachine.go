package raft

//保存并从log中恢复
type StateMachine interface {
	Save() ([]byte, error)
	Recovery([]byte) error
}
