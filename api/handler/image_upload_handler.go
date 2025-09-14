package handler

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/y-ttkt/baker/gen/api"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"sync"
)

type ImageUploadHandler struct {
	api.UnimplementedImageUploadServiceServer
	sync.Mutex
	files map[string][]byte
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		files: make(map[string][]byte),
	}
}

func (h *ImageUploadHandler) Upload(stream grpc.ClientStreamingServer[api.ImageUploadRequest, api.ImageUploadResponse]) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	meta := req.GetFileMeta()
	filename := meta.Filename

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	uuid := u.String()
	buf := &bytes.Buffer{}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.files[filename] = data
	err = stream.SendAndClose(&api.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    filename,
		ContentType: mimeType,
	})
	return err
}
