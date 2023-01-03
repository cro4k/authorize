package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID                            string `json:"id"`
	CID                           string `json:"cid"`
	Username                      string `json:"username"`
	CertificateStatus             int    `json:"cs"`
	OrganizationID                string `json:"oid,omitempty"`
	OrganizationRole              int    `json:"or,omitempty"`
	OrganizationCertificateStatus int    `json:"ocs,omitempty"`
	jwt.StandardClaims
}
