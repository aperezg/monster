package monster

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

type Monster struct {
	ID      string      `jsonapi:"primary,monsters"`
	Name    string      `jsonapi:"attr,name"`
	Attack  int         `jsonapi:"attr,attack"`
	Defense int         `jsonapi:"attr,defense"`
	Type    MonsterType `json:"type" jsonapi:"attr,type"`
}

func NewMonster() *Monster {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ID := strings.ToUpper(ulid.MustNew(ulid.Timestamp(t), entropy).String())

	return &Monster{ID: ID}
}

type MonsterType string

const (
	WaterType MonsterType = "water"
	FireType              = "fire"
	WindType              = "wind"
	EarthType             = "earth"
)

func allowedMonsterTypes() map[string]MonsterType {
	return map[string]MonsterType{
		string(WaterType): WaterType,
		string(FireType):  FireType,
		string(WindType):  WindType,
		string(EarthType): EarthType,
	}
}

func (t *MonsterType) Check() error {
	if _, ok := allowedMonsterTypes()[string(*t)]; !ok {
		return fmt.Errorf("The type %s is not supported", string(*t))
	}

	return nil
}

const (
	maxAttack  = 999
	maxDefense = 999
)

func (m *Monster) Validate() error {
	if m.Name == "" {
		return errors.New("The name is required")
	}

	if m.Attack > maxAttack {
		return fmt.Errorf("The monster only have max %d attack", maxAttack)
	}

	if m.Defense > maxDefense {
		return fmt.Errorf("The monster only have max % dedense", maxDefense)
	}

	if err := m.Type.Check(); err != nil {
		return err
	}

	return nil
}
