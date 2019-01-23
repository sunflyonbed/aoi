package olist

type MyNode interface {
	UniqueId() uint64
	OnMove(uint64)
	OnEnter(uint64)
	OnLeave(uint64)
}
