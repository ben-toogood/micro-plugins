package healtcheck

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
)

type enableHealthCheck struct {
	healthCheckPath string
}

func (eh *enableHealthCheck) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "health-check-path",
			Usage:  "The URL path for the health check (default is '/health')",
			EnvVar: "HEALTH_CHECK_PATH",
		},
	}
}

func (eh *enableHealthCheck) Commands() []cli.Command {
	return nil
}

func (eh *enableHealthCheck) Handler() plugin.Handler {
	return func(ha http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == eh.healthCheckPath {
				w.WriteHeader(200)
				w.Write([]byte("Micro Health Check"))
				return
			}

			ha.ServeHTTP(w, r)
		})
	}
}

func (eh *enableHealthCheck) Init(ctx *cli.Context) error {
	if p := ctx.String("health-check-path"); len(p) > 0 {
		eh.healthCheckPath = p
	}

	return nil
}

func (eh *enableHealthCheck) String() string {
	return "Health Check"
}

// NewPlugin creates the HealthCheck plugin
func NewPlugin() plugin.Plugin {
	return &enableHealthCheck{healthCheckPath: "/health"}
}
