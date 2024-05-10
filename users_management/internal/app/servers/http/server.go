package auth_server

import (
	"log"
	"net"
	"net/http"
	"time"

	"users_management/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"

	mw "users_management/internal/app/servers/middleware"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type AuthHttpServerConfig struct {
	Address string
	IsDev   bool
}

type authHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHttpServer struct {
	authHandler authHandler
}

// @contact.name   API Support
// @license.name  MIT
func NewAuthHttpServer(authHandler authHandler) *authHttpServer {
	return &authHttpServer{
		authHandler: authHandler,
	}
}

func (s *authHttpServer) Run(config AuthHttpServerConfig) error {
	r := chi.NewRouter()
	r.Use(mw.MetricsMiddleware)
	r.Use(mw.LoggerMiddleware)
	r.Use(middleware.Timeout(time.Second * 20))
	r.Post("/login", s.authHandler.Login)
	r.Post("/refresh", s.authHandler.Refresh)
	r.Post("/register", s.authHandler.Register)

	if config.IsDev {
		docs.SwaggerInfo.Title = "Swagger authorization service API"
		docs.SwaggerInfo.Description = "Authorization service"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.BasePath = ""
		docs.SwaggerInfo.Schemes = []string{"http"}

		r.Get(
			"/swagger/*",
			httpSwagger.Handler(httpSwagger.URL("./swagger/doc.json")),
		)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:      []string{"*"},
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
	})
	corsMux := c.Handler(r)

	conn, err := net.Listen("tcp", config.Address)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Printf("[http-server] Http server starting... [%s]", config.Address)
	return http.Serve(conn, corsMux)
}
