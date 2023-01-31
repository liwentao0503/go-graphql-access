package common

import (
	"context"
	"sync"
)

type Handle func() error

// GoAndWait 封装了sync.WaitGroup 直接构造好func 拉函数即可
func GoAndWait(ctx context.Context, handles []Handle) (err error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(handles))
	for i := range handles {
		go func(ctx context.Context, i int) {
			defer wg.Done()
			if gerr := handles[i](); gerr != nil {
				err = gerr
			}
		}(ctx, i)
	}
	wg.Wait()
	return
}
