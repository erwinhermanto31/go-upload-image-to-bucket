package module

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	cld "github.com/erwinhermanto31/go-upload-image-to-bucket/cloudinary"
)

type UploadImage struct {
	cloudinary *cloudinary.Cloudinary
}

func NewUploadImage() *UploadImage {
	return &UploadImage{
		cloudinary: cld.Cld,
	}
}

func (h *UploadImage) Handler(ctx context.Context, image_file string) error {
	result, err := h.cloudinary.Upload.Upload(ctx, image_file, uploader.UploadParams{PublicID: "image", Folder: "siup"})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("upload success")
	log.Println(result.SecureURL)
	return nil
}
