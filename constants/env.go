package constants

const (
	TestEnv       = "test"
	DevEnv        = "dev"
	StagingEnv    = "staging"
	ProductionEnv = "prod"
	DefaultEnv    = DevEnv
)

var (
	AllowedEnvs = map[string]string{
		TestEnv:       TestEnv,
		DevEnv:        DevEnv,
		StagingEnv:    StagingEnv,
		ProductionEnv: ProductionEnv,
	}
)
