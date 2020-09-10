#!/bin/sh
outputFile="$3"
export S3_ACCESS_KEY_ID="$1"
export S3_SECRET_ACCESS_KEY="$2"
export S3_BUCKET_REGION="$4"
export S3_FILE_PATH="$3"
#echo $S3_ACCESS_KEY_ID $S3_SECRET_ACCESS_KEY $S3_BUCKET_REGION $S3_FILE_PATH
checkoutput=`go run checkfileexist.go`
#checkoutput=true
echo $checkoutput
if [ $checkoutput == true ]; then
if [ -f "$outputFile" ]; then
    echo "$outputFile exists in directory"
else
    echo "$outputFile does not exist in directory"
    go run download.go
fi
else
  echo "$outputFile does not exist in bucket"
fi
