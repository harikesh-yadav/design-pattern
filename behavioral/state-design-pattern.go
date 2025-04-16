package main

import (
	"fmt"
	"time"
)

/***
DIFINITION:
The State design pattern in Go lets an object change its behavior based on its internal state
by using interfaces and state-specific structs, enabling clear and maintainable state transitions.
**/

type SignalState interface {
	State()
}

type YellowState struct{}

func (s *YellowState) State() {
	fmt.Println("This is yellow Signal!!!")
}

type GreenState struct{}

func (s *GreenState) State() {
	fmt.Println("This is green Signal!!!")
}

type RedState struct{}

func (s *RedState) State() {
	fmt.Println("This is Red Signal!!!")
}

type TRafficSignal struct {
	currentState SignalState
}

func (t *TRafficSignal) SetState(state SignalState) {
	t.currentState = state
}

func (t *TRafficSignal) GetState() {
	t.currentState.State()
}

func InitTRafficSignlal() *TRafficSignal {
	return &TRafficSignal{
		currentState: &YellowState{},
	}
}

func main() {
	signal := InitTRafficSignlal()
	signal.GetState()
	time.Sleep(3 * time.Second)
	signal.SetState(&GreenState{})
	signal.GetState()
	time.Sleep(3 * time.Second)
	signal.SetState(&RedState{})
	signal.GetState()
}
