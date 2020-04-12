package errors

import (
	"bytes"
	"fmt"
	"strings"
)

type ErrorType string

// 错误类型常量。
const (
	// ErrorTypeDownloader 代表下载器错误。
	ErrorTypeDownloader ErrorType = "downloader error"
	// ErrorTypeAnalyzer 代表分析器错误。
	ErrorTypeAnalyzer ErrorType = "analyzer error"
	// ErrorTypePipeline 代表条目处理管道错误。
	ErrorTypePipeline ErrorType = "pipeline error"
	// ErrorTypeScheduler 代表调度器错误。
	ErrorTypeScheduler ErrorType = "scheduler error"
)

// CrawlerError 代表爬虫错误的接口类型。
type SpiderError interface {
	// Type 用于获得错误的类型。
	Type() ErrorType
	// Error 用于获得错误提示信息。
	Error() string
}

// myCrawlerError 代表爬虫错误的实现类型。
type mySpiderError struct {
	// errType 代表错误的类型。
	errType ErrorType
	// errMsg 代表错误的提示信息。
	errMsg string
	// fullErrMsg 代表完整的错误提示信息。
	fullErrMsg string
}

func (c *mySpiderError) Type() ErrorType {
	return c.errType
}

func (c *mySpiderError) Error() string {
	if c.fullErrMsg == "" {
		c.genFullErrMsg()
	}
	return c.fullErrMsg
}

// genFullErrMsg 用于生成错误提示信息，并给相应的字段赋值。
func (c *mySpiderError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if c.errType != "" {
		buffer.WriteString(string(c.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(c.errMsg)
	c.fullErrMsg = fmt.Sprintf("%s", buffer.String())
}

func NewCrawlerError(errType ErrorType, errMsg string) SpiderError {
	return &mySpiderError{errType: errType, errMsg: strings.TrimSpace(errMsg)}
}

// NewCrawlerErrorBy 用于根据给定的错误值创建一个新的爬虫错误值。
func NewCrawlerErrorBy(errType ErrorType, err error) SpiderError {
	return NewCrawlerError(errType, err.Error())
}

//非法参数
type IllegalParameterError struct {
	msg string
}

// NewIllegalParameterError 会创建一个IllegalParameterError类型的实例。
func NewIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("illegal parameter: %s",
			strings.TrimSpace(errMsg)),
	}
}

func (ipe IllegalParameterError) Error() string {
	return ipe.msg
}
