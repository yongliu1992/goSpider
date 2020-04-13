package base

type LogFormat string

const (
	// FormatText 代表普通文本的日志格式。
	FormatText LogFormat = "text"
	// FormatJson json。
	FormatJson LogFormat = "json"
)

// TimeStampFormat 代表时间戳格式化字符串。
const TimeStampFormat = "2006-01-02T15:04:05.999"
