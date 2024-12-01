package middlewares

import (
	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/DoWithLogic/go-rbac/pkg/response"
	"github.com/DoWithLogic/go-rbac/pkg/security"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	jwt *security.JWTFactory
}

func NewMiddleware(jwt *security.JWTFactory) *Middleware {
	return &Middleware{jwt: jwt}
}

func (m *Middleware) JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get(constants.AuthorizationHeaderKey.String())
			if tokenString == "" {
				return response.ErrorBuilder(app_errors.Unauthorized(app_errors.ErrMissingJWTToken)).Send(c)
			}

			// Remove "Bearer " prefix from token string
			tokenString = tokenString[len("Bearer "):]
			if m.jwt.IsTokenBlacklisted(c.Request().Context(), tokenString) {
				return response.ErrorBuilder(app_errors.Unauthorized(app_errors.ErrInvalidJWTToken)).Send(c)
			}

			jwtClaims, err := m.jwt.VerifyJWT(tokenString)
			if err != nil {
				return response.ErrorBuilder(err).Send(c)
			}
			c.Set(constants.AuthCredentialContextKey.String(), jwtClaims)

			return next(c)
		}
	}
}

func (m *Middleware) PermissionsMiddleware(permissions ...constants.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtData, ok := c.Get(constants.AuthCredentialContextKey.String()).(*security.JWTClaims)
			if !ok {
				return response.ErrorBuilder(app_errors.Forbidden(app_errors.ErrAccessDenied)).Send(c)
			}

			// Check if at least one of the required permissions exists in the user's permissions
			if !m.hasRequiredPermission(jwtData.Permissions, permissions) {
				return response.ErrorBuilder(app_errors.Forbidden(app_errors.ErrAccessDenied)).Send(c)
			}

			// Store the token claims in the request context for later use
			c.Set(constants.AuthCredentialContextKey.String(), jwtData)

			// Continue to the next handler
			return next(c)
		}
	}
}

// hasRequiredPermission checks if any of the required permissions exist in the user's permissions
func (m *Middleware) hasRequiredPermission(userPermissions, requiredPermissions []constants.Permission) bool {
	// Create a map for fast lookups of required permissions
	requiredPermissionsMap := make(map[constants.Permission]bool)
	for _, permission := range requiredPermissions {
		requiredPermissionsMap[permission] = true
	}

	// Check if any of the user's permissions match the required permissions
	for _, permission := range userPermissions {
		if requiredPermissionsMap[permission] {
			return true
		}
	}

	return false
}

func (m *Middleware) RolesMiddleware(roles ...constants.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtData, ok := c.Get(constants.AuthCredentialContextKey.String()).(*security.JWTClaims)
			if !ok {
				return response.ErrorBuilder(app_errors.Forbidden(app_errors.ErrAccessDenied)).Send(c)
			}

			if !m.hasRequiredRole(jwtData.Role, roles) {
				return response.ErrorBuilder(app_errors.Forbidden(app_errors.ErrAccessDenied)).Send(c)
			}

			// Store the token claims in the request context for later use
			c.Set(constants.AuthCredentialContextKey.String(), jwtData)

			return next(c)
		}
	}
}

func (m *Middleware) hasRequiredRole(userRole constants.UserRole, roles []constants.UserRole) bool {
	for _, r := range roles {
		if r == userRole {
			return true
		}
	}
	return false
}
