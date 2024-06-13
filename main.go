package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

func replaceStringInFile(wg *sync.WaitGroup, filePath *string, oldPattern *string, newStr *string) error {
    defer wg.Done()
    // Open the file for reading and writing
    file, err := os.OpenFile(*filePath, os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    var lines []string

    // Compile the regular expression pattern
    regex := regexp.MustCompile(*oldPattern)

    // Read the file line by line
    for scanner.Scan() {
        line := scanner.Text()

        // Replace using regular expression
        newLine := regex.ReplaceAllString(line, *newStr)
        lines = append(lines, newLine)
    }

    // Write the modified content back to the file
    if err := scanner.Err(); err != nil {
        return err
    }

    // Truncate the file to remove old content
    if err := file.Truncate(0); err != nil {
        return err
    }

    // Move the file cursor to the beginning
    if _, err := file.Seek(0, 0); err != nil {
        return err
    }

    // Write the modified lines back to the file
    writer := bufio.NewWriter(file)
    for _, line := range lines {
        _, err := fmt.Fprintln(writer, line)
        if err != nil {
            return err
        }
    }

    // Flush and close the writer
    if err := writer.Flush(); err != nil {
        return err
    }

    return nil
}

func main() {
    fmt.Println("Welcome to string-changer!")
    if len(os.Args) <= 2 {
        fmt.Println(`Args required! e.g. : string-changer.exe case C:\dir`)
        os.Exit(0)
    }
    db := os.Args[1]
    projectDir := os.Args[2]
    localrgx := `.*=(2|4)`
    case1 := "change with this"
    case2 := "change with that"
    var wg sync.WaitGroup

    files := []string{
        `Dir\that\contains\file.json`,
    }
    for _, file := range files {
        filePath := projectDir + `\` + file
        wg.Add(1)
        switch db {
        case "case1":
            go replaceStringInFile(&wg, &filePath, &localrgx, &case1)
        case "case2":
            go replaceStringInFile(&wg, &filePath, &localrgx, &case2)
        default:
            fmt.Println("Invalid case!")
        }
    }
    wg.Wait()
    fmt.Println("String replaced successfully.")
}
