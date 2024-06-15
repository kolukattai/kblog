package util

import "sync"

func IterationConcurrently[T any](arr []T, yield func(item T, index int)) {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for i, v := range arr {
		go func(v T, i int) {
			defer wg.Done()
			yield(v, i)
		}(v, i)
	}
	wg.Wait()
}





