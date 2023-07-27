package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestMaxRetryCount(t *testing.T) {
	var err error
	r := NewRetry(0, 0, 10, 0)

	for {
		err = r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err == nil {
			break
		}
		if errors.Is(err, ErrMaxRetryCount) {
			break
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}

func TestMaxRetryTime(t *testing.T) {
	var err error
	r := NewRetry(0, 0, 0, 3*time.Second)

	for {
		err = r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err == nil {
			break
		}
		if errors.Is(err, ErrMaxRetryTime) {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}

func TestMaxRetryCountAndTime(t *testing.T) {
	var err error
	r := NewRetry(0, 0, 5000, 1*time.Second)

	for {
		err = r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err == nil {
			break
		}
		if errors.Is(err, ErrMaxRetryCount) || errors.Is(err, ErrMaxRetryTime) {
			break
		}
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}
