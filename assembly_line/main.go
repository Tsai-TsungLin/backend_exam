package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	ID             int
	ProcessedCount int
	mu             sync.Mutex
}

func (e *Employee) IncrementCount() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ProcessedCount++
}

func (e *Employee) GetCount() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ProcessedCount
}

type Item1 struct {
	ID int
}

func (i *Item1) Process() {
	time.Sleep(100 * time.Millisecond)
}

func (i *Item1) String() string {
	return fmt.Sprintf("Item1 #%d", i.ID)
}

type Item2 struct {
	ID int
}

func (i *Item2) Process() {
	time.Sleep(150 * time.Millisecond)
}

func (i *Item2) String() string {
	return fmt.Sprintf("Item2 #%d", i.ID)
}

type Item3 struct {
	ID int
}

func (i *Item3) Process() {
	time.Sleep(200 * time.Millisecond)
}

func (i *Item3) String() string {
	return fmt.Sprintf("Item3 #%d", i.ID)
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
	String() string
}

func main() {
	startTime := time.Now()

	// 創建物品
	items := make([]Item, 0, 30)
	for i := 0; i < 10; i++ {
		items = append(items, &Item1{ID: i + 1})
	}
	for i := 0; i < 10; i++ {
		items = append(items, &Item2{ID: i + 1})
	}
	for i := 0; i < 10; i++ {
		items = append(items, &Item3{ID: i + 1})
	}

	// 隨機打亂
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	// 創建任務 channel
	itemChan := make(chan Item, len(items))
	for _, item := range items {
		itemChan <- item
	}
	close(itemChan)

	// 創建員工
	employees := make([]*Employee, 5)
	for i := 0; i < 5; i++ {
		employees[i] = &Employee{ID: i + 1}
	}

	// 啟動員工 goroutines
	var wg sync.WaitGroup
	for _, emp := range employees {
		wg.Add(1)
		go func(e *Employee) {
			defer wg.Done()
			for item := range itemChan {
				processStart := time.Now()
				fmt.Printf("[%s] 員工 #%d 開始處理 %s\n",
					processStart.Format("2006-01-02 15:04:05.000"),
					e.ID,
					item.String())

				item.Process()

				processEnd := time.Now()
				duration := processEnd.Sub(processStart)
				fmt.Printf("[%s] 員工 #%d 完成處理 %s (耗時: %v)\n",
					processEnd.Format("2006-01-02 15:04:05.000"),
					e.ID,
					item.String(),
					duration)

				e.IncrementCount()
			}
		}(emp)
	}

	wg.Wait()

	totalTime := time.Since(startTime)

	// 統計輸出
	fmt.Println("\n========== 統計結果 ==========")
	fmt.Printf("總處理時間: %v\n", totalTime)
	totalProcessed := 0
	for _, emp := range employees {
		count := emp.GetCount()
		fmt.Printf("員工 #%d 處理了 %d 件物品\n", emp.ID, count)
		totalProcessed += count
	}
	fmt.Printf("總共處理: %d 件物品\n", totalProcessed)
}
