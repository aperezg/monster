package storage

import (
	"fmt"
	"sync"

	"github.com/aperezg/monster"
)

type MonsterRepository struct {
	mtx      sync.RWMutex
	monsters map[string]*monster.Monster
}

func NewMonsterRepository(monsters map[string]*monster.Monster) *MonsterRepository {
	if monsters == nil {
		monsters = make(map[string]*monster.Monster)
	}

	return &MonsterRepository{
		monsters: monsters,
	}
}

func (r *MonsterRepository) CreateMonster(m *monster.Monster) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExistsByName(m.Name); err != nil {
		return err
	}
	r.monsters[m.ID] = m
	return nil
}

func (r *MonsterRepository) FetchMonsters() ([]*monster.Monster, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*monster.Monster, 0, len(r.monsters))
	for _, value := range r.monsters {
		values = append(values, value)
	}
	return values, nil
}

func (r *MonsterRepository) DeleteMonster(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.monsters, ID)

	return nil
}

func (r *MonsterRepository) UpdateMonster(ID string, m *monster.Monster) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.monsters[ID] = m
	return nil
}

func (r *MonsterRepository) FetchMonsterByID(ID string) (*monster.Monster, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.monsters {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}

func (r *MonsterRepository) checkIfExistsByName(name string) error {
	for _, v := range r.monsters {
		if v.Name == name {
			return fmt.Errorf("The monster %s is already exist", name)
		}
	}

	return nil
}
