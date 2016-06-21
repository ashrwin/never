package main

type StringMatcher func(string) bool

func inSlice(a []string, matcher StringMatcher) bool {
	for _, val := range a {
		if matcher(val) {
			return true
		}
	}
	return false
}

func printSeperatorLine(f func(format string, a ...interface{})) {
	x := "------------------------------------------------------------------------------------"
	f("\n%s\n", x)
}

func checkError(err error) {
	if err != nil {
		fatal(err)
	}
}
