package main

//type Person struct {
//	Name        string    `json:"Имя"`
//	Email       string    `json:"Почта"`
//	DateOfBirth time.Time `json:"-"`
//}
//
//func main() {
//	man := Person{
//		Name:        "Алекс",
//		Email:       "alex@yandex.ru",
//		DateOfBirth: time.Now(),
//	}
//	jsMan, err := json.Marshal(man)
//	if err != nil {
//		log.Fatalln("unable marshal to json")
//	}
//	fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Алекс","Почта":"alex@yandex.ru"}
//}
