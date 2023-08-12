package main 

import (
    "fmt"
    "YadroTest/logger"
)

func main() {
    inputDir  := "/home/rus/GoFolder/YadroTest/input/"
    outputDir := "/home/rus/GoFolder/YadroTest/output/"
    path1 := inputDir  + "test6.txt"
    path2 := outputDir + "test6.txt"
    fmt.Println("Docker work")
    logger.Logger(path1, path2)
}
