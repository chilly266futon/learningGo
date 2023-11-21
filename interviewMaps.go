package main

//func main() {
//	arr := []int{25, 26, 27, 28, 29, 30}
//	k := 52
//	fmt.Println(find(arr, k))
//}
//func find(arr []int, k int) []int {
//	// Создадим пустую map
//	m := make(map[int]int)
//	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
//	for i, a := range arr {
//		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
//			return []int{i, j}
//		}
//		// если искомого значения нет, то добавляем текущий индекс и значение в map
//		m[a] = i
//	}
//	// не нашли пары подходящих чисел
//	return nil
//}

// как можно заметить, алгоритм пройдётся по массиву всего один раз
// если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений
