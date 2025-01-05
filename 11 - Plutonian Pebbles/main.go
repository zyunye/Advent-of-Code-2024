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

type StoneAttr struct {
	len      int
	children []int
}

type SafeMap[T comparable, P any] struct {
	mu    sync.Mutex
	cache map[T]P
}

type UpdatePacket struct {
	stone int
	attrs StoneAttr
}

type StoneJob struct {
	stone int
	count int
}

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

func bulk_process(stone int, count int, stone_attrs *StoneAttr, result chan map[int]int, updates chan<- UpdatePacket) {
	if stone == 0 {
		result <- map[int]int{1: count}
	} else {
		if stone_attrs == nil {
			num_len := get_num_len(stone)
			new_stone_attrs := StoneAttr{children: make([]int, 1)}

			if num_len%2 == 0 {
				l, r := split_stone(stone, num_len)
				new_stone_attrs.children[0] = l
				new_stone_attrs.children = append(new_stone_attrs.children, r)
			} else {
				new_stone_attrs.children[0] = stone * 2024
			}

			updates <- UpdatePacket{stone: stone, attrs: new_stone_attrs}
			stone_attrs = &new_stone_attrs
		}

		children := stone_attrs.children
		if len(children) == 1 {
			result <- map[int]int{children[0]: count}
		} else {
			result <- map[int]int{children[0]: count, children[1]: count}
		}
	}
}

func cache_updater(stone_cache *SafeMap[int, StoneAttr], updates <-chan UpdatePacket, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case pkt := <-updates:
			stone_cache.mu.Lock()
			stone_cache.cache[pkt.stone] = pkt.attrs
			stone_cache.mu.Unlock()
		}
	}
}

func pebbles_updater(results_buffer *SafeMap[int, int], results_ch <-chan map[int]int, wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case stone_update := <-results_ch:
			results_buffer.mu.Lock()
			for k, v := range stone_update {
				results_buffer.cache[k] += v
			}
			results_buffer.mu.Unlock()
			wg.Done()
		}
	}
}

func worker(
	jobs_ch <-chan StoneJob,
	results_ch chan map[int]int,
	updates_ch chan UpdatePacket,
	stone_cache *SafeMap[int, StoneAttr],
	wg *sync.WaitGroup,
	ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case stone_job := <-jobs_ch:
			stone, count := stone_job.stone, stone_job.count

			stone_cache.mu.Lock()
			attr, ok := stone_cache.cache[stone]
			stone_cache.mu.Unlock()

			if ok {
				bulk_process(stone, count, &attr, results_ch, updates_ch)
			} else {
				bulk_process(stone, count, nil, results_ch, updates_ch)
			}

			wg.Done()
		}
	}
}

func part_len_cache(file_name string) {
	pebbles := read_input(file_name)

	pebble_counts := SafeMap[int, int]{cache: make(map[int]int)}
	ret_buffer := SafeMap[int, int]{cache: make(map[int]int)}
	stone_cache := SafeMap[int, StoneAttr]{cache: make(map[int]StoneAttr)}

	worker_count := 16
	jobs_ch := make(chan StoneJob)
	updates_ch := make(chan UpdatePacket, worker_count)
	processed_ch := make(chan map[int]int, worker_count)

	var main_blocker sync.WaitGroup
	var process_blocker sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for _, v := range pebbles {
		pebble_counts.cache[v] = 1
	}

	for i := 0; i < worker_count; i++ {
		go worker(jobs_ch, processed_ch, updates_ch, &stone_cache, &process_blocker, ctx)
		go cache_updater(&stone_cache, updates_ch, ctx)
		go pebbles_updater(&ret_buffer, processed_ch, &process_blocker, ctx)
	}
	
	for blink := 0; blink < 25; blink++ {
		main_blocker.Add(1)
		go func() {
			for k, v := range pebble_counts.cache {
				if v == 0 {
					continue
				}
				process_blocker.Add(2)
				pebble_counts.cache[k] = 0
				jobs_ch <- StoneJob{stone: k, count: v}

			}
			process_blocker.Wait()
			main_blocker.Done()
		}()
		main_blocker.Wait()
		pebble_counts, ret_buffer = ret_buffer, pebble_counts
	}
	cancel()

	final_pebble_count := 0
	for _, v := range pebble_counts.cache {
		final_pebble_count += v
	}

	fmt.Printf("P.1_len: %d\n", final_pebble_count)

}

func main() {
	file_name := "test.txt"
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
