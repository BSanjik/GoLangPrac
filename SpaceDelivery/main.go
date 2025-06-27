package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	dockNums  = 25
	orderNums = 50
)

type Result struct {
	OrderId int
	Success bool
}

type Order struct {
	ID int
}

func main() {
	startTime := time.Now()

	wg := sync.WaitGroup{}
	orderChan := make(chan Order)
	resultChan := make(chan Result)

	for i := 1; i <= dockNums; i++ {
		wg.Add(1)
		go acceptOrder(i, orderChan, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	go func() {
		for i := 1; i <= orderNums; i++ {
			orderChan <- Order{ID: i}
		}
		close(orderChan)
	}()

	successCount := 0
	timeoutCount := 0

	for res := range resultChan {
		if res.Success {
			fmt.Printf("✅ Заказ #%02d доставлен\n", res.OrderId)
			successCount++
		} else {
			fmt.Printf("❌ Заказ #%02d просрочен\n", res.OrderId)
			timeoutCount++
		}
	}

	fmt.Printf("📦 Итого:\n")
	fmt.Printf("Успешных доставок: %d\n", successCount)
	fmt.Printf("Просрочено: %d\n", timeoutCount)

	endTime := time.Since(startTime)
	fmt.Printf("Время обработки %v секунд\n", endTime.Seconds())
}

func acceptOrder(id int, orderChan chan Order, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orderChan {
		fmt.Printf("🚀 Получен заказ#%02d\n", order.ID)
		fmt.Printf("🛠️ Док #%02d обрабатывает заказ #%02d...\n", id, order.ID)

		ctx, cancle := context.WithTimeout(context.Background(), time.Second*2)

		orderSuccess := make(chan bool)
		go func() {
			workTime := time.Second * time.Duration(rand.IntN(3)+1)
			time.Sleep(workTime)
			orderSuccess <- true
		}()

		select {
		case <-ctx.Done():
			resultChan <- Result{OrderId: order.ID, Success: false}
		case <-orderSuccess:
			resultChan <- Result{OrderId: order.ID, Success: true}
		}
		defer cancle()
	}
}
