package constants

type PostgreSSLMode string

const (
	SSLModeVerifyCa   = "verify-ca"
	SSLModeVerfiyFull = "verify-full"
	SSLModeDisable    = "disable"
)
