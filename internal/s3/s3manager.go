package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/davido912/cshex-utils/internal/filemanager"
)

type Manager struct {
	downloader  *s3manager.Downloader
	fileManager *filemanager.FileManager
}

type ObjectDownloadInput struct {
	Bucket  string
	Key     string
	DstPath string
}

func (m *Manager) DownloadFiles(ctx context.Context, objectDownloadInputs ...ObjectDownloadInput) {
	for _, objectDownloadInput := range objectDownloadInputs {
		if _, err := m.download(objectDownloadInput, ctx); err != nil {
			fmt.Printf("failed to download file, err: %s\n", err)
		}
	}
}

func (m *Manager) download(objectDownloadInput ObjectDownloadInput, ctx context.Context) (int64, error) {
	return m.downloader.DownloadWithContext(
		ctx,
		m.fileManager.MustCreate(objectDownloadInput.DstPath),
		&s3.GetObjectInput{
			Bucket: &objectDownloadInput.Bucket,
			Key:    &objectDownloadInput.Key,
		},
	)
}

func New() *Manager {
	sess := session.Must(session.NewSession())
	return &Manager{
		downloader:  s3manager.NewDownloader(sess),
		fileManager: filemanager.New(),
	}
}
