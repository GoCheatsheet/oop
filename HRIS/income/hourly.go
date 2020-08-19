package income

type hourly struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (tm hourly) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm hourly) source() string {
	return tm.projectName
}

func NewHourly(projectName string, noOfHours int, hourlyRate int) hourly {
	hr := hourly{projectName, noOfHours, hourlyRate}
	return hr
}
