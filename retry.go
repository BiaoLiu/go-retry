package retry

import (
	"errors"
	"time"
)

var (
	ErrMaxRetryCount = errors.New("retry count has reached maximum limit")
	ErrMaxRetryTime  = errors.New("retry time has reached the maximum limit")
)

type RetryableFunc func(firstRetryTime int64, retriedCount int64) error

// Retry .
type Retry struct {
	firstRetryTime int64         // 首次重试时间
	retriedCount   int64         // 已重试次数
	maxRetryCount  int64         // 最大重试次数
	maxRetryTime   time.Duration // 最大重试时间
}

// NewRetry new a Retry.
func NewRetry(firstRetryTime, retriedCount, maxRetryCount int64, maxRetryTime time.Duration) *Retry {
	return &Retry{
		firstRetryTime: firstRetryTime,
		retriedCount:   retriedCount,
		maxRetryCount:  maxRetryCount,
		maxRetryTime:   maxRetryTime,
	}
}

func (r *Retry) FirstRetryTime() int64 {
	return r.firstRetryTime
}

func (r *Retry) RetriedCount() int64 {
	return r.retriedCount
}

// Do retry
func (r *Retry) Do(retryableFunc RetryableFunc) error {
	if r.firstRetryTime <= 0 {
		r.firstRetryTime = time.Now().Unix()
		r.retriedCount = 0
	}
	if r.retriedCount < 0 {
		r.retriedCount = 0
	}
	// firstRetryTimeStr := time.Unix(r.firstRetryTime, 0).Format("2006-01-02 15:04:05")
	if r.maxRetryTime > 0 {
		if time.Now().Add(-r.maxRetryTime).Unix() > r.firstRetryTime {
			return ErrMaxRetryTime
			// err := fmt.Errorf("重试时间已达到最大限制，终止处理... 首次重试时间:%v 重试次数:%v 最大重试时间:%v", firstRetryTimeStr, r.retryCount, r.maxRetryTime)
			// return err
		}
	}
	if r.maxRetryCount > 0 {
		if r.retriedCount >= r.maxRetryCount {
			return ErrMaxRetryCount
			// err := fmt.Errorf("重试次数已达到最大限制，终止处理... 首次重试时间:%v 重试次数:%v 最大重试次数:%v", firstRetryTimeStr, r.retryCount, r.maxRetryCount)
			// return err
		}
	}
	r.retriedCount += 1
	err := retryableFunc(r.firstRetryTime, r.retriedCount)
	return err
}
