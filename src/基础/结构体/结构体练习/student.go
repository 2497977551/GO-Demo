package main

import "fmt"

type student struct {
	id          int    // 学员id
	name, class string // 学员姓名与班级
}

// student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type setStudent struct {
	students []*student
}

func newSetStudent() *setStudent {
	return &setStudent{
		students: make([]*student, 0, 20),
	}
}

//	添加学员
func (s *setStudent) addStudent(newSt *student) {
	s.students = append(s.students, newSt)
}

// 编辑学员
func (s *setStudent) UpdStudent(newSt *student) {
	for k, v := range s.students {
		if v.id == newSt.id {
			s.students[k] = newSt
			return
		}
	}
	fmt.Println("输入的学号不存在")
}

// 展示所有学员,
func (s *setStudent) ShowStudent() {
	for _, v := range s.students {
		fmt.Printf("学号:%d，姓名:%s，班级%s\n", v.id, v.name, v.class)
	}
}
