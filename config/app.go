package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xframe-go/template/app/providers"
	"github.com/xframe-go/template/public"
	"github.com/xframe-go/template/routes"
	"github.com/xframe-go/x"
	"github.com/xframe-go/x/contracts"
	"github.com/xframe-go/x/frame/app"
	"github.com/xframe-go/x/frame/xhttp/middlewares"
	"github.com/xframe-go/x/net/xhttp"
	"github.com/xframe-go/x/utils/env"
	"log/slog"
)

func init() {
	x.App().Config().Register(&app.Config{
		// Application Name
		// This value is the name of your application, which will be used when the
		// framework needs to place the application's name in a notification or
		// other UI elements where an application name needs to be displayed.
		Name: env.String("APP_NAME", ""),

		Env: env.String("APP_ENV", "development"),

		Servers: map[string]contracts.ServerConfig{
			"default": {
				// HTTP SERVER PORT
				Port: env.Int("APP_PORT", 0),

				// HTTP SERVER HOST
				Host: env.String("APP_HOST", "127.0.0.1"),

				PublicFS: &public.Public,

				RoutingProviders: []contracts.RoutingProvider{
					routes.ApiRoutingProvider,
				},

				Middlewares: []contracts.Middleware{
					middlewares.RequestHandled,

					xhttp.EchoMiddleware(middleware.GzipWithConfig(middleware.GzipConfig{
						Level: 5,
					})),

					//xhttp.EchoMiddleware(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
					//	Store: middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{
					//		Rate: rate.Limit(10),
					//	}),
					//})),

					xhttp.EchoMiddleware(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
						LogURI:       true,
						LogStatus:    true,
						LogMethod:    true,
						LogLatency:   true,
						LogUserAgent: true,
						LogRemoteIP:  true,
						LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
							x.App().Log().Info(v.URI,
								slog.String("method", v.Method),
								slog.String("duration", v.Latency.String()),
								slog.String("client", v.UserAgent),
								slog.String("ip", v.RemoteIP),
							)
							return nil
						},
					})),
				},
			},
		},

		Providers: []contracts.Provider{
			&providers.AppServiceProvider{},

			//&pulse.Provider{},
		},
	})
}
