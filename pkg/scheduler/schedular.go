package schedular

/*
Schedular picks the right candiate from the worker nodes
*/
type Schedular interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
