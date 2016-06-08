package main

import (
	"fmt"
	"github.com/pivotal-golang/bytefmt"
	"gopkg.in/amz.v3/aws"
	"gopkg.in/amz.v3/s3"
)

type s3Files struct {
	totalFolders int64
	totalFiles   int64
	size         uint64
}

func getBucketList(bucketname string, key string, secret string, region aws.Region, prefix string, delim string, marker string, max int) (err error) {

	AWSAuth := aws.Auth{
		AccessKey: key,
		SecretKey: secret,
	}

	connection := s3.New(AWSAuth, region)
	bucket, err := connection.Bucket(bucketname)
	if err != nil {
		return err
	}

	lr, err := bucket.List(prefix, delim, marker, max)
	if err != nil {
		return err
	}

	more := true
	lastKey := ""

	s3ftotal := &s3Files{}

	for more {
		lastKey = processBucketList(lr, s3ftotal)
		if lr.IsTruncated {

			lr, err = bucket.List(prefix, delim, lastKey, max)
			if err != nil {
				return err
			}
		}
		more = lr.IsTruncated
	}

	fmt.Printf("\n\ntotal files: %v\ntotal folders: %v\nsize: %v (%v bytes)\n\n", s3ftotal.totalFiles, s3ftotal.totalFolders, bytefmt.ByteSize(s3ftotal.size), s3ftotal.size)

	return nil
}

func processBucketList(lr *s3.ListResp, s3f *s3Files) (lastKey string) {

	for _, v := range lr.Contents {
		if v.Size == 0 && string(v.Key[len(v.Key)-1]) == "/" {
			s3f.totalFolders++
		} else {
			s3f.totalFiles++
			s3f.size = s3f.size + uint64(v.Size)
			lastKey = v.Key
		}
	}

	return
}
