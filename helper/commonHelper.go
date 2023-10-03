package helper

import "os"

func GetFilePath(params string) string {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fileDomain := os.Getenv("FILE_URL")

	return fileDomain +"/"+ params
}