package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var System = runtime.GOOS

type SetCursorPosition struct {
	Top  int
	Left int
}

var CursorPosition *SetCursorPosition = &SetCursorPosition{}

func RandomString(Length int) string {
	Random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	Letters := []rune("abcdefghijklmnopqrstuwxyz123456789_")
	B := make([]rune, Length)
	for i := range B {
		B[i] = Letters[Random.Intn(len(Letters))]
	}
	return string(B)
}

func RandomStringUpper(Length int) string {
	Random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	Letters := []rune("QWERTYUIOPASDFGHJKLZXCVBNM1234567890")
	B := make([]rune, Length)
	for i := range B {
		B[i] = Letters[Random.Intn(len(Letters))]
	}
	return string(B)
}

func RandomStringNumber(Length int) string {
	Random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	Numbers := []rune("1234567890")
	B := make([]rune, Length)
	for i := range B {
		B[i] = Numbers[Random.Intn(len(Numbers))]
	}
	return string(B)
}

func Remove(index *int, Array []string) []string {
	return append(Array[:*index], Array[*index+1:]...)
}

func CreateFile(Path, Text string) os.File {
	File, _ := os.OpenFile(Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer File.Close()
	File.WriteString(Text)
	return *File
}

func ReWriteList(Path string, Array *[]string) {
	os.Remove(Path)
	for _, S := range *Array {
		CreateFile(Path, fmt.Sprintf("%s\r\n", S))
	}
}

func CreateFileOnly(Path string) os.File {
	File, err := os.Create(Path)
	if err != nil {
		log.Fatalln(File)
	}
	return *File
}

func ClearConsole() string {
	if System == "windows" {
		Cmd := exec.Command("cmd", "/c", "cls")
		Cmd.Stdout = os.Stdout
		Cmd.Run()
	} else {
		Cmd := exec.Command("clear")
		Cmd.Stdout = os.Stdout
		Cmd.Run()
	}
	return System
}

func RemoveFromFile(L string, List *[]string, Dir string) {
	var Temp string
	for F, LL := range *List {
		if LL != L {
			if F == len(*List)-1 {
				Temp += LL
			} else {
				Temp += LL + "\r\n"
			}
		}
	}
	O := os.Remove(Dir)
	if O != nil {
		//fmt.Print(O)
	}
	Er := ioutil.WriteFile("./"+Dir, []byte(Temp), os.ModePerm)
	if Er != nil {
		//fmt.Println("Holy fock")
	}
	LLL, Erra := ioutil.ReadFile(Dir)
	if Erra != nil {
		//fmt.Println(Erra)
	}
	*List = strings.Split(string(LLL), "\r\n")
}

func LoadFile(Name string, Path string) ([]string, error) {
	File, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	var LIST []string
	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		ReplaceStrings := strings.Join(strings.Fields(Scanner.Text()), "")
		ReplaceStrings = strings.Replace(ReplaceStrings, "\n", "", -1)
		ReplaceStrings = strings.Replace(ReplaceStrings, " ", "", -1)
		if ReplaceStrings != "" {
			LIST = append(LIST, ReplaceStrings)
		}
	}
	File.Close()
	return LIST, err
}
