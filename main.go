package main

import (
	"bufio"
	"fmt"
	"github.com/mind1949/googletrans"
	"golang.org/x/text/language"
	"io"
	"os"
	"strings"
)

func Translate(input string) string {
	params := googletrans.TranslateParams{
		Src:  "auto",
		Dest: language.SimplifiedChinese.String(),
		Text: input,
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
	return translated.Text
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	config_file, err := os.Open("from.ini")
	if err != nil {
		panic(err)
	}
	defer config_file.Close()
	rd := bufio.NewReader(config_file)

	var filename = "./to.ini"
	var f *os.File
	f, _ = os.OpenFile(filename, os.O_CREATE, 0666)

	for {
		line, err := rd.ReadString('\n')
		if err == nil || (err != nil && io.EOF == err) {
			if line != "" && strings.Contains(line, "=") {
				first_index := strings.Index(line, "\"")
				end_index := strings.LastIndex(line, "\"")
				if first_index > 0 && end_index > first_index {
					fmt.Println(line[first_index+1 : end_index])
					trans := Translate(line[first_index+1 : end_index])
					fmt.Println(trans)
					io.WriteString(f, line[0:first_index+1]+trans+line[end_index:])
				} else {
					io.WriteString(f, line)
				}
			} else {
				io.WriteString(f, line)
			}
			if io.EOF == err {
				break
			}
		} else {
			break
		}
	}
}
