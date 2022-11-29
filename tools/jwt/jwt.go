package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义JWT的过期时间
const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于(签名)加盐的字符串
var CustomSecret = []byte("夏天夏天悄悄过去")


// CustomClaims 自定义声明类型 并内嵌（继承）jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录username和userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID		int64 `json:"userid"`
	Username    string `json:"username"`
	jwt.RegisteredClaims // 内嵌标准声明，相当于继承
}


// GenToken 生成JWT
func GenToken(username string, userid int64) (string, error){
	// 创建一个自定义声明
	claims := CustomClaims{
		userid,
		username,
		jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
					Issuer: "bluebell", // 签发人
		},
	}
	// 使用指定的方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整编码后的字符串token
	return token.SignedString(CustomSecret)
}


// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}