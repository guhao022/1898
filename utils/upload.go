package utils

import (
	"os"
	"strings"
	"net/http"
	"io"
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

func Upload(r *http.Request) (string, error) {

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")

		if err != nil {
			return "", err
		}

		defer file.Close()

		// Create ./tmp if it doesn't exist
		if _, err := os.Stat("./tmp"); err != nil {
			if os.IsNotExist(err) {
				os.Mkdir("./tmp", 0777)
			}
		}

		hashed := SHA1(handler.Filename)
		hashed_base := hashed[0:8]

		extensions := strings.Split(handler.Filename, ".")
		extension := extensions[len(extensions)-1]

		hashed_filename := hashed_base + "." + extension

		f, err := os.OpenFile("./tmp/" + hashed_filename, os.O_WRONLY | os.O_CREATE, 0666)

		if err != nil {
			return "", err
		}

		if (!is_valid_image(f)) {
			return "", errors.New("Invalid image")
		}

		defer f.Close()
		io.Copy(f, file)

		return hashed_filename, nil
	}

	return "", errors.New("transmit the request must using POST")
}

func SHA1(s string) (string) {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func is_valid_image(file *os.File) (bool) {
	return true
}
