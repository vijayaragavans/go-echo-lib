package goecholib

func Echo(input string) *string {
	str := "Welcome Mr." + input
	return &str
}
