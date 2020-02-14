package sample

/*
	$ gocyclo cycmplx/sample.go
	4 sample doSometing2 cycmplx/sample.go:13:1
	2 sample doSometing1 cycmplx/sample.go:7:1
	1 sample doSometing cycmplx/sample.go:3:1
*/

func doSometing() {

}

func doSometing1(s string) {
	if s == "1" {

	}
}

func doSometing2(s string) {
	switch s {
	case "1":

	case "2":

	default:

	}
}
