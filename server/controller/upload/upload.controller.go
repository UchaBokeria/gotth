package upload

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"main/server/common/controller"
	"main/server/common/globals"
	"main/server/common/storage"
	"main/server/model"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadResponse struct { ID int; Message string; Success bool } 

func FileUpload(ctx *controller.Context) error {
	// var About model.Interface_about
	// result := storage.DB.Last(&About)

	// if result.Error != nil {
	// 	return ctx.Html(view.ErrorPage())
	// }

	// return ctx.Html(view.Terms(About.Terms))
	// Retrieve the file from form data
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Error retrieving file from form data", Success: false },
		)
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Error opening received file", Success: false },
		)
	}
	defer src.Close()
	
	// Calculate SHA-256 hash of the file contents
	hash := sha256.New()
	if _, err := io.Copy(hash, src); err != nil {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Error calculating hash", Success: false },
		)
	}

	// Reset src to the beginning to read again
	src.Seek(0, 0)
	extension := GetFileExtension(file)
	hashName := hex.EncodeToString(hash.Sum(nil))

	// Create a new file on the server to store the uploaded file
	dst, err := os.Create(globals.Env.Uploads + hashName + extension)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Error creating file on server", Success: false },
		)
	}
	defer dst.Close()

	// Copy the file from the form data to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Error copying file to destination", Success: false },
		)
	}

	if len(extension) < 2 {
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "File type " + extension + " has a problem", Success: false },
		)
	}

	var Type model.File_types
	result := storage.DB.Where(&model.File_types{Ext: extension[1:]}).Last(&Type)

	if result.Error != nil {
		log.Print(result)
		return ctx.JSON(
			http.StatusBadRequest, 
			&UploadResponse{ ID: -1, Message: "Server can't accept " + extension + " type files", Success: false },
		)
	}

	var File model.Files = model.Files{
		Name: hashName + extension,
		Original: file.Filename,
		Size: int(file.Size),
		Location: globals.Env.Uploads,
		Path: globals.Env.Uploads + hashName + extension,
		Compressed: false,
		Base64: "",
		TypeID: int(Type.ID),
	}

	Result := storage.DB.Create(&File)
	if Result.Error != nil || Result.RowsAffected < 1 {
		log.Print(Result)
		return ctx.JSON(
			http.StatusNotAcceptable, 
			&UploadResponse{ ID: -1, Message: "File uploaded but was not saved in database", Success: false },
		)
	}

	return ctx.JSON(
		http.StatusOK, 
		&UploadResponse{ ID: int(File.ID), Message: "Successfully uploaded", Success: true },
	)
}

func GetFileExtension(file *multipart.FileHeader) string {
	header, err := file.Open()
	if err != nil {
		return ""
	}
	defer header.Close()

	extension := ""
	for i := len(file.Filename) - 1; i >= 0; i-- {
		if file.Filename[i] == '.' {
			extension = file.Filename[i:]
			break
		}
	}
	return extension
}