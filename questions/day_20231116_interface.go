package questions

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func Q20231116() {
	//var peo People = Student{}
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

func q20231116() {
	//cannot use Student{} (value of type Student) as People value in variable declaration: Student does not implement People (method Speak has pointer receiver)
}
