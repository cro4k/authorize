package service

import (
	"github.com/cro4k/authorize/internal/service/auth"
	//"github.com/cro4k/authorize/internal/service/oauth2"
)

var (
	Auth = auth.NewService()
	//OAuth2 = oauth2.NewService()
	//FS     *minio.Client
)

//func init() {
//	var err error
//	FS, err = minio.New(config.C().Minio.Endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(config.C().Minio.AccessID, config.C().Minio.AccessKey, ""),
//		Secure: false,
//	})
//	if err != nil {
//		logrus.Fatal(err)
//	}
//}
