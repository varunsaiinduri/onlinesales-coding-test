package main
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
    "os"
)
func main() {
    awsAccessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
    awsSecretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")
    awsBucketRegion := os.Getenv("S3_BUCKET_REGION")
    awsFilePath := os.Getenv("S3_FILE_PATH")

    bucket := "onlinesales-coding-test"
    //item := "test.zip"
    file, err := os.Create(awsFilePath)
    if err != nil {
        exitErrorf("Unable to open file %q, %v", err)
    }

    defer file.Close()

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String(awsBucketRegion,),
                Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
                },)

    downloader := s3manager.NewDownloader(sess)
    // sess := session.New()
   //  downloader := s3manager.NewDownloader(sess, func(d *s3manager.Uploader) {
   //  d.PartSize = 64 * 1024 * 1024
    // })

    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String(bucket),
            Key:    aws.String(awsFilePath),
        })
    if err != nil {
        exitErrorf("Unable to download awsFilePath %q, %v", awsFilePath, err)
    }

    fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}