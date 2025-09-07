package main

import (
	"fmt"
)

func main() {
	p := base.NewPerson("Василиса Премудрая")
	fmt.Println(p[len(p)-1])
}

//надо понять
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//const layout = "02.01.2006"
//
//func Age(birthday string) (int, error) {
//
//	birthDate, err := time.Parse(layout, birthday)
//	if err != nil {
//		return 0, fmt.Errorf("неверный формат даты")
//	}
//
//	now := time.Now()
//
//	if birthDate.After(now) {
//		return 0, fmt.Errorf("дата рождения не может быть в будущем")
//	}
//
//	age := now.Year() - birthDate.Year()
//
//	if now.Month() < birthDate.Month() ||
//		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
//		age--
//	}
//
//	return age, nil
//}
//
//func main() {
//	for _, v := range []string{"04.01.1969", "28.07.1995",
//		"16.12.2007", "07.03.1947"} {
//		age, err := Age(v)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		fmt.Println(v, ":", age)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func whatTime(friend string) time.Time {
//	var friends = map[string]string{
//		"Серёга": "Омск",
//		"Соня":   "Москва",
//		"Дима":   "Екатеринбург",
//		"Алина":  "Новосибирск",
//		"Егор":   "Калининград",
//	}
//	var offsetUTC = map[string]int{
//		"Санкт-Петербург": 3,
//		"Москва":          3,
//		"Самара":          4,
//		"Новосибирск":     7,
//		"Екатеринбург":    5,
//		"Казань":          3,
//		"Омск":            6,
//		"Хабаровск":       10,
//		"Пермь":           5,
//		"Краснодар":       3,
//		"Калининград":     2,
//	}
//
//	utcf := time.Now().UTC()
//	city, exists := friends[friend]
//	if !exists {
//		return utcf
//	}
//	offset, exists := offsetUTC[city]
//	if !exists {
//		return utcf
//	}
//	resultTime := utcf.Add(time.Duration(offset) * time.Hour)
//	return resultTime
//}
//
//func main() {
//	for _, friend := range []string{"Соня", "Дима", "Егор"} {
//		t := whatTime(friend)
//		fmt.Printf("%s: %d ч.\n", friend, t.Hour())
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// создайте переменные типа time.Time
//	gola := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
//	gol := time.Date(2023, 8, 8, 0, 0, 0, 0, time.UTC)
//	// получите интервал между датами
//	unter := gol.Sub(gola)
//	// вычислите количество дней
//	days := int(unter.Hours() / 24)
//	fmt.Printf("Между выходами версий прошло %d дней", days)
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	// дата выхода первой серии
//	begin := time.Date(2011, 4, 17, 0, 0, 0, 0, time.UTC)
//	// укажите дату выхода последней серии
//	end := time.Date(2019, 4, 15, 0, 0, 0, 0, time.UTC)
//	// вычислите, сколько времени шёл сериал
//	duration := end.Sub(begin)
//	// разделите количество часов duration.Hours() на 24,
//	// чтобы получить количество дней
//	fmt.Printf("Сериал шёл %d дней", int(duration.Hours()/24))
//}

//надо понять

