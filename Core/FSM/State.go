package FSM

type State interface {
	Enter()
	Execute()
	Exit()
}
