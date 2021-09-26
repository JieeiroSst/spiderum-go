package middleware

import (
	"errors"
	"fmt"
	cacheErr "github.com/allegro/bigcache"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"gitlab.com/Spide_IT/spide_it/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/Spide_IT/spide_it/pkg/bigcache"
)

type Authorization struct {
	cache bigcache.Cache
	jwt *jwt.TokenUser
}

func NewAuthorization(cache bigcache.Cache,jwt *jwt.TokenUser) *Authorization{
	return &Authorization{
		cache:cache,
		jwt:jwt,
	}
}

func (a *Authorization) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken:=c.Request.Header.Get("Authorization")
		if bearToken == ""{
			c.JSON(403,map[string]interface{}{"message":"Authentication failure: Token not provided","status":403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		strArr := strings.Split(bearToken, " ")
		message,err:=a.jwt.ParseToken(strArr[1])
		if err!=nil{
			c.JSON(403,map[string]interface{}{"message":message,"status":403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		sessionId, _ := c.Cookie("current_subject")
		sub,err:= a.cache.Get(sessionId)
		if errors.Is(err, cacheErr.ErrEntryNotFound) {
			c.JSON(401,map[string]interface{}{"message":"user hasn't logged in yet","status":401})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

func (a *Authorization) Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("current_subject")
		val, existed := a.cache.Get(cookie)
		if existed != nil {
			c.JSON(401,map[string]interface{}{"message":"user hasn't logged in yet","status":401})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ok, err := enforce(string(val), obj, act, adapter)
		if err != nil {
			c.JSON(403,map[string]interface{}{"message":"error occurred when authorizing user","status":403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !ok {
			c.JSON(403,map[string]interface{}{"message":"forbidden","status":403})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("pkg/conf/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}

