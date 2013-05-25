package bloom

import (
	"fmt"
	"crypto/sha256"
	"math"
	"bytes"
	"encoding/binary"
)

type Filter struct {
	bitCount      int
	hashFunctions int
	chunkSize     int
	bitArray      []byte
}

func New(capacity int, errorRate float64) Filter {
	hashFunctions := int(math.Floor(math.Log2(1.0 / errorRate)))
	bitCount := int(math.Pow(2, math.Ceil(math.Log2(float64(-capacity)*math.Log(errorRate)/math.Pow(math.Log(2), 2)))))
	addressBits := math.Log2(float64(bitCount))
	chunkSize := int(math.Ceil(float64(addressBits / 8)))
	bitArray := make([]byte, bitCount>>3)

	return Filter{
		bitCount,
		hashFunctions,
		chunkSize,
		bitArray,
	}
}

func (f *Filter) Add(key []byte) {
	hash := f.GenerateHashForKey(key)

	for index := 0; index < f.hashFunctions; index++ {
		unpackedBytes := hash[index*f.chunkSize : (index+1)*f.chunkSize]
		unpackedBytesBuffer := bytes.NewBuffer(unpackedBytes)
		unpackedValue, _ := binary.ReadUvarint(unpackedBytesBuffer)
		subHash := unpackedValue % uint64(f.bitCount)

        fmt.Println("unpackedValue", unpackedValue % uint64(f.bitCount), "subHash", subHash)

		word := subHash >> 3
		f.bitArray[word] = f.bitArray[word] | 1<<uint(subHash%8)
	}
}

func (f *Filter) Contains(key []byte) bool {
	hash := f.GenerateHashForKey(key)

	for index := 0; index < f.hashFunctions; index++ {
		unpackedBytes := hash[index*f.chunkSize : (index+1)*f.chunkSize]
		unpackedBytesBuffer := bytes.NewBuffer(unpackedBytes)
		unpackedValue, _ := binary.ReadUvarint(unpackedBytesBuffer)
		subHash := unpackedValue % uint64(f.bitCount)

		word := subHash >> 3
		if f.bitArray[word]&(1<<uint(subHash%8)) == 0 {
			return false
		}
	}

	return true
}

func (f *Filter) GenerateHashForKey(key []byte) []byte {
	var completeHash []byte
	var hash = sha256.New()

	hash.Write(key)
	completeHash = append(completeHash, hash.Sum(nil)...)

	for f.chunkSize*f.hashFunctions > len(completeHash) {
		hash.Reset()
		hash.Write(completeHash)
		completeHash = append(completeHash, hash.Sum(nil)...)
	}

	return completeHash
}
