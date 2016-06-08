# s3size
Amazon S3 Bucket size calculator

cli tool to calculate an Amazon S3 bucket size (or any path into a bucket)

#### Usage

The following example will calculate the whole bucket size:

_./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion_


You can also calculate only a path into the bucket:

_./s3size --key yourKey --secret yourSecret --bucket yourBucketName --region anyAWSregion **--path anyFolder/**_

The above call will calculate only the size of the anyFolder folder into the yourBucketName bucket.


#### Output example

total files: 54
total folders: 6
size: 730.6M (766053540 bytes)