//package main
//
//import (
//	"fmt"
//	"slices"
//)
//
//// Mode возвращает моды числовой последовательности.
//func Mode(s []int) ([]int, int) {
//	if len(s) == 0 {
//		return []int{}, 1
//	}
//	m := make(map[int]int, len(s))
//	for _, v := range s {
//		m[v]++
//	}
//	maxCount := 1
//	for _, count := range m {
//		// встроенные функции min/max появились в go 1.21
//		// поэтому можно использовать
//		maxCount = max(maxCount, count)
//		// вместо
//		// if count > maxCount {
//		//    maxCount = count
//		// }
//	}
//	result := make([]int, 0)
//	if maxCount <= 1 {
//		return result, 1
//	}
//
//	for i, count := range m {
//		if count == maxCount {
//			result = append(result, i)
//		}
//	}
//	slices.Sort(result)
//	return result, maxCount
//}
//
//func main() {
//	lists := [][]int{
//		{},
//		{57},
//		{78, -7},
//		{99, 200, 0},
//		{4, 4, 4, 3},
//		{102, -7, 44, -7, 102},
//		{82, -23, 1, 5, 98, 100},
//		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
//	}
//	modes := [][]int{
//		{}, {}, {}, {},
//		{4},
//		{-7, 102}, {},
//		{20000},
//	}
//	mcount := []int{
//		1, 1, 1, 1, 3, 2, 1, 3,
//	}
//	for i, list := range lists {
//		mode, count := Mode(list)
//		if len(mode) != len(modes[i]) {
//			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
//		} else {
//			for j, v := range mode {
//				if v != modes[i][j] {
//					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
//				}
//			}
//		}
//		if count != mcount[i] {
//			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
//		}
//	}
//	fmt.Println("Тестирование завершено")
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//// Average возвращает среднее арифметическое элементов слайса []int.
//// Напишите код функции
//func Average(numbers []int) float64 {
//	if len(numbers) == 0 {
//		return 0.0
//	}
//
//	sum := 0.0
//	for _, number := range numbers {
//		sum += float64(number)
//	}
//
//	return sum / float64(len(numbers))
//}
//
//// Range возвращает размах числовой последовательности.
//// Напишите код функции
//func Range(numbers []int) int {
//	if len(numbers) == 0 {
//		return 0
//	}
//	min := numbers[0]
//	max := numbers[0]
//	for _, number := range numbers {
//		if number < min {
//			min = number
//		}
//		if number > max {
//			max = number
//		}
//	}
//	return max - min
//}
//
//func main() {
//	lists := [][]int{
//		{},
//		{57},
//		{78, -7},
//		{99, 200, 0},
//		{4, 4, 4, 3},
//		{102, -7, 44, -7, 102},
//		{82, -23, 1, 5, 98, 100},
//		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
//	}
//	averages := []float64{
//		0, 57, 36, 100, 4, 47, 44, 39938,
//	}
//	ranges := []int{
//		0, 0, 85, 200, 1, 109, 123, 80000,
//	}
//
//	for i, list := range lists {
//		if average := math.Round(Average(list)); average != averages[i] {
//			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
//		}
//		if r := Range(list); r != ranges[i] {
//			fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
//		}
//	}
//	fmt.Println("Тестирование завершено")
//}

//package main
//
//import (
//	"errors"
//	"fmt"
//	"strconv"
//)
//
//// создайте новые ошибки
//var (
//	ErrDivisionByZero = errors.New("division by zero")
//	ErrConversion     = errors.New("conversion error")
//)
//
//func div(num1, num2 string) (int, error) {
//	a, err1 := strconv.Atoi(num1)
//	if err1 != nil {
//		return 0, fmt.Errorf("%w: %w", ErrConversion, err1)
//	}
//	b, err2 := strconv.Atoi(num2)
//	if err2 != nil {
//		return 0, fmt.Errorf("%w: %w", ErrConversion, err2)
//	}
//	if b == 0 {
//		return 0, ErrDivisionByZero
//	}
//	return a / b, nil
//}
//
//func main() {
//	res, err := div("5", "0")
//	if err != nil {
//		if errors.Is(err, strconv.ErrSyntax) {
//			fmt.Println("ошибка преобразования строки в целое")
//		} else if errors.Is(err, ErrDivisionByZero) {
//			fmt.Println("деление на 0")
//		}
//	} else {
//		fmt.Println(res)
//	}
//
//	res, err = div("abc", "def")
//	if err != nil {
//		if errors.Is(err, strconv.ErrSyntax) {
//			fmt.Println("ошибка преобразования строки в целое")
//		} else if errors.Is(err, ErrDivisionByZero) {
//			fmt.Println("деление на 0")
//		}
//	} else {
//		fmt.Println(res)
//	}
//}

//package main
//
//import (
//	"log"
//	"os"
//)
//
//func main() {
//	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer flog.Close()
//	log.SetOutput(flog)
//	// создаем логер, который будем использовать
//	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//
//	log.Print("Logging to a file in Go!")
//}

