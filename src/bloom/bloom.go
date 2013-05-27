package bloom

import (
	"crypto/sha256"
	"math"
	"math/big"
)

type Filter struct {
	Capacity      int
	ErrorRate     float64
	bitTotal      uint64
	bitArray      []byte
	hashFunctions int
	chunkSize     int
}

func New(capacity int, errorRate float64) Filter {
	bitTotal := uint64(math.Pow(2, math.Ceil(math.Log2(-float64(capacity)*math.Log(errorRate)/math.Pow(math.Log(2), 2)))))
	bitArray := make([]byte, bitTotal>>3)
	hashFunctions := int(math.Floor(math.Log2(1.0 / errorRate)))
	chunkSize := int(math.Ceil(math.Log2(float64(bitTotal)) / 8))

	return Filter{
		capacity,
		errorRate,
		bitTotal,
		bitArray,
		hashFunctions,
		chunkSize,
	}
}

func (f *Filter) Add(key []byte) {
	hash := f.generateHashForKey(key)

	for index := 0; index < f.hashFunctions; index++ {
		byteIndex, bitMask := f.findBitArrayIndexAndMask(hash, index)

		f.bitArray[byteIndex] |= bitMask
	}

	// fmt.Printf("\n")
}

func (f *Filter) Contains(key []byte) bool {
	hash := f.generateHashForKey(key)

	for index := 0; index < f.hashFunctions; index++ {
		byteIndex, bitMask := f.findBitArrayIndexAndMask(hash, index)

		if f.bitArray[byteIndex]&bitMask == 0 {
			return false
		}
	}

	return true
}

func (f *Filter) generateHashForKey(key []byte) []byte {
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

func (f *Filter) findBitArrayIndexAndMask(hash []byte, index int) (bitArrayIndex uint64, mask byte) {
	var subHashInt big.Int

	subHashBytes := hash[index*f.chunkSize : (index+1)*f.chunkSize]
	subHashInt.SetBytes(subHashBytes)
	position := subHashInt.Uint64()

	position %= f.bitTotal

	bitArrayIndex = uint64(position >> 3)
	mask = 1 << uint(position%8)
	return
}

func (f *Filter) GetSizeInBytes() uint64 {
	return f.bitTotal >> 3
}
