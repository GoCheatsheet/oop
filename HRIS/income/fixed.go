package income

type fixed struct {
	projectName  string
	biddedAmount int
}

func (fb fixed) calculate() int {
	return fb.biddedAmount
}

func (fb fixed) source() string {
	return fb.projectName
}

func NewFixed(projectName string, biddedAmount int) fixed {
	fb := fixed{projectName, biddedAmount}
	return fb
}
