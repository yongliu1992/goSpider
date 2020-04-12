package module

import (
	"math"
	"sync"
)

type SNGenerator interface {
	//Start 用于获取预设的最小序列号。
	Start() uint64
	// Max 用于获取预设的最大序列号。
	Max() uint64
	// Next 用于获取下一个序列号。
	Next() uint64
	// CycleCount 用于获取循环计数。
	CycleCount() uint64
	// Get 用于获得一个序列号并准备下一个序列号。
	Get() uint64
}

type myGenerator struct {
	// start 代表序列号的最小值。
	start uint64
	// max 代表序列号的最大值。
	max uint64
	// next 代表下一个序列号。
	next uint64
	// cycleCount 代表循环的计数。
	cycleCount uint64
	// lock 代表读写锁。
	lock sync.RWMutex
}

func (m myGenerator) Start() uint64 {
	return m.start
}
func (m *myGenerator) Max() uint64 {
	return m.max
}

func (m *myGenerator) Next() uint64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.next
}

func (m *myGenerator) CycleCount() uint64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.cycleCount
}

func (m *myGenerator) Get() uint64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	ID := m.next
	if ID == m.max {
		m.next = m.start
		m.cycleCount++
	} else {
		m.next++
	}
	return ID
}

func NewSNGenerator(start uint64, max uint64) SNGenerator {
	if max == 0 {
		max = math.MaxUint64
	}
	return &myGenerator{start: start, max: max, next: start}
}
