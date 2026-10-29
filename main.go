// package main

// import "fmt"

// type Engine struct {
// 	Power int
// 	Fuel  string
// }

// func (e Engine) Start() string{
// 	return "я запустился р-р-р"
// }

// type Car struct {
// 	Engine Engine
// 	Brand  string
// }



// type Student struct {
// 	Name string
// 	Average int
// }

// func (s *Student) changeAverage(p int){
// 	s.Average = p
// }

// type Group struct {
// 	Name string
// 	Students []*Student
// }



// type Logger struct{
// 	Prefix string
// }

// func (l Logger) Log(msg string){
// 	fmt.Println(l.Prefix, msg)
// }

// type Service struct{
// 	Logger
// 	Name string
// }


// func main() {

// 	s := Service{Logger: Logger{Prefix: "[svc]"}, Name: "auth"}
// 	s.Log("started")
// 	fmt.Println(s.Prefix)

// 	fmt.Println("----------------------------")

// 	s1 := &Student{Name: "Vera", Average: 5}
// 	s2 := &Student{Name: "Mark", Average: 5}

// 	group := Group{
// 		Name: "IS223",
// 		Students: []*Student{s1, s2},
// 	}


// 	group.Students[0].changeAverage(6)
// 	fmt.Println(s1.Average)
	
// 	group := Group{
// 		Name: "Electron",
// 		Students: []*Student{s1},

// 	}




// 	c := Car{
// 		Engine: Engine{Power: 84, Fuel: "gasoline"},
// 		Brand: "Lada",
// 	}

// 	fmt.Println(c.Engine.Start())

// 	c2 := c
// 	c2.Engine.Power = 200
// 	fmt.Println(c.Engine.Power)
// }

// package main

// type Author struct {
// 	FirstName string
// 	LastName  string
// }

// type Book struct {
// 	Title      string
// 	Year       int
// 	BookAuthor Author
// }


// func main() {
// 	books := []Book{}
// 	book1 := Book{
// 		Title: "Война и мир",
// 		Year:  1869,
// 		BookAuthor: Author{
// 			FirstName: "Лев",
// 			LastName:  "Толстой",
// 		},
// 	}
// 	book2 := Book{
// 		Title: "Гуру дизайна",
// 		Year:  2024,
// 		BookAuthor: Author{
// 			FirstName: "Апполон",
// 			LastName:  "Чиколаев",
// 		},
// 	}
// 	book3 := Book{
// 		Title: "Как переносить абьюз программистов. Том 1. Плачем вместе",
// 		Year:  2025,
// 		BookAuthor: Author{
// 			FirstName: "Апполон",
// 			LastName:  "Чиколаев",
// 		},
// 	}
// 	books = append(books, book1, book2, book3)
// 	for _, books 
// }

package main

func (b Book) String() string {
	return fmt.Sprintf(`"%s (%s, %d)"`, b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
     if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не      может получить книгу.", reader.FirstName, reader.LastName)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}


func main{

}