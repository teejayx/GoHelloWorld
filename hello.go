package main

import "fmt"

func main(){
	fmt.Println(Hello("Chris", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Holla, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	

	return greetingPrefix(language) + name
}


func greetingPrefix(language string) (prefix string){

switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
return 


}