# s3size
## Amazon S3 Bucket size calculator - cli tool to calculate an Amazon S3 bucket size (or any path into a bucket)

### Usage

The following example will calculate the whole bucket size:

```
./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion
```


You can also calculate only a path into the bucket:

```
./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion **--path anyFolder/**
```

Use the above call to calculate only the size of the _anyFolder_ folder into the _yourBucketName_ bucket.


### Output example

```
total files: 54
total folders: 6
size: 730.6M (766053540 bytes)
```