package main

//func main() {
//	input := []string{
//		"cat",
//		"dog",
//		"bird",
//		"dog",
//		"parrot",
//		"cat",
//	}
//	fmt.Println(input)
//
//	fmt.Println(RemoveDuplicates(input))
//
//}
//
//func RemoveDuplicates(input []string) []string {
//	output := make([]string, 0)
//	inputSet := make(map[string]struct{}, len(input))
//	for _, v := range input {
//		if _, ok := inputSet[v]; !ok {
//			output = append(output, v)
//		}
//		inputSet[v] = struct{}{}
//	}
//	return output
//}
