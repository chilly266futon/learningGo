package main

//
//import (
//	"fmt"
//	"os"
//	"path/filepath"
//	"strings"
//)
//
//func main() {
//	//PrintAllFilesWithFilter(".", "task")
//	//PrintFilesWithFuncFilter(".", containsDot)
//	PrintFilesWithFuncFilter(".", containsTask)
//}
//
//func PrintAllFilesWithFilter(path string, filter string) {
//	// получаем список всех элементов в папке (и файлов, и директорий)
//	files, err := os.ReadDir(path)
//	if err != nil {
//		fmt.Println("unable to get list of files", err)
//		return
//	}
//	//  проходим по списку
//	for _, f := range files {
//		// получаем имя элемента
//		// filepath.Join — функция, которая собирает путь к элементу с разделителями
//		filename := filepath.Join(path, f.Name())
//		// печатаем имя элемента, если путь к нему содержит filter
//		if strings.Contains(filename, filter) {
//			fmt.Println(filename)
//		}
//		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
//		if f.IsDir() {
//			PrintAllFilesWithFilter(filename, filter)
//		}
//	}
//}
//
//func containsDot(s string) bool {
//	return strings.Contains(s, ".")
//}
//
//func containsTask(s string) bool {
//	return strings.Contains(s, "task")
//}
//
//func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
//	// получаем список всех элементов в папке (и файлов, и директорий)
//	files, err := os.ReadDir(path)
//	if err != nil {
//		fmt.Println("unable to get list of files", err)
//		return
//	}
//	//  проNavig по списку
//	for _, f := range files {
//		// получаем имя элемента
//		// filepath.Join — функция, которая собирает путь к элементу с разделителями
//		filename := filepath.Join(path, f.Name())
//		// печатаем имя элемента, если путь к нему содержит filter
//		if predicate(filename) {
//			fmt.Println(filename)
//		}
//		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
//		if f.IsDir() {
//			PrintFilesWithFuncFilter(filename, predicate)
//		}
//	}
//}
