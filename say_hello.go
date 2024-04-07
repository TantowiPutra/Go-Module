package go_say_hello

// func SayHello(name string) (number int, message string) {
// 	number = 1
// 	message = "Hello, " + name

// 	return number, message
// }

// ! Major Changes

func SayHello(first_name string, second_name string) string{
	return "Hello, " + first_name + " " + second_name
}
