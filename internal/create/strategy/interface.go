package strategy

const (
	//handler
	FIBER    = "fiber"
	NET_HTTP = "net/http"
	CHI      = "chi"
	//logger
	ZEROLOG = "zerolog"
	ZAP     = "zap"
)

type CreateStrategy interface {
	Execute() error
}
