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

	const (
		bucketName = "sample-bucket"
		fileKey    = "some_file.txt"
	)
	// TODO: 別ファイルに切り出す、あるいはCLIにする。
	// writeObject(client, bucketName, fileKey, nil)
	// if err != nil {
	// 	log.Fatalf("failed to write object: %v", err)
	// }
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

func genReadBucketPath(bucketName string) string {
	return bucketPathPrefix + bucketName
}

func genReadObjectPath(objectName string) string {
	return objectPathPrefix + objectName
}

func downloadFile(client *storage.Client, bucketName, fileKey string) ([]byte, error) {
	// NOTE: fsouza/fake-gcs-serverのルータが期待するパスに生成する。
	// ref. https://github.com/fsouza/fake-gcs-server/blob/main/fakestorage/server.go
	reader, err := client.Bucket(genReadBucketPath(bucketName)).Object(genReadObjectPath(fileKey)).NewReader(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to create a reader: %v", err)
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

func writeObject(client *storage.Client, bucketName, fileKey string, data []byte) error {
	/*
		NOTE: fsouza/fake-gcs-serverのルータが期待するパスに生成する。
		ref. https://github.com/fsouza/fake-gcs-server/blob/main/fakestorage/server.go#L191

		コンテナ内部では、/storage/<bucket_name>/に書き込まれる。
		/ # ls storage/sample-bucket/
		some_file.txt    some_file_2.txt
	*/
	writer := client.Bucket(bucketName).Object(fileKey).NewWriter(context.TODO())
	defer writer.Close()
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}
