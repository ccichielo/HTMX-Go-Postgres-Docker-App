package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"example.com/lib"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	lib.TestDbConnection()
	imgByte := lib.GetImageBytesById(1)

	/*img, err := jpeg.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}*/

	// Generate HTML using a template
	htmlTemplate := `<html>
        <body>
            <img src="data:image/{{.ImageType}};base64,{{.ImageData}}" alt="Decoded Image">
        </body>
    </html>`

	data := struct {
		ImageType string
		ImageData string
	}{
		ImageType: "jpeg", // or "png"
		ImageData: base64.StdEncoding.EncodeToString(imgByte),
	}

	tmpl, err := template.New("html").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing HTML template:", err)
	}
}

func main() {
	http.HandleFunc("/", WebHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
