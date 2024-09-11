package ratel

import (
	. "nostr.mleku.dev"
	"ratel.mleku.dev/keys/index"
	"ratel.mleku.dev/keys/serial"
)

// GetCounterKey returns the proper counter key for a given event ID.
func GetCounterKey(ser *serial.T) (key B) {
	key = index.Counter.Key(ser)
	// Log.T.F("counter key %d %d", index.Counter, ser.Uint64())
	return
}
