package server

import (
	"fmt"
	"net/http"

	"github.com/DoWithLogic/go-rbac/internal/middlewares"
	"github.com/DoWithLogic/go-rbac/internal/users"
	userApi "github.com/DoWithLogic/go-rbac/internal/users/delivery/api"
	userRepo "github.com/DoWithLogic/go-rbac/internal/users/repository"
	userUsecase "github.com/DoWithLogic/go-rbac/internal/users/usecase"
	"github.com/DoWithLogic/go-rbac/pkg/redis"
	"github.com/DoWithLogic/go-rbac/pkg/security"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func (s *Server) mapHandler() error {
	// Define the domain and create a versioned group for API endpoints.
	var domain = s.echo.Group("api/v1/rbac")

	{
		domain.GET("/ping", func(c echo.Context) error {
			return c.JSON(http.StatusOK, fmt.Sprintf("welcome to Role and Scope Based Access Control (RBAC) version %s 👋", s.cfg.App.Version))
		})

		// Serve Swagger documentation
		domain.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	var (
		redis      = redis.NewRedis(s.redisConn)
		jwtFactory = security.NewJWTFactory(s.cfg.Security.JWT, redis)
		mw         = middlewares.NewMiddleware(jwtFactory)

		repositories = newRepositories(s.db)
		usecases     = newUsecases(repositories)
	)
	userHandler := userApi.NewUsersHandler(usecases.userUC)
	userApi.MapUserRoutes(domain, userHandler, mw)

	return nil
}

type repositories struct {
	authRepo users.Repositories
}

func newRepositories(db *gorm.DB) repositories {
	return repositories{
		authRepo: userRepo.NewUsersRepository(db),
	}
}

type usecases struct {
	userUC users.Usecases
}

func newUsecases(r repositories) usecases {
	return usecases{
		userUC: userUsecase.NewUsersUC(userUsecase.Dependencies{
			Repositories: userUsecase.Repositories{
				UserRepo: r.authRepo,
			},
		}),
	}
}
