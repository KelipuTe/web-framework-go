package v22

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type HookFunc func(c context.Context) error

// NotifyShutdownToGateway 通知网关，服务要下线了
func NotifyShutdownToGateway(p1c context.Context) error {
	fmt.Println("notify gateway server will shutdown")
	time.Sleep(time.Second * 2)
	return nil
}

// ServiceShutdownBuilder 关闭服务
func ServiceShutdownBuilder(arr1p1s ...Service) HookFunc {
	return func(c context.Context) error {
		syncwg := sync.WaitGroup{}
		chansyncwg := make(chan struct{})

		syncwg.Add(len(arr1p1s))
		for _, t1p1s := range arr1p1s {
			go func(p1s Service) {
				err := p1s.Shutdown(c)
				if err != nil {
					fmt.Printf("service shutdown err: %+v\r\n", err)
				}
				time.Sleep(time.Second)
				syncwg.Done()
			}(t1p1s)
		}

		go func() {
			syncwg.Wait()
			chansyncwg <- struct{}{}
		}()

		select {
		case <-chansyncwg:
			fmt.Printf("ServiceShutdown, close all servers\r\n")
			return nil
		case <-c.Done():
			fmt.Printf("ServiceShutdown, Context Done\r\n")
			return errors.New("ServiceShutdown, Context Done")
		}
	}
}
