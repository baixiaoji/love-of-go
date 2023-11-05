package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoStorer struct{}

func (m MongoStorer) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MongoStorer) Put(number int) error {
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

// interfaces
func main() {
	apiServer := ApiServer{
		numberStore: MongoStorer{},
	}

	nums, err := apiServer.numberStore.GetAll()

	if err != nil {
		fmt.Println(" get all has the error")
	}

	for i, num := range nums {
		fmt.Printf("index %d, value %d\n", i, num)
	}
}
