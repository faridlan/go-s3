package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

type Upload struct {
	ImageUrl string `json:"image_url,omitempty"`
}

func uploadS3(image string) Upload {
	LoadEnv()
	awsAccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	fmt.Println("My access key ID is ", awsAccessKeyID)

	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	s := RandStringRunes(10)
	f, err := os.Open(image)
	if err != nil {
		panic(err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(GetEnvWithKey("BUCKET_NAME")),
		Key:         aws.String("test/" + s + ".png"),
		Body:        f,
		ContentType: aws.String("image/png"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		panic(err)
	}

	return Upload{
		ImageUrl: aws.StringValue(&result.Location),
	}
	// fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func RequestFromStruct(r *http.Request, result any) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(result)
	if err != nil {
		panic(err)
	}
}

func ToJson(w http.ResponseWriter, response any) {

	w.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	if err != nil {
		panic(err)
	}
}

func GetUser() User {

	x := User{
		Name: "Faridlan",
		Age:  23,
	}

	return x

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(err)
	}

	f, _, err := r.FormFile("testImage")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	x := User{}
	RequestFromStruct(r, &x)

	ToJson(w, x)
}

func main() {

	sm := http.NewServeMux()
	sm.HandleFunc("/", HelloWorld)

	server := http.Server{
		Addr:    ":8080",
		Handler: sm,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
