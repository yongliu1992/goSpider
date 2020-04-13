package base

// LoggerLevel 代表日志输出级别。
type LogLevel uint8

const (
	// LevelDebug 代表调试级别，是最低的日志等级。
	LevelDebug LogLevel = iota + 1
	// LevelInfo 代表信息级别，是最常用的日志等级。
	LevelInfo
	// LevelWarn 代表警告级别，是适合输出到错误输出的日志等级。
	LevelWarn
	// LevelError 代表普通错误级别，是建议输出到错误输出的日志等级。
	LevelError
	// LevelFatal 代表致命错误级别，是建议输出到错误输出的日志等级。
	// 此级别的日志一旦输出就意味着`os.Exit(1)`立即会被调用。
	LevelFatal
	// LevelPanic 代表恐慌级别，是最高的日志等级。
	// 此级别的日志一旦输出就意味着运行时恐慌立即会被引发。
	LevelPanic
)
