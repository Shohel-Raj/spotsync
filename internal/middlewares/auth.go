package middlewares

import (
	"net/http"
	"spotssync/internal/auth"
	"strings"

	"github.com/labstack/echo/v5"
)

const (
	roleAdmin  = "admin"
	driverRole = "driver"
)

func AuthMiddleware(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// extract token from authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"success": "false",
					"massage": "Missing authorization header",
				})
			}

			// check bearer scheme
			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"success": "false",
					"error":   "invalid authorization header format",
				})
			}

			tokenString := parts[1]

			// validate token

			claims, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"success": "false",
					"error":   "invalid or expired token",
				})
			}
			// store user info in context for handlers
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_name", claims.Name)
			c.Set("user_role", claims.Role)

			return next(c)
		}
	}
}
func AdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			userRole := c.Get("user_role").(string)
			if userRole != roleAdmin {
				return c.JSON(http.StatusForbidden, map[string]string{
					"success": "false",
					"message": "You do not have permission to access this resource",
					"error":   "admin permission required",
				})
			}

			return next(c)
		}
	}
}
