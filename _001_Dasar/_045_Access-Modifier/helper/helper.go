package helper

var version string = "1.0.0" // tidak bisa diakses dari luar
var Application string = "Golang"

func sayGoodBye() string { // tidak bisa diakses dari luar
	return "Bye!"
}

func SayHello(name string) string {
	return "Hello " + name
}
