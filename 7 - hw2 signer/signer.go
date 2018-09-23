package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func worker(wg *sync.WaitGroup, in, out chan interface{}, job job) {
	defer wg.Done()
	defer close(out)
	job(in, out)
}

func SingleHash(in, out chan interface{}) {

	singleWg, cache, mu := &sync.WaitGroup{}, make(map[string]map[string]string), &sync.Mutex{}
	for income := range in {

		singleWg.Add(1)
		go func(data interface{}) {
			defer singleWg.Done()

			d, ok := data.(int)
			if !ok {
				panic("Singe hash error, not int input" + data.(string))
			}

			i := strconv.Itoa(d)

			crc32, cachedSingle := cache[i]
			if !cachedSingle {

				crc32 = map[string]string{
					"string": i,
					"md5string": func() string {
						mu.Lock()
						r := DataSignerMd5(i)
						mu.Unlock()
						return r
					}(),
				}

				wg := &sync.WaitGroup{}
				muCrc32 := &sync.Mutex{}
				for key := range crc32 {
					wg.Add(1)
					go func(key string) {
						defer wg.Done()
						r := DataSignerCrc32(crc32[key])
						muCrc32.Lock()
						crc32[key] = r
						muCrc32.Unlock()
					}(key)
				}

				wg.Wait()
				cache[i] = crc32
			}
			out <- crc32["string"] + "~" + crc32["md5string"]
		}(income)
	}

	singleWg.Wait()
}

func MultiHash(in, out chan interface{}) {
	multiWg, cache := &sync.WaitGroup{}, make(map[string]string)

	for income := range in {

		multiWg.Add(1)
		go func(data interface{}) {
			defer multiWg.Done()
			str, ok := data.(string)
			if !ok {
				panic("Multi hash error")
			}

			_, foundCached := cache[str]
			if !foundCached {
				wg := &sync.WaitGroup{}
				th := make([]string, 6)
				for i := 0; i < 6; i++ {
					th[i] = str
					wg.Add(1)
					go func(thNum int) {
						defer wg.Done()
						th[thNum] = DataSignerCrc32(strconv.Itoa(thNum) + th[thNum])
					}(i)
				}
				wg.Wait()
				cache[str] = strings.Join(th, "")
			}
			out <- cache[str]
		}(income)
	}

	multiWg.Wait()
}

func CombineResults(in, out chan interface{}) {

	results := make([]string, 0)
	for income := range in {
		data, ok := income.(string)
		if !ok {
			panic("Result error")
		}
		results = append(results, data)
	}

	sort.Strings(results)
	out <- strings.Join(results, "_")
}

func ExecutePipeline(jobs ...job) {
	in, out, wg := make(chan interface{}, MaxInputDataLen), make(chan interface{}, MaxInputDataLen), &sync.WaitGroup{}
	for _, job := range jobs {
		wg.Add(1)
		go worker(wg, in, out, job)
		in = out
		out = make(chan interface{}, MaxInputDataLen)
	}

	wg.Wait()
}
