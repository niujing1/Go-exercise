type Talk interface {
	Hello(userName string) string
	Talk(heart string)(saying string, end bool, err error)
}
