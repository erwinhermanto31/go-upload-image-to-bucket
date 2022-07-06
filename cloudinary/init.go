package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go"
)

var Cld *cloudinary.Cloudinary

func Init() {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Println("Failed to intialize Cloudinary", err)
	}
	log.Println("cloudinary Success connected!!")

	Cld = cld
}
