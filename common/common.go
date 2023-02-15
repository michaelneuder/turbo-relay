// Package common provides things used by various other components
package common

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrServerAlreadyRunning = errors.New("server already running")

	SlotsPerEpoch    = 32
	DurationPerSlot  = time.Second * 12
	DurationPerEpoch = DurationPerSlot * time.Duration(SlotsPerEpoch)
)

// HTTPServerTimeouts are various timeouts for requests to the mev-boost HTTP server
type HTTPServerTimeouts struct {
	Read       time.Duration // Timeout for body reads. None if 0.
	ReadHeader time.Duration // Timeout for header reads. None if 0.
	Write      time.Duration // Timeout for writes. None if 0.
	Idle       time.Duration // Timeout to disconnect idle client connections. None if 0.
}

// BuilderStatus configures how builder blocks are processed.
type BuilderStatus struct {
	IsHighPrio    bool
	IsBlacklisted bool
	IsDemoted     bool
}

type Profile struct {
	Unzip       uint64
	Decode      uint64
	CacheRead   uint64
	RandaoLock1 uint64
	DutiesLock  uint64
	Checks      uint64
	RandaoLock2 uint64
	Simulation  uint64
	RedisUpdate uint64
	Submission  uint64
}

func (p *Profile) String() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", p.Unzip, p.Decode, p.CacheRead, p.RandaoLock1, p.DutiesLock, p.Checks, p.RandaoLock2, p.Simulation, p.RedisUpdate, p.Submission)
}
