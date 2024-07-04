package main

import (
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	listLength := 0
	cursor := head

	for cursor != nil {
		listLength++
		cursor = cursor.Next
	}

	cursor = head

	for i := 0; i < listLength/2; i++ {
		cursor = cursor.Next
	}

	return cursor
}

func canConstructWithArr(ransomNote string, magazine string) bool {
	magazines := strings.Split(magazine, "")
	ransomNotes := strings.Split(ransomNote, "")
	finalLen := 0

	for i := 0; i < len(ransomNotes); i++ {
		for j := 0; j < len(magazines); j++ {
			if magazines[j] == ransomNotes[i] {
				magazines[j] = ""
				finalLen++
				break
			}
		}
	}

	return finalLen == len(ransomNotes)
}

func canConstruct(ransomNote string, magazine string) bool {
	resultMap := make(map[string]int)

	for _, char := range magazine {
		resultMap[string(char)]++
	}

	for _, char := range ransomNote {
		key := string(char)

		_, isOk := resultMap[key]
		if isOk {
			if resultMap[key] == 1 {
				delete(resultMap, key)
			} else {
				resultMap[key]--
			}
		}
	}

	return len(resultMap) == 0
}

func main() {
	nodes := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	middleNode(nodes)
}
