package main

import "fmt"

type Displayable interface {
	Display()
}

type Employee struct {
	Surname  string
	Name     string
	Age      int
	Position string
	Salary   int
}

func (employee Employee) Display() {
	fmt.Printf(
		"%s %s, %s, возраст: %d, зарплата: %d руб.\n",
		employee.Surname,
		employee.Name,
		employee.Position,
		employee.Age,
		employee.Salary,
	)
}

func AddEmployees(employees []Employee, newEmployees []Employee) []Employee {
	return append(employees, newEmployees...)
}

func FilterEmployees(employees []Employee, minAge int, minSalary int) []Employee {
	var filtered []Employee

	for _, employee := range employees {
		if employee.Age >= minAge && employee.Salary >= minSalary {
			filtered = append(filtered, employee)
		}
	}

	return filtered
}

func DisplayEmployees(employees []Employee) {
	for _, employee := range employees {
		var displayable Displayable = employee
		displayable.Display()
	}
}

func main() {
	employees := []Employee{
		{
			Surname:  "Patrushev",
			Name:     "Alexey",
			Age:      30,
			Position: "Developer",
			Salary:   350000,
		},
		{
			Surname:  "Ivanov",
			Name:     "Ivan",
			Age:      40,
			Position: "CTO",
			Salary:   500000,
		},
		{
			Surname:  "Andrey",
			Name:     "Patrushev",
			Age:      18,
			Position: "Intern",
			Salary:   50000,
		},
		{
			Surname:  "Petr",
			Name:     "Petrov",
			Age:      25,
			Position: "QA",
			Salary:   170000,
		},
	}

	newEmployees := []Employee{
		{
			Surname:  "Torvalds",
			Name:     "Linus",
			Age:      56,
			Position: "God",
			Salary:   12000000,
		},
		{
			Surname:  "OpenAI",
			Name:     "GPT",
			Age:      4,
			Position: "Helper",
			Salary:   0,
		},
	}

	fmt.Println("Старый список сотрудников:")
	DisplayEmployees(employees)

	fmt.Println("\nСписок кандидатов у HR:")
	DisplayEmployees(newEmployees)

	employees = AddEmployees(employees, newEmployees)

	fmt.Println("\nАктуальный список сотрудников:")
	DisplayEmployees(employees)

	var minAge int

	fmt.Println("\nВведите минимальный возраст:")
	_, err := fmt.Scanln(&minAge)

	if err != nil || minAge < 0 {
		fmt.Println("Ошибка: некорректный возраст")
		return
	}

	var minSalary int

	fmt.Println("Введите минимальную зарплату:")
	_, err = fmt.Scanln(&minSalary)

	if err != nil || minSalary < 0 {
		fmt.Println("Ошибка: некорректная зарплата")
		return
	}

	filteredEmployees := FilterEmployees(
		employees,
		minAge,
		minSalary,
	)

	fmt.Printf(
		"\nСотрудники возрастом от %d лет и зарплатой от %d рублей:\n",
		minAge,
		minSalary,
	)

	if len(filteredEmployees) == 0 {
		fmt.Println("Подходящих сотрудников не найдено")
		return
	}

	DisplayEmployees(filteredEmployees)
}
