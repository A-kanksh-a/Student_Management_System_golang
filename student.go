package main

type student struct {
	id    string
	name  string
	marks string
}

func (s *student) update(id string, name string, marks string) {
	s.name = name
	s.marks = marks
}

func createStudent(id string, name string, marks string) (s student) {
	s1 := student{id, name, marks}
	return s1
}
