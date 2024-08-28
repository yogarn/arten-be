package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/yogarn/arten/model/transcribe"
	"google.golang.org/grpc"
)

type ITranscribeService interface {
	TranscribeEnglish(ctx context.Context, file multipart.File) (*transcribe.TranscriptionResponse, error)
	TranscribeIndonesian(ctx context.Context, file multipart.File) (*transcribe.TranscriptionResponse, error)
}

type TranscribeService struct {
	EnglishClient    transcribe.EnglishTranscriptionServiceClient
	IndonesianClient transcribe.IndonesianTranscriptionServiceClient
}

func NewTranscribeService(conn *grpc.ClientConn) ITranscribeService {
	return &TranscribeService{
		EnglishClient:    transcribe.NewEnglishTranscriptionServiceClient(conn),
		IndonesianClient: transcribe.NewIndonesianTranscriptionServiceClient(conn),
	}
}

func (s *TranscribeService) TranscribeEnglish(ctx context.Context, file multipart.File) (*transcribe.TranscriptionResponse, error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	req := &transcribe.TranscriptionRequest{
		FileData: fileBytes,
	}

	return s.EnglishClient.TranscribeAudio(ctx, req)
}

func (s *TranscribeService) TranscribeIndonesian(ctx context.Context, file multipart.File) (*transcribe.TranscriptionResponse, error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	req := &transcribe.TranscriptionRequest{
		FileData: fileBytes,
	}

	return s.IndonesianClient.TranscribeAudio(ctx, req)
}
