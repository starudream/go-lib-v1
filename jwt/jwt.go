package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/seq"

	"github.com/starudream/go-lib/internal/ilog"
)

const (
	_alg    = "HS256"
	_secret = "e76573a0-4867-41d2-9287-b1099799f0ad"
	_valid  = 7 * 24 * time.Hour
)

var (
	method jwt.SigningMethod

	rsa        bool
	secretKey  []byte
	publicKey  any
	privateKey any
	valid      time.Duration

	NowFunc = time.Now
)

func init() {
	alg := config.GetString("jwt.alg")
	if alg == "" {
		alg = _alg
	}
	method = jwt.GetSigningMethod(strings.ToUpper(alg))

	secret, public, private := config.GetString("jwt.secret"), config.GetString("jwt.public"), config.GetString("jwt.private")

	switch method {
	case jwt.SigningMethodHS256, jwt.SigningMethodHS512:
		if secret == "" {
			secret = _secret
		}
		secretKey = []byte(secret)
	case jwt.SigningMethodRS256, jwt.SigningMethodRS512:
		if public == "" || private == "" {
			ilog.X.Fatal().Msg("jwt public or private key is empty")
		}
		var err error
		publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(public))
		if err != nil {
			ilog.X.Fatal().Msgf("parse jwt public key failed: %v", err)
		}
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(private))
		if err != nil {
			ilog.X.Fatal().Msgf("parse jwt private key failed: %v", err)
		}
		rsa = true
	default:
		ilog.X.Fatal().Msgf("unsupported jwt alg: %s", alg)
	}

	valid = config.GetDuration("jwt.valid")
	if valid < time.Hour {
		valid = _valid
	}
}

type Claims struct {
	Id       string `json:"jti,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	Subject  string `json:"sub,omitempty"`
	IssuedAt int64  `json:"iat,omitempty"`

	Extra map[string]any `json:"ext,omitempty"`
}

var _ jwt.Claims = (*Claims)(nil)

func (c *Claims) Valid() error {
	if c == nil {
		return fmt.Errorf("claims is nil")
	}

	if c.IssuedAt == -1 {
		return nil
	}

	if NowFunc().Add(-valid).Unix() > c.IssuedAt {
		return fmt.Errorf("token is expired")
	}

	return nil
}

func (c *Claims) Sign() (string, error) {
	if c == nil {
		return "", fmt.Errorf("claims is nil")
	}

	if c.Id == "" {
		c.Id = seq.NextId()
	}
	if c.IssuedAt == 0 {
		c.IssuedAt = NowFunc().Unix()
	}

	token := jwt.NewWithClaims(method, c)

	if rsa {
		return token.SignedString(privateKey)
	}
	return token.SignedString(secretKey)
}

func New(id, iss, sub string, iat time.Time, ext map[string]any) *Claims {
	return &Claims{
		Id:       id,
		Issuer:   iss,
		Subject:  sub,
		IssuedAt: iat.Unix(),
		Extra:    ext,
	}
}

func Parse(raw string, skipValidate ...bool) (*Claims, error) {
	c := &Claims{}

	_, err := jwt.ParseWithClaims(raw, c, func(token *jwt.Token) (any, error) {
		if rsa {
			return publicKey, nil
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if len(skipValidate) > 0 && skipValidate[0] {
		return c, c.Valid()
	}

	return c, nil
}
