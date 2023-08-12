package logger 

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "testing"
)

var loggerTests = []string {
    "test1.txt",
    "test2.txt",
    "test3.txt",
    "test4.txt",
    "test5.txt",
    "test6.txt",
    "test8.txt",
    "test9.txt",
    "test10.txt",
}

func reader(path string) string {
    file, _ := os.Open(path)

    scanner := bufio.NewScanner(file)

    str := ""
    for scanner.Scan() {
        str += scanner.Text()
    }
   
    return strings.TrimRight(str, "\n")
}

func TestLogger(t *testing.T) {
    inputDirectory  := "/home/rus/GoFolder/YadroTest/input/" 
    outputDirectory := "/home/rus/GoFolder/YadroTest/output/"
    expectDirectory := "/home/rus/GoFolder/YadroTest/expect/"

    for i, val := range loggerTests {
        Logger(inputDirectory + val, outputDirectory + val)

        output   := reader(outputDirectory + val)
        expected := reader(expectDirectory + val)

        if output != expected {
            t.Error("Не пройден", strconv.Itoa(i + 1) + "-й тест (" + val + ")")

        } else {
            fmt.Println("Пройден", strconv.Itoa(i + 1) + "-й тест (" + val + ")")
        }
    }
}
