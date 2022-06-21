package constants

type RateLimitKey string

const (
	/**
	* 请求URI前缀
	 */
	URI_PREFIX = "ratelimit:uri:"

	/**
	* 上次访问时间
	 */
	LAST_REFILL_TIME = "lastRefillTime"
	/**
	 * 剩余可用访问次数
	 */
	REMAIN_TOKEN = "remainToken"
)
