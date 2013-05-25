package bloom

import (
	"fmt"
	"crypto/sha256"
	"math"
	// "math/big"
	// "strconv"
	"bytes"
	"encoding/binary"
)

type Filter struct {
	bitCount      int
	hashFunctions int
	bitMask       int
	chunkSize     int
	bitArray      []byte
}

func New(capacity int, errorRate float64) Filter {
	hashFunctions := int(math.Floor(math.Log2(1.0 / errorRate)))
	bitCount := int(math.Pow(2, math.Ceil(math.Log2(float64(-capacity)*math.Log(errorRate)/math.Pow(math.Log(2), 2)))))
	addressBits := math.Log2(float64(bitCount))
	bitMask := (1 << uint(addressBits)) - 1
	chunkSize := int(math.Ceil(float64(addressBits / 8)))
	bitArray := make([]byte, bitCount>>3)

	return Filter{
		bitCount,
		hashFunctions,
		bitMask,
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

		subHash := (unpackedValue % uint64(f.bitCount)) & uint64(f.bitMask)

        fmt.Println("unpackedValue", unpackedValue % uint64(f.bitCount), "subHash", subHash)

		word := subHash >> 3
		// fmt.Printf("%d-", f.bitArray[word])
		f.bitArray[word] = f.bitArray[word] | 1<<uint(subHash%8)
		// fmt.Printf("%d ", f.bitArray[word])
	}

	// fmt.Println()
}

// var word *big.Int = new(big.Int)
// var modulus *big.Int = new(big.Int)

// i := new(big.Int)
// unpackedValue := hash[index*f.chunkSize : (index+1)*f.chunkSize]
// unpackedValueBuf := bytes.NewBuffer(unpackedValue)
// in, _ := binary.ReadUvarint(unpackedValueBuf)
// fmt.Println(in)
// subHash := (i.SetBytes(unpackedValue).Mod(i, big.NewInt(int64(f.bitCount)))).And(i, big.NewInt(int64(f.bitMask)))

// word.DivMod(subHash, big.NewInt(8), modulus)

// bitIndex, _ := strconv.ParseInt(word.String(), 10, 64)
// toShift, _ := strconv.ParseInt(modulus.String(), 10, 8)
// f.bitArray[bitIndex] = f.bitArray[bitIndex] | 1<<uint(toShift)

func (f *Filter) Contains(key []byte) bool {
	hash := f.GenerateHashForKey(key)

	for index := 0; index < f.hashFunctions; index++ {
		unpackedBytes := hash[index*f.chunkSize : (index+1)*f.chunkSize]
		unpackedBytesBuffer := bytes.NewBuffer(unpackedBytes)
		unpackedValue, _ := binary.ReadUvarint(unpackedBytesBuffer)
		subHash := (unpackedValue % uint64(f.bitCount)) & uint64(f.bitMask)

		word := subHash >> 3
		// fmt.Println("compare", (f.bitArray[word] & (1<<uint(subHash % 8))))
		if f.bitArray[word]&(1<<uint(subHash%8)) == 0 {
			// fmt.Println("bitArray[word]", f.bitArray[word], "shifted", 1<<uint(subHash % 8))
			return false
		}

		// var word *big.Int = new(big.Int)
		// var modulus *big.Int = new(big.Int)

		// i := new(big.Int)
		// unpackedValue := hash[index*f.chunkSize : (index+1)*f.chunkSize]
		//       // fmt.Printf("%x %d", unpackedValue, unpackedValue);
		// subHash := (i.SetBytes(unpackedValue).Mod(i, big.NewInt(int64(f.bitCount)))).And(i, big.NewInt(int64(f.bitMask)))

		// word.DivMod(subHash, big.NewInt(8), modulus)

		// bitIndex, _ := strconv.ParseInt(word.String(), 10, 64)
		// toShift, _ := strconv.ParseInt(modulus.String(), 10, 8)

		// if (f.bitArray[bitIndex] & (1 << uint(toShift))) == 0 {
		// 	return false
		// }
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
