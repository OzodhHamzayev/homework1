package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Order struct {
	ID 	   int
	Status string
	mu 	   sync.Mutex
}



func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i+1, Status: "pending"}
	}
	return orders
}


func proccessOrders(orders []*Order) {
	for _, order := range orders { 
		time.Sleep(
			time.Duration(rand.IntN(500)) * 
			time.Millisecond,
		)
		fmt.Println(order)
	}
}

func updateOrderStatuses(orders []*Order) {

	for _, order := range orders { 
		time.Sleep(
			time.Duration(rand.IntN(300)) *
			time.Millisecond,
		)
		status := []string{
			"proccessing", "shipped", "delivered",
		}[rand.IntN(3)]
		order.Status = status
		fmt.Println(order.ID, order.Status)
	}

}
func main() {
	var wg sync.WaitGroup
	wg.Add(2) 
	// wg := sync.WaitGroup{} 
	orders :=  generateOrders(20)

	go func ()  {
		defer wg.Done()
		proccessOrders(orders)	
	}()
	go func ()  {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()
	wg.Wait()

	fmt.Println("All operations completed.")
}
