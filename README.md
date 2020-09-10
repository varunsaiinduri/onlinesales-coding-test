# onlinesales-coding-test
this project used for download file from s3 bucket with inputs as access key, secret key, filename, region.
you can find 3 files in the root folder of repository.
1) download.sh ---> this script is used for exporting environment variables and calling below scripts which is based on go lang programming with logic.
2) checkfileexist.go ----> this go program is used check whether input file is residing in s3 bucket or not, it will give response as true when file exist and false when file does not exist in the s3 bucket.it will take all inputs via environment variables.
3) download.go ----> this program is used to download file from s3 bucket in to local directory. it will give response as downloaded with filename and file size. this program will consume all inputs from environment variables which we will give thorugh shell arguments.

command to run:sh download.sh "<accesskey>" "<secretkey>" "<filename>" "<region>"
