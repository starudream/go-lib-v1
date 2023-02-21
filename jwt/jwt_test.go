package jwt

import (
	_ "embed"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/starudream/go-lib/testx"
)

var (
	// public openssl rsa -in private -pubout -out public
	//go:embed public
	public []byte
	// private openssl genrsa -out private 2048
	//go:embed private
	private []byte
)

func Test(t *testing.T) {
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(public)
	priKey, _ := jwt.ParseRSAPrivateKeyFromPEM(private)

	type c struct {
		method     jwt.SigningMethod
		rsa        bool
		secretKey  []byte
		publicKey  any
		privateKey any
		excepted   string
	}

	cs := []c{
		{
			method:    jwt.SigningMethodHS256,
			secretKey: []byte("1234567890"),
			excepted:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIweDAiLCJpc3MiOiJhZG1pbiIsInN1YiI6IjEwMDAwIiwiaWF0IjoxNjcyNTMxMjAwfQ.KQka5LFaEUfJh9SVS7ccplbV8Rz0xjzHkzyGYsQxpf8",
		},
		{
			method:    jwt.SigningMethodHS512,
			secretKey: []byte("1234567890"),
			excepted:  "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIweDAiLCJpc3MiOiJhZG1pbiIsInN1YiI6IjEwMDAwIiwiaWF0IjoxNjcyNTMxMjAwfQ.Myrzt46rPJlY9aGx0XZLTHuMGJacRkTHYRsXg2Y3pcAejgJBvw-RdzYXp2aQGLSSoCKC-wb91gPoGCHKneSvoA",
		},
		{
			method:     jwt.SigningMethodRS256,
			rsa:        true,
			publicKey:  pubKey,
			privateKey: priKey,
			excepted:   "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIweDAiLCJpc3MiOiJhZG1pbiIsInN1YiI6IjEwMDAwIiwiaWF0IjoxNjcyNTMxMjAwfQ.z8fVvKXQRFkB6uPoatgc0qag1qTTujARdzPTfbuF_qFxRTz8ZbblaIb5NqICce7X-8ri7t9-QQaDTxM5qa7_Zws_qVtZlCdNE5kmbNLoZvrlBCf46bnSGH0NRKYvw5mDNI6DVU30r5PgIn1x3IPKi6QZI9NhfIXoLUcg4-vx-qYWzRdxQtwK1itOiBsUO6bZG3OtrlLd7H-lCC_jAaOwW_uEcxEwIRF0uChvXBvO2Tg76UPXGF5vpzsn2YAzrG6otKvBGO9cbSYJw1QuSRjXFL-vquwgulLw9kAMv3Taj0bONbFlmJvrMbJmxDQ7M3oUbu8WDYAMbJDy3AlDanCu3A",
		},
		{
			method:     jwt.SigningMethodRS512,
			rsa:        true,
			publicKey:  pubKey,
			privateKey: priKey,
			excepted:   "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIweDAiLCJpc3MiOiJhZG1pbiIsInN1YiI6IjEwMDAwIiwiaWF0IjoxNjcyNTMxMjAwfQ.HHDeNJddoqBSVxcR8-yMiry9kPqfuTz4hHFQJfnBd3kusfHJDq6bfPnXql5FGIaa8_yH5yO48npsJk5TarnoF-qmL7NEv5PbXUL9JxXt567VH2K1Iz3gF6WNjtHp4UVv7KZPds3BwLWPyxQ0MebHjaaWHLQqlliK8nWIAxPsDAxiUe7IqutoPEQ1Q-XIb2GjczneMZ9G6stgEYT_BuXXQFiClDPMNuzib5pzIXOv82_BGAulXA1A4zYcEtOpz7tigZ3yWoivOIZ4od9FwRZYsykBC_Ufpk-xytcKVLKa8Csza7i0mWzh4SOOCCynyrUvVJogoXiMIS-GQuOMplnrrw",
		},
	}

	for i, v := range cs {
		t.Run(v.method.Alg(), func(t *testing.T) {
			method = v.method
			rsa = v.rsa
			secretKey = v.secretKey
			publicKey = v.publicKey
			privateKey = v.privateKey

			token, err := New("0x0", "admin", "10000", date, nil).Sign()
			testx.RequireNoErrorf(t, err, strconv.Itoa(i))
			testx.RequireEqualf(t, v.excepted, token, strconv.Itoa(i))

			claims, err := Parse(token)
			testx.P(t, err, claims)
		})
	}
}
