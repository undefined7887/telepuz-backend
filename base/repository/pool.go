package repository

import "sync"

type Pool struct {
	mux  sync.Mutex
	pool map[string]PoolItem
}

type PoolItem interface {
	GetId() string
}

func NewPool() *Pool {
	return &Pool{pool: make(map[string]PoolItem)}
}

func (p *Pool) Add(item PoolItem) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pool[item.GetId()] = item
}

func (p *Pool) Get(id string) PoolItem {
	p.mux.Lock()
	defer p.mux.Unlock()

	return p.pool[id]
}

func (p *Pool) GetAll() []PoolItem {
	p.mux.Lock()
	defer p.mux.Unlock()

	items := make([]PoolItem, 0, len(p.pool))

	for _, val := range p.pool {
		items = append(items, val)
	}

	return items
}

func (p *Pool) Remove(id string) {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.pool, id)
}

func (p *Pool) RemoveAll() {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pool = make(map[string]PoolItem)
}
