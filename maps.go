package main

//func main() {
//	prices := make(map[string]int)
//	prices["хлеб"] = 50
//	prices["молоко"] = 100
//	prices["масло"] = 200
//	prices["колбаса"] = 500
//	prices["соль"] = 20
//	prices["огурцы"] = 200
//	prices["cыр"] = 600
//	prices["ветчина"] = 700
//	prices["буженина"] = 900
//	prices["помидоры"] = 250
//	prices["рыба"] = 300
//	prices["хамон"] = 1500
//
//	fmt.Println("Деликатесы: ")
//	for k, v := range prices {
//		if v > 500 {
//			fmt.Println(k)
//		}
//	}
//
//	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
//	total := 0
//	//for i := range order {
//	//	for k, v := range prices {
//	//		if order[i] == k {
//	//			total += v
//	//		}
//	//	}
//	//}
//
//	for _, v := range order {
//		total += prices[v]
//	}
//	fmt.Print("Стоимость заказа: ", total)
//}
