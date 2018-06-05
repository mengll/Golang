package cdn

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"smallgamepk.qcwanwan.com/config"
	"github.com/labstack/echo"
	"os"
	"smallgamepk.qcwanwan.com/utils"

	"strings"
	"net/http"
)

const DB_NAME = "cdn_alibaba"



func GetBucket(bucketName string) (*oss.Bucket, error) {
	// New Client
	client, err := oss.New(config.GetConfig(DB_NAME,"endpoint").(string), config.GetConfig(DB_NAME,"access_id").(string), config.GetConfig(DB_NAME,"access_key").(string))
	if err != nil {
		return nil, err
	}

	// Get Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func HandleError(err error) {
	utils.Log(err)
	os.Exit(-1)
}

//上传文件到阿里cdn
func UploadFileToCdn( c echo.Context) error{
	utils.Log(config.GetConfig(DB_NAME,"bucket_name").(string))
	bucket, err := GetBucket(config.GetConfig(DB_NAME,"bucket_name").(string))
	if err != nil {
		HandleError(err)
	}

	file, err := c.FormFile("game_img")
	if err != nil {
		return err
	}
	file_handele, err := file.Open()
	if err != nil {
		return err
	}

	defer file_handele.Close()
	paths := []string{"pkgame","assets",file.Filename}
	img_url := strings.Join(paths,"/")
	err = bucket.PutObject(img_url, file_handele)
	
	if err != nil {
		HandleError(err)
	}

	base_url := config.GetConfig(DB_NAME,"base_url").(string) + img_url
	bp := make(map[string]interface{})
	bp["img_url"] = base_url
	return c.JSON(http.StatusOK,bp)
}


