package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"os"
	"time"
)

func GetFileDetailSchema() *schema.Schema {
	log.Println("Entering GetFileDetailSchema()")
	defer log.Println("Exiting GetFileDetailSchema()")
	return &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
		Elem: &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func GetFileDetail(context context.Context, fileNameValue interface{}) (map[string]interface{}, error) {
	log.Println("Entering GetFileDetail(Context, string)")
	defer log.Println("Exiting GetFileDetail(Context, string)")
	var (
		fileName = ""
		ok       bool
	)
	fileDetails := make(map[string]interface{})
	if fileName, ok = fileNameValue.(string); !ok {
		return fileDetails, errors.New(fmt.Sprintf("file name passed %v is not string", fileNameValue))
	}
	if fileName == "" {
		return fileDetails, errors.New("empty file name")
	}
	var fileObject os.FileInfo
	var err error
	if fileObject, err = os.Stat(fileName); err != nil {
		return fileDetails, err
	} else {
		fileDetails["path"] = fileName
		fileDetails["name"] = fileObject.Name()
		fileDetails["isDirectory"] = fmt.Sprintf("%t", fileObject.IsDir())
		fileDetails["size"] = fmt.Sprintf("%d", fileObject.Size())
		fileDetails["mode"] = fmt.Sprintf("%v", fileObject.Mode())
		fileDetails["last_modified"] = fmt.Sprintf("%v", fileObject.ModTime().Format(time.RFC3339))
	}
	return fileDetails, nil
}
