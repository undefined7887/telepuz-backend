package repository

import "sync"

type Pool struct {
	mux  sync.Mutex
	pool map[string]interface{}
}

func NewPool() *Pool {
	return &Pool{pool: make(map[string]interface{})}
}

func (p *Pool) Add(id string, value interface{}) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pool[id] = value
}

func (p *Pool) Get(id string) interface{} {
	p.mux.Lock()
	defer p.mux.Unlock()

	return p.pool[id]
}

func (p *Pool) GetAll() interface{} {
	p.mux.Lock()
	defer p.mux.Unlock()

	items := make([]interface{}, 0, len(p.pool))

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

	p.pool = make(map[string]interface{})
}
