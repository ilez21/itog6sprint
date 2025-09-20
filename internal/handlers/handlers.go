package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	content := string(data)
	result, err := service.AutoConvert(content)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(result)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
