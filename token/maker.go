package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

var _ Maker = (*JWTMaker)(nil)
var _ Maker = (*PasetoMaker)(nil)
