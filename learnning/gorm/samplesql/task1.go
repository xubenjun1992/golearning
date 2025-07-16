package samplesql

import (
	"fmt"
	"log"
	"strings"

	"main.go/gorm/samplesql/core"
)

type Students struct {
	Id    int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}

func SaveStudent(student *Students) {
	//insert into students(`name`, `age`, `grade`) values ('张三', 20, '三年级')
	core.DB.Create(student)
	log.Fatal("插入成功")
}

func QueryByCondition(group core.ConditionGroup) []Students {
	studentList := []Students{}
	whereClauses := []string{}
	var args []interface{}
	for _, v := range group.Conditions {
		whereClauses = append(whereClauses, fmt.Sprintf("%s %s ?", v.ColumnName, v.Cond))
		args = append(args, v.Val)
	}
	var joinWith = "AND"
	if group.JoinWith == "OR" {
		joinWith = "OR"
	}
	whereSql := strings.Join(whereClauses, joinWith)
	if whereSql != "" {
		core.DB.Where(whereSql, args...).Scan(&studentList)
	} else {
		core.DB.Find(&studentList)
	}

	return studentList
}

func UpdateStudent() {
	core.DB.Where("name = ?", "张三").Update("grade", "四年级")
}

func DeleteStudent() {
	core.DB.Where("age < 15").Delete(&Students{})
}
