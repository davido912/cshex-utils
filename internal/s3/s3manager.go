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

// func downloadFromS3(cmd *cobra.Command, args []string) error {
// 	s3manager := s3.New()
// 	bs, err := os.ReadFile("input")
// 	if err != nil {
// 		panic(err)
// 	}
// 	splitted := strings.Split(string(bs), ",")
// 	files := make([]s3.ObjectDownloadInput, len(splitted))
// 	for _, s := range splitted {
// 		k := strings.Trim(strings.Replace(s, "'", "", -1), " ")
// 		files = append(files,
// 			s3.ObjectDownloadInput{
// 				Bucket:  "traderepublic-timeline-mqlog",
// 				Key:     k,
// 				DstPath: fmt.Sprintf("./output/%s", path.Base(k)),
// 			},
// 		)
// 	}

// 	for _, f := range files {
// 		fmt.Printf("%+v\n", f)
// 	}
// 	s3manager.DownloadFiles(context.Background(), files...)
// 	return nil

// }

// // MustCli instantiates the CLI with the given entrypoint function passed to it
// func MustCli(rootCmd *cobra.Command) {
// 	if err := rootCmd.Execute(); err != nil {
// 		logger.Fatal().Err(err)
// 	}
// }
