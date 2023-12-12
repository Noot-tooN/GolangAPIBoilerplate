package routes

type HealthcheckGroup struct {
	Server   string
	Postgre string
}

var (
	Healthcheck = HealthcheckGroup{
		Server:   "/server",
		Postgre: "/postgre",
	}
)