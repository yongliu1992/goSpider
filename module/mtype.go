package module

// Type 代表组件的类型。
type Type string

const (
	// TypeDownloader 代表下载器。
	TypeDownloader Type = "downloader"
	// TypeAnalyzer 代表分析器。
	TypeAnalyzer Type = "analyzer"
	// TypePipeline 代表条目处理管道。
	TypePipeline Type = "pipeline"
)

// legalLetterTypeMap 代表合法的组件类型-字母的映射。
var legalLetterTypeMap = map[Type]string{
	TypeDownloader: "D",
	TypeAnalyzer:   "A",
	TypePipeline:   "P",
}

// CheckType 用于判断组件实例的类型是否匹配。
func CheckType(moduleType Type, model Module) bool {
	if moduleType == "" || module == nil {
		return false
	}
	switch moduleType {
	case TypeDownloader:
		if _, ok := model.(Downloader); ok {
			return true
		}
	case TypeAnalyzer:
		if _, ok := model.(Analyzer); ok {
			return true
		}
	case TypePipeline:
		if _, ok := model.(Pipeline); ok {
			return true
		}
	}
	return false
}
