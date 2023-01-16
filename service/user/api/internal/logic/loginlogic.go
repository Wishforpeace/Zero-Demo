package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"go-zero-demo/common/errorx"
	"go-zero-demo/service/user/model"
	"strings"
	"time"

	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//fmt.Println("---------------------------------")
	if len(strings.TrimSpace(req.UserNumber)) == 0 || len(strings.TrimSpace(req.Passoword)) == 0 {
		return nil, errors.New("参数错误")
	}
	//fmt.Println(req)
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.UserNumber)
	//fmt.Println(userInfo)
	switch err {
	case nil:

	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Passoword {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}
	//------start------
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	fmt.Println(accessExpire)
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	//fmt.Println(jwtToken)
	//------end------
	return &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	//fmt.Printf("%v\n %v\n %v\n %v\n", secretKey, iat, seconds, userId)
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err := token.SignedString([]byte(secretKey))
	return jwtToken, err
}
