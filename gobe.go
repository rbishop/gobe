package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "path/filepath"
)

func baseFileName(fileName string) string {
  extension := filepath.Ext(fileName)
  endOfFileNameIndex := len(fileName) - len(extension)
  return fileName[0:endOfFileNameIndex]
}

func main() {

  if len(os.Args) > 1 {
    fileName := os.Args[1]
    programName := "./" + baseFileName(fileName)
    programArgs := os.Args[2:]

    exec.Command("go", "build", fileName).Run()
    out, err := exec.Command(programName, programArgs...).Output()
    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("%s", out)

    os.Remove(programName)
  } else {
    fmt.Println("Please specify the file you would like to build and execute")
  }

}