//package main
//
//import (
//	"errors"
//	"fmt"
//	"strconv"
//)
//
//// создаем переменную ошибки тут
//var (
//	ErrConvToFloat       = errors.New("ошибка преобразования строки в число с плавающей точкой")
//	ErrNotPositiveNumber = errors.New("число должно быть больше нуля")
//)
//
//// ниже ваша функция сложения
//func addPositive(num1, num2 string) (float64, error) {
//	divisible, err := strconv.ParseFloat(num1, 64)
//	if err != nil {
//		return 0, ErrConvToFloat
//	}
//	divider, err := strconv.ParseFloat(num2, 64)
//	if err != nil {
//		return 0, ErrConvToFloat
//	}
//	if divider == 0.0 {
//		return 0, ErrNotPositiveNumber
//	}
//	return float64(divisible + divider), nil
//}
//
//func main() {
//	num1 := "5.2"
//	num2 := "2.7"
//
//	result, err := addPositive(num1, num2)
//
//	if err != nil {
//		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
//		fmt.Println(err)
//	} else {
//		fmt.Println("результат сложения", result)
//	}
//
//	num1 = "two"
//	num2 = "five"
//
//	result, err = addPositive(num1, num2)
//	if err != nil {
//		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
//		fmt.Println(err)
//	} else {
//		fmt.Println("результат сложения:", result)
//	}
//
//	num1 = "5.4"
//	num2 = "0.0"
//
//	result, err = addPositive(num1, num2)
//	if err != nil {
//		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
//		fmt.Println(err)
//	} else {
//		fmt.Println("результат сложения:", result)
//	}
//}

//package main
//
//import "fmt"
//
//const (
//	electricity   = 5.88
//	hotWater      = 156.28
//	coldWater     = 32.39
//	waterDrainage = 29.69
//)
//
//func payment(lastCounter, currentCounter int, tariff float64) float64 {
//	return float64(currentCounter-lastCounter) * tariff
//}
//
//func electricityCalc(lastCounter, currentCounter int) float64 {
//	return payment(lastCounter, currentCounter, electricity)
//}
//
//func hotWaterCalc(lastCounter, currentCounter int) float64 {
//	return payment(lastCounter, currentCounter, hotWater)
//}
//
//func coldWaterCalc(lastCounter, currentCounter int) float64 {
//	return payment(lastCounter, currentCounter, coldWater)
//}
//
//func waterDrainageCalc(lastHotCounter, currentHotCounter, lastColdWater, currentColdWater int) float64 {
//	return payment(lastHotCounter, currentHotCounter, waterDrainage) + payment(lastColdWater, currentColdWater, waterDrainage)
//}
//
//// ваша функция для расчета общей суммы ниже
//func f(arg ...float64) float64 {
//	total := 0.0
//	for _, payments := range arg {
//		total += payments
//	}
//	return total
//}
//
//func main() {
//	paymentsElectocity := electricityCalc(13130, 13383)
//
//	fmt.Println("В этом месяце платеж за электроэнергию составил", paymentsElectocity)
//
//	paymentsHotWater := hotWaterCalc(57, 60)
//
//	fmt.Println("В этом месяце платеж за горячую воду составил", paymentsHotWater)
//
//	paymentsColdWater := coldWaterCalc(191, 199)
//
//	fmt.Println("В этом месяце платеж за холодную воду составил", paymentsColdWater)
//
//	paymentWaterDrainage := waterDrainageCalc(57, 60, 191, 199)
//
//	fmt.Println("В этом месяце платеж за водоотведение составил", paymentWaterDrainage)
//
//	// упаковываем все платежи в слайс
//	payments := []float64{paymentsElectocity, paymentsHotWater, paymentsColdWater, paymentWaterDrainage}
//
//	// вызовите функцию для общей суммы ниже, передайте в качестве аргумента слайс payments
//	total := f(payments...)
//
//	fmt.Println("Всего придется заплатить", total)
//}

//package main
//
//import "fmt"
//
//func printFriendsCount(friendsCount int) {
//	if friendsCount == 0 {
//		fmt.Println("У тебя нет друзей")
//	} else if friendsCount == 1 {
//		fmt.Println("У тебя", friendsCount, "друг")
//	} else if friendsCount >= 2 && friendsCount <= 4 {
//		fmt.Println("У тебя", friendsCount, "друга")
//	} else if friendsCount >= 5 && friendsCount < 20 {
//		fmt.Println("У тебя", friendsCount, "друзей")
//	} else {
//		fmt.Println("Ого, сколько у тебя друзей! Целых ", friendsCount)
//	}
//}
//
//func main() {
//	// Ниже закончите цикл, в котором будет вызываться функция printFriendsCount
//	// с аргументом от 0 до 20 включительно
//
//	for i := 0; i <= 20; i++ {
//		printFriendsCount(i)
//
//	}
//
//}

