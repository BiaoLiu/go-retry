## go-retry
go-retry，支持配置最大重试次数与最大重试时间。如果同时配置两个条件，只要当其中某一条件最先达成，将退出重试。

### 使用

配置最大重试次数：
```
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
```

配置最大重试时间：
```
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
```

同时配置最大重试次数与最大重试时间：
```
r := NewRetry(0, 0, 5000, 1 * time.Second)

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
```