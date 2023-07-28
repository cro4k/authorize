package service

import (
	"github.com/cro4k/authorize/internal/db"
	"github.com/cro4k/authorize/internal/service/auth"
	"github.com/cro4k/authorize/internal/service/oauth2"
	"github.com/cro4k/authorize/internal/service/perm"
	//"github.com/cro4k/authorize/internal/service/oauth2"
)

var (
	Auth   = auth.NewService()
	OAuth2 = oauth2.NewService()
	Casbin *perm.Service

	//FS     *minio.Client
)

func Init() {
	var err error
	Casbin, err = perm.NewService(db.DB())
	if err != nil {
		panic(err)
	}
}

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
