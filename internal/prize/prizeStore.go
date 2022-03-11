package prize

import (
	"fmt"
	"sync"
)

type PrizeStore struct {
	mutex sync.Mutex

	prizes map[int]Prize
	nextId int
}

type ResponsePrize struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func New() *PrizeStore {
	store := &PrizeStore{}
	store.prizes = make(map[int]Prize)
	store.nextId = 1

	return store
}

func (st *PrizeStore) CreatePrize(description string) ResponsePrize {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	prize := Prize{
		id:          st.nextId,
		description: description,
	}

	fmt.Println(prize.id)
	fmt.Println(prize.description)

	st.prizes[st.nextId] = prize
	st.nextId++

	return GetResponsePrize(prize)
}

func (st *PrizeStore) GetPrizes() []ResponsePrize {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	prizes := make([]ResponsePrize, 0, len(st.prizes))
	for _, prize := range st.prizes {
		prizes = append(prizes, GetResponsePrize(prize))
	}

	return prizes
}

func GetResponsePrize(prize Prize) ResponsePrize {
	return ResponsePrize{
		Id:          prize.id,
		Description: prize.description,
	}
}
