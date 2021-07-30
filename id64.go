package id64

import (
	"sync/atomic"
	"time"

	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
)

var lastSec uint32
var Gen *Generator
var offset2021 uint32 // MinDateOffset = Offset2020.Unix() or MinNanoDateOffset = Offset2020.UnixNano()

func init() {

	offset2021 = uint32(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC).Unix())

	Gen = &Generator{}
}

func ID() id64 {
	return Gen.ID()
}

type Generator struct {
	AtomicCounter uint32
}

func (gen *Generator) ID() id64 {
	now := uint32(fastime.UnixNow())
	counter := atomic.AddUint32(&gen.AtomicCounter, 1)
	if now != lastSec {
		atomic.SwapUint32(&lastSec, now)
		atomic.SwapUint32(&gen.AtomicCounter, 0) // ignore old value
		counter = atomic.AddUint32(&gen.AtomicCounter, 1)
	}
	return id64((uint64(lastSec-offset2021) << 32) + uint64(counter))
}

type id64 uint64

func (i id64) Time() time.Time {
	return time.Unix(int64(i>>32)+int64(offset2021), 0)
}

func (i id64) Counter() uint32 {
	return uint32(i % (256 * 256 * 256 * 256))
}

func (i id64) String() string {
	return S.EncodeCB63(int64(i), 0)
}
