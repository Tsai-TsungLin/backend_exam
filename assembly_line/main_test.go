package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// TestItemInterface 驗證所有 Item 類型實現 interface
func TestItemInterface(t *testing.T) {
	var _ Item = &Item1{}
	var _ Item = &Item2{}
	var _ Item = &Item3{}
}

// TestItemProcessingTime 驗證每種 Item 的處理時間
func TestItemProcessingTime(t *testing.T) {
	tests := []struct {
		name     string
		item     Item
		expected time.Duration
		tolerance time.Duration
	}{
		{"Item1", &Item1{ID: 1}, 100 * time.Millisecond, 10 * time.Millisecond},
		{"Item2", &Item2{ID: 1}, 150 * time.Millisecond, 10 * time.Millisecond},
		{"Item3", &Item3{ID: 1}, 200 * time.Millisecond, 10 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			tt.item.Process()
			duration := time.Since(start)

			if duration < tt.expected-tt.tolerance || duration > tt.expected+tt.tolerance {
				t.Errorf("%s processing time = %v, want %v (±%v)",
					tt.name, duration, tt.expected, tt.tolerance)
			}
		})
	}
}

// TestEmployeeIncrementCount 驗證員工計數器的正確性
func TestEmployeeIncrementCount(t *testing.T) {
	emp := &Employee{ID: 1}

	if count := emp.GetCount(); count != 0 {
		t.Errorf("Initial count = %d, want 0", count)
	}

	emp.IncrementCount()
	if count := emp.GetCount(); count != 1 {
		t.Errorf("After increment count = %d, want 1", count)
	}

	emp.IncrementCount()
	emp.IncrementCount()
	if count := emp.GetCount(); count != 3 {
		t.Errorf("After 3 increments count = %d, want 3", count)
	}
}

// TestEmployeeConcurrentCount 驗證員工計數器的並發安全性
func TestEmployeeConcurrentCount(t *testing.T) {
	emp := &Employee{ID: 1}
	var wg sync.WaitGroup
	iterations := 1000

	// 並發增加計數
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			emp.IncrementCount()
		}()
	}

	wg.Wait()

	if count := emp.GetCount(); count != iterations {
		t.Errorf("Concurrent count = %d, want %d", count, iterations)
	}
}

// TestAssemblyLine_ItemCount 驗證處理所有30件物品
func TestAssemblyLine_ItemCount(t *testing.T) {
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
				item.Process()
				e.IncrementCount()
			}
		}(emp)
	}

	wg.Wait()

	// 驗證總處理數量
	totalProcessed := 0
	for _, emp := range employees {
		count := emp.GetCount()
		totalProcessed += count
	}

	if totalProcessed != 30 {
		t.Errorf("Total processed items = %d, want 30", totalProcessed)
	}
}

// TestAssemblyLine_EmployeeDistribution 驗證員工分配的合理性
func TestAssemblyLine_EmployeeDistribution(t *testing.T) {
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
				item.Process()
				e.IncrementCount()
			}
		}(emp)
	}

	wg.Wait()

	// 驗證每個員工都處理了物品
	for _, emp := range employees {
		count := emp.GetCount()
		if count == 0 {
			t.Errorf("員工 #%d 沒有處理任何物品", emp.ID)
		}
		if count > 30 {
			t.Errorf("員工 #%d 處理了 %d 件物品，超過總數", emp.ID, count)
		}
	}

	// 驗證總和
	totalProcessed := 0
	for _, emp := range employees {
		totalProcessed += emp.GetCount()
	}

	if totalProcessed != 30 {
		t.Errorf("所有員工處理總和 = %d, want 30", totalProcessed)
	}
}

// TestAssemblyLine_ProcessingTime 驗證總處理時間合理性
func TestAssemblyLine_ProcessingTime(t *testing.T) {
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
				item.Process()
				e.IncrementCount()
			}
		}(emp)
	}

	wg.Wait()

	totalTime := time.Since(startTime)

	// 計算理論最短時間：總時間 / 員工數
	// Item1: 10 * 100ms = 1000ms
	// Item2: 10 * 150ms = 1500ms
	// Item3: 10 * 200ms = 2000ms
	// 總計: 4500ms
	// 5個員工並發，理論最短約 900ms (4500/5)
	minExpectedTime := 900 * time.Millisecond

	// 最長時間：所有物品串行處理
	maxExpectedTime := 5000 * time.Millisecond

	if totalTime < minExpectedTime {
		t.Errorf("總處理時間 %v 太短，可能測試有誤", totalTime)
	}

	if totalTime > maxExpectedTime {
		t.Errorf("總處理時間 %v 超過預期最大值 %v", totalTime, maxExpectedTime)
	}
}

// TestAssemblyLine_RaceCondition 驗證無 race condition（需使用 -race 執行）
func TestAssemblyLine_RaceCondition(t *testing.T) {
	// 多次執行以增加檢測 race condition 的機會
	for run := 0; run < 5; run++ {
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
					item.Process()
					e.IncrementCount()
					// 同時讀取以增加 race 檢測機會
					_ = e.GetCount()
				}
			}(emp)
		}

		wg.Wait()
	}
}
