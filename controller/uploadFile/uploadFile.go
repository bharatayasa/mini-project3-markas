package uploadfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan path file gambar: ")
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	path = strings.TrimSpace(path)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(file.Name()))

	dst, err := os.Create("./images/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