//import "fmt"
//
//func main() {
//	s := make(map[string]int)
//	s["f"] = 1
//	s["kok"] = 90
//	s["f"] = 10
//	fmt.Println(s["f"])
//	_, ok := s["f"]
//	if !ok {
//		fmt.Println("no")
//	} else {
//		fmt.Println("yes")
//	}
//	fmt.Println(len(s))
//}

//package main
//
//import "fmt"
//
//const (
//	CmdAdd   = iota // сложить два числа в стеке (top-1) = (top-1) + top
//	CmdSub          // вычесть (top-1) = (top-1) - top
//	CmdMul          // умножить (top-1) = (top-1) * top
//	CmdDiv          // разделить (top-1) = (top-1) / top
//	CmdPush         // поместить следующее число в стек
//	CmdPop          // убрать число из стека
//	CmdPrint        // печать верхнего элемента стека
//	CmdSave         // сохранить число из стека в ячейку
//	CmdLoad         // переместить число из ячейки в стек
//)
//
//func main() {
//	program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush,
//		-13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul,
//		CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7,
//		CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
//		CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub,
//		CmdPush, 1000, CmdDiv, CmdPrint}
//	stack := make([]int, 0, 100)
//	registers := make(map[rune]int)
//	for i := 0; i < len(program); i++ {
//		cmd := program[i]
//		top := len(stack) - 1
//		switch cmd {
//		case CmdAdd:
//			stack[top-1] = stack[top-1] + stack[top]
//			stack = stack[:top]
//		case CmdSub:
//			stack[top-1] = stack[top-1] - stack[top]
//			stack = stack[:top]
//		case CmdMul:
//			stack[top-1] = stack[top-1] * stack[top]
//			stack = stack[:top]
//		case CmdDiv:
//			if len(stack) < 2 {
//				fmt.Println("no element ")
//			} else {
//				if stack[top] == 0 {
//					fmt.Println("nolzen ")
//				} else {
//					stack[top-1] = stack[top-1] / stack[top]
//					stack = stack[:top]
//				}
//			}
//		case CmdPush:
//			stack = append(stack, program[i+1])
//			i++
//		case CmdPop:
//			stack = stack[:top]
//		case CmdPrint:
//			fmt.Println(stack[top])
//		case CmdSave:
//			registers[rune(program[i+1])] = stack[top]
//			i++
//		case CmdLoad:
//			stack = append(stack, registers[rune(program[i+1])])
//			i++
//		}
//
//	}
//
//}

//package main
//
//import "fmt"
//
//func main() {
//	titles := map[string][]string{
//		"Что делать?":               {"планы", "думы"},
//		"Где отдохнуть в выходные":  {"отдых", "планы"},
//		"Кто виноват?":              {"поиск", "думы", "общество"},
//		"Лучшие рестораны города":   {"еда", "отдых"},
//		"С заботой о народе":        {"думы", "общество"},
//		"Как попасть на дегустацию": {"еда", "планы", "отдых"},
//	}
//	tags := make(map[string][]string)
//	for title, tag := range titles {
//		for _, ta := range tag {
//			tags[ta] = append(tags[ta], title)
//		}
//	}
//	fmt.Println(len(tags["думы"]), len(tags["общество"]), tags["поиск"])
//}

// 1 переоброзование
//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	// прайс-лист выпечки
//	bakery := map[string]int{
//		"Хлеб белый":          43,
//		"Хлеб Дарницкий":      47,
//		"Батон":               52,
//		"Булочка Домашняя":    35,
//		"Хачапури":            63,
//		"Сосиска в тесте":     70,
//		"Беляш":               82,
//		"Растегай":            87,
//		"Самса":               91,
//		"Пирожок с картошкой": 37,
//	}
//	var list = make([]string, 0, len(bakery))
//
//	for s := range bakery {
//		list = append(list, s)
//
//	}
//	sort.Strings(list)
//	for _, j := range list {
//		g := bakery[j]
//		fmt.Println(g, ":", j)
//	}
//}
