package constants

import "github.com/go-playground/validator/v10"

type ContextKey string

const (
	UserKey         ContextKey = "user"
	SessionKey      ContextKey = "session"
	NavItemsKey     ContextKey = "navItems"
	AllNavItemsKey  ContextKey = "allNavItems"
	TxKey           ContextKey = "poolTx"
	PoolKey         ContextKey = "pool"
	ParamsKey       ContextKey = "params"
	LoggerKey       ContextKey = "logger"
	AppKey          ContextKey = "app"
	LogoKey         ContextKey = "logo"
	HeadKey         ContextKey = "head"
	SidebarPropsKey ContextKey = "sidebarProps"
	TabsKey         ContextKey = "tabs"
	RequestStart    ContextKey = "requestStart"

	PageContext ContextKey = "pageContext"
	TenantIDKey ContextKey = "tenant"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())
