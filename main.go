package main

import (
	"context"

	"github.com/davido912/cshex-utils/internal/s3"
)

func main() {
	s3manager := s3.New()
	s3manager.DownloadFiles(
		context.Background(),
		
	)
}
