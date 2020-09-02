package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	httpClient := &http.Client{Transport: transCfg}
	client, err := storage.NewClient(context.TODO(), option.WithEndpoint("https://storage.gcs.127.0.0.1.nip.io:4443/storage/v1/"), option.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatal(err)
	}
	buckets, err := list(client, "test")
	if err != nil {
		log.Fatalf("failed to list: %v", err)
	}
	fmt.Printf("buckets: %+v\n", buckets)

	const (
		bucketName = "sample-bucket"
		fileKey    = "some_file.txt"
	)
	data, err := downloadFile(client, bucketName, fileKey)
	if err != nil {
		log.Fatalf("failed to download file: %v", err)
	}
	fmt.Printf("contents of %s/%s: %s\n", bucketName, fileKey, data)
}

func list(client *storage.Client, projectID string) ([]string, error) {
	var buckets []string
	it := client.Buckets(context.TODO(), projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets, nil
}

const (
	bucketPathPrefix = "storage/v1/b/"
	objectPathPrefix = "o/"
)

func genBucketPath(bucketName string) string {
	return bucketPathPrefix + bucketName
}

func genObjectPath(objectName string) string {
	return objectPathPrefix + objectName
}

func downloadFile(client *storage.Client, bucketName, fileKey string) ([]byte, error) {
	// NOTE: fsouza/fake-gcs-serverのルータが期待するパスに生成する。
	// ref. https://github.com/fsouza/fake-gcs-server/blob/main/fakestorage/server.go
	reader, err := client.Bucket(genBucketPath(bucketName)).Object(genObjectPath(fileKey)).NewReader(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to create a reader: %v", err)
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}
