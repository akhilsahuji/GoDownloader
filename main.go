package main

import (
	"flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    // specify the URL of the file to download
    urlFlag := flag.String("u", "", "URL of the file to download")
    folderFlag := flag.String("f", "", "path to the folder where the file should be saved")



	// parse the command-line flags
    flag.Parse()

    // make sure the URL and folder path are provided
    if *urlFlag == "" {
        fmt.Println("Please provide the URL using the -u flag")
        return
    }

	// Get the current working directory
    folderPath, err := os.Getwd()
    if err != nil {
        // Handle error
        fmt.Println("Error:", err)
        return
    }


    // extract the file name from the URL
    fileName := filepath.Base(url)

    // create the file to save the download
    filePath := filepath.Join(folderPath, fileName)
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // download the file
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer response.Body.Close()

    // copy the downloaded content to the file
    _, err = io.Copy(file, response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Download completed successfully.")
}
