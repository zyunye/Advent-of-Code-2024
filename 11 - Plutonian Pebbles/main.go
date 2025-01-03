package main

import (
	. "aoc"
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

func read_input(file_name string) []int {
	file, err := os.Open(file_name)
	CheckErr(err)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	str_nums := strings.Fields(line)

	ret := make([]int, len(str_nums))
	for i, v := range str_nums {
		ret[i], err = strconv.Atoi(v)
		CheckErr(err)
	}

	return ret
}

func get_num_len(num int) int {
	if num == 0 {
		return 1
	}

	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}

	return digits
}

func split_stone(num int, num_len int) (int, int) {
	div := int(math.Pow10(num_len / 2))
	left := num / div
	right := num - (left * div)

	return left, right
}

func process_stone(num int, len_cache *SafeLenCache, ret_ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	len_cache.mu.Lock()
	num_len, ok := len_cache.cache[num]
	if !ok {
		num_len = get_num_len(num)
		len_cache.cache[num] = num_len
	}
	len_cache.mu.Unlock()

	if num == 0 {
		ret_ch <- 1
	} else if num_len%2 == 0 {
		l, r := split_stone(num, num_len)
		ret_ch <- l
		ret_ch <- r
	} else {
		ret_ch <- num * 2024
	}
}

func worker(jobs_ch <-chan int, len_cache *SafeLenCache, ret_ch chan<- int, wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case stone := <-jobs_ch:
			process_stone(stone, len_cache, ret_ch, wg)
		}
	}
}

type SafeLenCache struct {
	mu    sync.Mutex
	cache map[int]int
}

func part_len_cache(file_name string) {
	pebbles := read_input(file_name)
	pebbles_buffer := make([]int, 0)
	pebbles_buffer = append(pebbles_buffer, pebbles...)

	len_cache := SafeLenCache{cache: make(map[int]int)}

	var worker_lock sync.WaitGroup
	// var main_lock sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	numWorkers := 16
	jobs_ch := make(chan int, numWorkers)
	ret_ch := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(jobs_ch, &len_cache, ret_ch, &worker_lock, ctx)
	}

	for blink := 0; blink < 75; blink++ {
		pebbles, pebbles_buffer = pebbles_buffer, pebbles
		fmt.Printf("Blink: %d, %d\n", blink, len(pebbles))

		go func() {
			for i := 0; i < len(pebbles); i++ {
				num := pebbles[i]
				worker_lock.Add(1)
				jobs_ch <- num
			}
			worker_lock.Wait()
			ret_ch <- -1
		}()


		i := 0
		for stone := range ret_ch {
			if stone == -1 {
				break
			}
			if i < len(pebbles_buffer) {
				pebbles_buffer[i] = stone
				i++
			} else {
				pebbles_buffer = append(pebbles_buffer, stone)
				i++
			}
		}
	}
	cancel()

	fmt.Printf("P.1_len: %d\n", len(pebbles_buffer))

}

func main() {
	file_name := "input.txt"
	var start time.Time

	f, err := os.Create("perf.prof")
	CheckErr(err)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// start := time.Now()
	// part1(file_name)
	// fmt.Printf("P1 time: %s", time.Since(start))

	start = time.Now()
	part_len_cache(file_name)
	fmt.Printf("P1 time: %s", time.Since(start))

}
