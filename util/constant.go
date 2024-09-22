package util

type AnimeListStatus int
const (
	PlanToWatch	AnimeListStatus = iota + 1 // EnumIndex = 1
	Watching     						// EnumIndex = 2
	Completed    	
	OnHold      	
	Dropped      	
)	

func (a AnimeListStatus) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[a-1]
}
func (a AnimeListStatus) EnumIndex() int {
	return int(a)
}