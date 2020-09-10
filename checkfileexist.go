package main
import (
        "fmt"
        "time"
        "os"
        "github.com/AdRoll/goamz/aws"
        "github.com/AdRoll/goamz/s3"
)
func main() {
        awsAccessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
        awsSecretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")
        awsBucketRegion := os.Getenv("S3_BUCKET_REGION")
        awsFilePath := os.Getenv("S3_FILE_PATH")
//        fmt.Println(awsAccessKeyID)
        now := time.Now()
        threeDays := time.Hour * 24 * 3
        diff := now.Add(threeDays)
        auth, _ := aws.GetAuth(awsAccessKeyID, awsSecretAccessKey, "", diff)

        region := aws.GetRegion(awsBucketRegion) //example us-east-1

        S3 := s3.New(auth, region) //creates new S3
        bucket := S3.Bucket("onlinesales-coding-test") //example myBucket
        exists, _ := bucket.Exists(awsFilePath) //example /myPic.jpg
        //return true/false, error
        fmt.Println(exists)
}