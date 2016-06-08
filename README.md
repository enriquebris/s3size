# s3size . Amazon S3 Bucket size calculator
## cli tool to calculate an Amazon S3 bucket size (or any path inside a S3 bucket)


### Installation

Make sure you have a working Go environment in order to get the source code and compile it to your architecture.

```
go get github.com/enriquebris/s3size
```

### Supported platforms

s3size is currently tested on:
 1. Mac OS X 10.8.5
 2. Windows 10 (64-bit)
 3. Ubuntu Linux 14.04 (64-bit)

 **Note:** In order to get better performance the suggestion is: run it directly inside an AWS box.

### Usage

The following example will calculate the whole bucket size:

```
./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion
```


You can also calculate only a path inside the bucket:

```
./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion --path anyFolder/
```



### Output example

```
total files: 54
total folders: 6
size: 730.6M (766053540 bytes)
```