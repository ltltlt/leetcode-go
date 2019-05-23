package main

import "fmt"

func getNext(path string) string {
	// abc/cde => abc/
	// abc	=> abc
	var i int
	for i = 0; i < len(path) && path[i] != '/'; i++ {
	}
	if i == len(path) {
		i -= 1
	}
	return path[:i+1]
}
func simplifyPath(path string) string {
	if len(path) == 0 {
		return path
	}
	var indexes = []int{0} // index of /
	newPath := "/"
	var next string
Loop:
	for i := 0; i < len(path); i += len(next) {
		next = getNext(path[i:])
		switch next {
		case "":
			break Loop
		case "/", ".", "./":
		case "../", "..":
			if len(indexes) > 1 {
				indexes = indexes[:len(indexes)-1]
			}
			newPath = newPath[:indexes[len(indexes)-1]+1]
		default:
			newPath += next
			indexes = append(indexes, len(newPath)-1)
		}
	}
	if length := len(newPath); length > 1 && newPath[length-1] == '/' {
		newPath = newPath[:length-1]
	}
	return newPath
}
