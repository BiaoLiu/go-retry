package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestMaxRetryCount(t *testing.T) {
	r := NewRetry(0, 0, 10, 0)

	for i := 0; i < 20; i++ {
		err := r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}

func TestMaxRetryTime(t *testing.T) {
	r := NewRetry(0, 0, 0, 3*time.Second)

	for i := 0; i < 20; i++ {
		err := r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}

func TestMaxRetryCountAndTime(t *testing.T) {
	d := 1 * time.Second
	r := NewRetry(0, 0, 5000, d)

	for {
		err := r.Do(func(firstRetryTime int64, retriedCount int64) error {
			fmt.Println("retriedCount...", retriedCount)
			return nil
		})
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("total retriedCount", r.RetriedCount())
}
