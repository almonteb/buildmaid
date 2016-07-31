package fileman

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
	"github.com/minio/minio-go"
	"net/http"
	"strings"
)

type FileManS3 struct {
	client *minio.Client
	bucket string
}

func NewFileManS3(accessKey, secretKey, bucket, host string) (*FileManS3, error) {
	client, err := minio.New(host, accessKey, secretKey, true)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.SetCustomTransport(tr)

	if err != nil {
		return nil, err
	}
	return &FileManS3{
		client: client,
		bucket: bucket,
	}, nil
}

func (fm *FileManS3) GetBuilds(root string) ([]Build, error) {
	var builds []Build
	err := fm.forEachObject(root, func(o minio.ObjectInfo) error {
		splitPath := strings.Split(o.Key, "/")
		builds = append(builds, Build{Name: splitPath[len(splitPath)-2], Path: o.Key})
		return nil
	})
	return builds, err
}

func (fm *FileManS3) Delete(build Build) error {
	return fm.recursiveDelete(build.Path)
}

func (fm *FileManS3) recursiveDelete(path string) error {
	err := fm.forEachObject(path, func(o minio.ObjectInfo) error {
		if isDirectory(o.Key) {
			if err := fm.recursiveDelete(o.Key); err != nil {
				return err
			}
		} else {
			if err := fm.client.RemoveObject(fm.bucket, o.Key); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return fm.client.RemoveObject(fm.bucket, path)
}

func (fm *FileManS3) forEachObject(path string, f func(minio.ObjectInfo) error) error {
	doneCh := make(chan struct{})
	defer close(doneCh)

	p := ensureTrailingSlash(path)
	objectCh := fm.client.ListObjects(fm.bucket, p, false, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return object.Err
		}
		if object.Key == p {
			continue
		}
		log.Debugf("Object: %+v", object)
		if err := f(object); err != nil {
			return err
		}
	}
	return nil
}

func ensureTrailingSlash(str string) string {
	if !isDirectory(str) {
		str += "/"
	}
	return str
}

func isDirectory(p string) bool {
	return p[len(p)-1] == '/'
}
