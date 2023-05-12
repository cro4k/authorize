package oauth2

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"

	"github.com/cro4k/authorize/internal/dao"
	"github.com/cro4k/authorize/internal/db"
)

type Service struct {
	*server.Server
}

func NewService() *Service {
	s := new(Service)
	s.init()
	return s
}

func (s *Service) init() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapClientStorage(dao.Application(db.DB()))

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		//log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		//log.Println("Response Error:", re.Error.Error())
	})

	s.Server = srv

	//http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
	//	err := srv.HandleAuthorizeRequest(w, r)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//	}
	//})
	//
	//http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
	//	srv.HandleTokenRequest(w, r)
	//})
	//
	//log.Fatal(http.ListenAndServe(":9096", nil))
}
