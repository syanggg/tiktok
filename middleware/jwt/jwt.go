package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/util"
	"time"
)

type Claims struct {
	UserId int32  //保证命名开头大写，开头大写生成token时才生效
	jwt.StandardClaims
}

// JwtMiddleware 鉴权
func JwtMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		token := context.Query("token")
		if token == ""{
			token = context.PostForm("token")
		}

		if token == ""{
			context.JSON(http.StatusOK,models.CommonResponse{
				StatusCode: 401,
				StatusMsg: "token is not found",
			})
			context.Abort() //阻止执行
			return
		}

		//解析并验证token
		claims,err := ParseToken(token)
		if err != nil{
			context.JSON(http.StatusOK,models.CommonResponse{
				StatusCode: 403,
				StatusMsg: "token incorrect",
			})
			context.Abort() //阻止执行
			return
		}

		//判断token超时
		if claims.ExpiresAt < time.Now().Unix(){
			context.JSON(http.StatusOK,models.CommonResponse{
				StatusCode: 402,
				StatusMsg: "token timeout",
			})
			context.Abort()
			return
		}

		context.Set("user_id",claims.UserId)
		context.Next()
	}
}

func FeedMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		token := context.Query("token")
		//token存在时
		if token != ""{
			//记录当前token是否出现错误 出现错误按照用户未登录处理
			state := false
			//解析并验证token
			claims,err := ParseToken(token)
			if err != nil{
				context.JSON(http.StatusOK,models.CommonResponse{
					StatusCode: 403,
					StatusMsg: "token incorrect",
				})
				state = true	//出现错误 记录
			}

			//判断token超时
			if claims.ExpiresAt < time.Now().Unix(){
				context.JSON(http.StatusOK,models.CommonResponse{
					StatusCode: 402,
					StatusMsg: "token timeout",
				})
				state = true	//出现错误 记录
			}

			//未出现错误
			if !state {
				context.Set("user_id",claims.UserId)
			}
		}


		context.Next()
	}
}


//ReleaseToken 发放token
func ReleaseToken(user *models.UserLogin) (string,error){
	//指定到期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			//过期时间（时间戳）
			ExpiresAt: expirationTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			//发放者
			Issuer: "tiktok_syang",
			//主题
			Subject: "user token",
		},
	}

	//指定加密算法加密claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//利用密钥加密header和payload
	tokenString,err :=  token.SignedString([]byte(util.Info.Jwt.Key))
	if err != nil{
		return "",err
	}

	return tokenString,nil
}




// ParseToken 解析token
func ParseToken(tokenString string) (*Claims,error){
	claims := &Claims{}

	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.Info.Jwt.Key),nil
	})
	if err != nil{
		return nil,err
	}

	if !token.Valid {
		return nil,errors.New("invalid token")
	}

	return claims,nil

}
