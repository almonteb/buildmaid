package fileman

import (
	"crypto/tls"
	"github.com/minio/minio-go"
	"net/http"
)

type fileManS3 struct {
	client *minio.Client
	bucket string
}

func NewFileManS3(accessKey, secretKey, bucket, host string) (*fileManS3, error) {
	client, err := minio.New(host, accessKey, secretKey, true)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.SetCustomTransport(tr)

	if err != nil {
		return nil, err
	}
	return &fileManS3{
		client: client,
		bucket: bucket,
	}, nil
}

func (fm *fileManS3) GetDirectories(path string) ([]string, error) {
	doneCh := make(chan struct{})
	defer close(doneCh)

	var dirs []string
	objectCh := fm.client.ListObjects(fm.bucket, path, false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		dirs = append(dirs, object.Key)
	}

	return dirs, nil
}

func (fm *fileManS3) Delete(path string) error {
	return nil
}
