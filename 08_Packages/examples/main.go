package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a","b"},"b"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))

	arr := []byte("test")
	str := string([]byte {'t','e','s','t'})

	fmt.Println(str, "in binary is", arr)

	// Input/Output
	var buf bytes.Buffer
	buf.Write([]byte("test"))

	// Files and Folders
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	strBs := string(bs)
	fmt.Println(strBs)

	shorterBS
}