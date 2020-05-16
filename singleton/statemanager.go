package singleton

import "sync"

type manager struct {
	state string
}

var singleton *manager
var once sync.Once

// GetManager return manage object
func GetManager() *manager {
	once.Do(func() {
		singleton = &manager{state: "off"}
	})
	return singleton
}

func (sm *manager) GetState() string {
	return sm.state
}

func (sm *manager) SetState(s string) {
	sm.state = s
}
