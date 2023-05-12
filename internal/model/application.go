package model

type Application struct {
	UIDModel
	Secret string
	Domain string
}

func (a *Application) GetID() string {
	return a.ID
}

func (a *Application) GetSecret() string {
	return a.Secret
}

func (a *Application) GetDomain() string {
	return a.Domain
}

func (a *Application) IsPublic() bool {
	return false
}

func (a *Application) GetUserID() string {
	//DO NOTHING
	return ""
}

type ApplicationGrant struct {
	UIDModel
}
