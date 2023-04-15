package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id, Name, Address, Job, Reason string
}

var students []Student = []Student{
	{
		Id:      "1",
		Name:    "Naufal",
		Address: "Jalan Raya Durian Runtuh",
		Job:     "Mahasiswa",
		Reason:  "Ingin berkarir sebagai Front End Developer",
	},
	{
		Id:      "2",
		Name:    "Hilmy",
		Address: "Jalan Raya Kelapa Sawit",
		Job:     "Mahasiswa",
		Reason:  "Ingin berkarir sebagai Back End Developer",
	},
	{
		Id:      "3",
		Name:    "Musawwir",
		Address: "Jalan Raya Tanah Mas",
		Job:     "Mahasiswa",
		Reason:  "Ingin berkarir sebagai Full Stack Developer",
	},
}

func findStudentByName(studentName string) (Student, error) {
	for _, value := range students {
		if value.Name == studentName {
			return value, nil
		}
	}
	return Student{}, errors.New("Student not found")
}

func nameFriend(input []string) (string, error) {
	if len(input) < 2 {
		return "", errors.New("message : mohon untuk menjalankan program dengan menambahkan go run main.go [nama]")
	}
	return input[1], nil
}

func studentOutput(student Student) {
	fmt.Println("ID        : ", student.Id)
	fmt.Println("Nama      : ", student.Name)
	fmt.Println("Alamat    : ", student.Address)
	fmt.Println("Pekerjaan : ", student.Job)
	fmt.Println("Alasan    : ", student.Reason)
}

func main() {
	name, err := nameFriend(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}

	result, err := findStudentByName(name)
	if err != nil {
		log.Fatalln(err.Error())
	}

	studentOutput(result)
}
