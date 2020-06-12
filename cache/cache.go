package cache

import "sync"

type Item interface {
	GetId() string
}

type Pool struct {
	mux  sync.Mutex
	pool map[string]Item
}

func NewPool() *Pool {
	return &Pool{pool: make(map[string]Item)}
}

func (p *Pool) Add(item Item) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pool[item.GetId()] = item
}

func (p *Pool) Get(id string) Item {
	p.mux.Lock()
	defer p.mux.Unlock()

	return p.pool[id]
}

func (p *Pool) GetAll() []Item {
	p.mux.Lock()
	defer p.mux.Unlock()

	items := make([]Item, 0, len(p.pool))

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

	p.pool = make(map[string]Item)
}
