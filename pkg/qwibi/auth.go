package qwibi

import (
	"github.com/qwibi/qwibi-go-sdk/pkg/auth"
	"github.com/qwibi/qwibi-go-sdk/pkg/auth/session"
	"github.com/qwibi/qwibi-go-sdk/pkg/metadata"
	"github.com/qwibi/qwibi-go-sdk/pkg/qlog"
	"github.com/qwibi/qwibi-go-sdk/proto"
)

// Auth ...
func (c *QApiClient) Auth(a auth.QAuth) (*session.QSession, error) {

	switch v := a.(type) {
	case *auth.QAnonymousAuth:
		return c.AnonymousAuth(v.Token)
	case *auth.QBasicAuth:
		return c.BasicAuth(v.AccountId, v.Password)
	default:
		return nil, qlog.Error("Unknown auth type")
	}
	// md := metadata.Pairs("token", res.Session.Token)
	// c.ctx = metadata.NewOutgoingContext(context.Background(), md)
	// return res.Session, nil
}

func (c *QApiClient) AnonymousAuth(token string) (*session.QSession, error) {
	req := &proto.QPBxAuthRequest{
		Auth: &proto.QPBxAuthRequest_Anonym{
			Anonym: &proto.QPBxAnonymAuth{
				Token: token,
			},
		},
	}

	res, err := c.apiClient.Auth(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	session, err := session.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}

func (c *QApiClient) BasicAuth(login string, password string) (*session.QSession, error) {
	req := &proto.QPBxAuthRequest{
		Auth: &proto.QPBxAuthRequest_Basic{
			Basic: &proto.QPBxBasicAuth{
				AccountId: login,
				Password:  password,
			},
		},
	}

	res, err := c.apiClient.Auth(c.ctx, req)
	if err != nil {
		return nil, qlog.Error(err)
	}

	session, err := session.NewSessionPb(res.Session)
	if err != nil {
		return nil, qlog.Error(err)
	}

	c.ctx = metadata.SetContextToken(c.ctx, session.Token)

	return session, nil
}
