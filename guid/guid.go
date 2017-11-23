package guid

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	timestampBits uint = 30
	versionBits   uint = 2
	nodeBits      uint = 4
	sequenceBits  uint = 12

	sequenceMask = -1 ^ (-1 << sequenceBits)

	timestampShift uint = 34
	versionShift   uint = 32
	nodeShift      uint = 28
	sequenceShift  uint = 16

	epoch int64 = 1487235352
)

// Generator ...
//
// |-----------------|-----|------|--------|-------------|
// |        30       |  2  |  4   |   12   |     16      |
// |-----------------|-----|------|--------|-------------|
// |     timestamp   | ver | node |  seq   |    blank    |
// |-----------------|-----|------|--------|-------------|
type Generator struct {
	time    int64
	version int64
	node    int64
	seq     int64
	sync.Mutex
}

// NewGenerator ...
func NewGenerator(node int64) (*Generator, error) {
	if node < 0 || node > (1<<nodeBits) {
		return nil, errors.New("invalid node id")
	}
	return &Generator{
		time:    0,
		version: 1,
		node:    node,
		seq:     0,
	}, nil
}

// ID ...
type ID int64

// Generate ...
func (g *Generator) Generate() ID {
	g.Lock()

	curSec := time.Now().UnixNano() / 1000000000
	if g.time == curSec {
		g.seq = (g.seq + 1) & int64(sequenceMask)
		if g.seq == 0 {
			for curSec <= g.time {
				time.Sleep(50 * time.Nanosecond)
				curSec = time.Now().UnixNano() / 1000000000
			}
		}
	} else {
		g.seq = 0
	}

	g.time = curSec

	id := ID((curSec-epoch)<<timestampBits |
		(g.version << versionShift) |
		(g.node << nodeShift) |
		(g.seq << sequenceShift))

	g.Unlock()
	return id
}
