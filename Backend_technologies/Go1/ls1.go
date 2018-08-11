package main

import	"fmt"

func main()  {
	// по умолчанию инициализирует сам    
	var num int
	// короткое объявление 
	num1 := 10
	// несколько 
	weight, height = 11,21
	// есть int от 8 до 64
	// есть комплексные переменные 
	// строки
	var str string
	// из строки коробки utf-8
	// строки неизменяемые
	// длина строки в байтах
	byteLen := len(str) 
	//константы
	const (
		hello = "asdf"
		e = 2.324
	)

	const (
		zero = iota
		- //пустая переменная
		three // 3
		// можно объявлять нетипизированные константы
	)

	// размер массива является частью его типа

	// инициализация значениями по-умолчанию
	var a1 [3]int // [0,0,0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool // [false,false,false,false]
	fmt.Println("a2", a2)

	// определение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	// проверка при компиляции или при выполнении
	// invalid array index 4 (out of bounds for 3-element array)
	// a3[idx] = 12

	// инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// сразу с нужной ёмкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	// если ключа нет - вернёт значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	// удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)

	// указатели
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // новый указатель на переменную a

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12
}