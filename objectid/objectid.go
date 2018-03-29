package objectid

//mongodb的objectid算法

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

var staticMachine = getMachineHash()
var staticIncrement = getRandomNumber()
var staticPid = int32(os.Getpid())

type ObjectId struct {
	timestamp int64
	machine   int32
	pid       int32
	increment int32
}

func New() *ObjectId {
	timestamp := time.Now().Unix()
	return &ObjectId{timestamp, staticMachine, staticPid, atomic.AddInt32(&staticIncrement, 1) & 0xffffff}
}

func Parse(input string) (*ObjectId, error) {
	if len(input) == 0 {
		return nil, errors.New("输入为空")
	}
	if value, ok := tryParse(input); ok {
		return value, nil
	}
	return nil, errors.New("objectID无效")
}

func (this *ObjectId) Timestamp() int64 {
	return this.timestamp
}

func (this *ObjectId) Machine() int32 {
	return this.machine
}

func (this *ObjectId) Pid() int32 {
	return this.pid
}

func (this *ObjectId) Increment() int32 {
	return this.increment & 0xffffff
}

func (this *ObjectId) CreationTime() time.Time {
	return time.Unix(this.timestamp, 0)
}

func (this *ObjectId) Equal(other *ObjectId) bool {
	return this.timestamp == other.timestamp &&
		this.machine == other.machine &&
		this.pid == other.pid &&
		this.increment == other.increment
}

func (this *ObjectId) String() string {
	array := []byte{
		byte(this.timestamp >> 0x18),
		byte(this.timestamp >> 0x10),
		byte(this.timestamp >> 8),
		byte(this.timestamp),
		byte(this.machine >> 0x10),
		byte(this.machine >> 8),
		byte(this.machine),
		byte(this.pid >> 8),
		byte(this.pid),
		byte(this.increment >> 0x10),
		byte(this.increment >> 8),
		byte(this.increment),
	}
	return hex.EncodeToString(array)
}

func getMachineHash() int32 {
	machineName, err := os.Hostname()
	if err != nil {
		machineName = "localhost"
	}
	buf := md5.Sum([]byte(machineName))
	return (int32(buf[0])<<0x10 + int32(buf[1])<<8) + int32(buf[2])
}

func getRandomNumber() int32 {
	rand.Seed(rand.Int63n(time.Now().UnixNano()))
	return rand.Int31()
}

func tryParse(input string) (*ObjectId, bool) {
	if len(input) != 0x18 {
		return nil, false
	}
	array, err := hex.DecodeString(input)
	if err != nil {
		return nil, false
	}
	return &ObjectId{
		timestamp: int64(array[0])<<0x18 + int64(array[1])<<0x10 + int64(array[2])<<8 + int64(array[3]),
		machine:   int32(array[4])<<0x10 + int32(array[5])<<8 + int32(array[6]),
		pid:       int32(array[7])<<8 + int32(array[8]),
		increment: int32(array[9])<<0x10 + (int32(array[10]) << 8) + int32(array[11]),
	}, true
}
