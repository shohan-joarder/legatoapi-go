package helper

import "os"

func GetFilePath(params string) string {

	fileDomain := os.Getenv("FILE_URL")

	return fileDomain + "/" + params
}
