package promo

import (
	"fmt"
	"src/internal/participant"
	"src/internal/prize"
	"sync"
)

type PromoStore struct {
	mutex sync.Mutex

	promos map[int]Promo
	nextId int
}

type ResponsePromo struct {
	Id           int                       `json:"id"`
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	Prizes       []prize.Prize             `json:"prizes"`
	Participants []participant.Participant `json:"participants"`
}

func New() *PromoStore {
	store := &PromoStore{}
	store.promos = make(map[int]Promo)
	store.nextId = 1

	return store
}

func (st *PromoStore) CreatePromo(name string, description string) ResponsePromo {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	promo := Promo{
		Id:           st.nextId,
		Name:         name,
		Description:  description,
		Prizes:       make([]prize.Prize, 0),
		Participants: make([]participant.Participant, 0),
	}

	st.promos[st.nextId] = promo
	st.nextId++

	return GetResponsePromo(&promo)
}

func (st *PromoStore) GetProducts() []ResponsePromo {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	promos := make([]ResponsePromo, 0, len(st.promos))
	for _, promo := range st.promos {
		promos = append(promos, GetResponsePromo(&promo))
	}

	return promos
}

func (st *PromoStore) GetPromo(id int) (ResponsePromo, error) {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	prod, ok := st.promos[id]
	if ok {
		return GetResponsePromo(&prod), nil
	} else {
		return ResponsePromo{}, fmt.Errorf("There is no promo with id=%d", id)
	}
}

func (st *PromoStore) DeletePromo(id int) error {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	if _, ok := st.promos[id]; !ok {
		return fmt.Errorf("There is no product with id=%d", id)
	}
	delete(st.promos, id)
	return nil
}

func GetResponsePromo(promo *Promo) ResponsePromo {
	return ResponsePromo{
		Id:           promo.Id,
		Name:         promo.Name,
		Description:  promo.Description,
		Prizes:       promo.Prizes,
		Participants: promo.Participants,
	}
}
