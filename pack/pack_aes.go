package pack

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	. "qora/global"
	. "qora/utils"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// PackAES function
// input source file list and dest package path, output error information
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// dest file also support both absolute and relative paths, like 'C:\\package.pak' or '../test/data/package.pak'
// dest file name suffix can be any type such as '.pak', '.dat', even none is ok
// return err indicate the success or failure function execute
func PackAES(src []string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// clear global variable
	atomic.StoreInt64(&Done, 0)
	// first, split the pre-crypt files
	r := make([][]byte, len(src)+4)
	for k, v := range src {
		wg.Add(1)
		go PackAESOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			s := fmt.Sprintf("Error aes pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPackAES{}
	head.Name = make([]byte, 32)
	head.Author = make([]byte, 16)
	head.Type = make([]byte, 8)
	head.Number = make([]byte, 4)
	_, name := filepath.Split(dest)
	if len([]byte(name)) > 32 {
		log.Println("Error dest file name length:", err)
		return
	}
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Author), []byte("Alopex6414"))
	BytesCopy(&(head.Type), []byte("AES"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(dest, s, 0644)
	if err != nil {
		log.Println("Error write aes file:", err)
	}
	return err
}

// PackAESConfine function
// it common with function PackAES, just restrict goroutine when running
func PackAESConfine(src []string, dest string) (err error) {
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, 5)
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// clear global variable
	atomic.StoreInt64(&Done, 0)
	// first, split the pre-crypt files
	r := make([][]byte, len(src)+4)
	for k, v := range src {
		wg.Add(1)
		ch <- struct{}{}
		go PackAESOneConfineGo(v, &r[k+4], wg, ch)
		//go PackAESOneGo(v, &r[k+4], wg)
	}
	wg.Wait()
	// second, check goroutine whether success or not
	for i := 0; i < len(src); i++ {
		if bytes.Equal(r[i+4], []byte("")) {
			s := fmt.Sprintf("Error aes pack one file: %v", src[i])
			err = errors.New(s)
			return err
		}
	}
	// third, fill the header
	head := TPackAES{}
	head.Name = make([]byte, 32)
	head.Author = make([]byte, 16)
	head.Type = make([]byte, 8)
	head.Number = make([]byte, 4)
	_, name := filepath.Split(dest)
	if len([]byte(name)) > 32 {
		log.Println("Error dest file name length:", err)
		return
	}
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Author), []byte("Alopex6414"))
	BytesCopy(&(head.Type), []byte("AES"))
	BytesCopy(&(head.Number), IntToBytes(len(src)))
	r[0] = head.Name
	r[1] = head.Author
	r[2] = head.Type
	r[3] = head.Number
	// finally, write to dest file
	s := bytes.Join(r, []byte(""))
	err = ioutil.WriteFile(dest, s, 0644)
	if err != nil {
		log.Println("Error write aes file:", err)
	}
	return err
}

// PackAESWorkCalculate function
// it will calculate the total work value which you input files
// it will be call in progress pack files
// input src files same as you pack src files
// output work value and err
// return err indicate the success or failure function execute
func PackAESWorkCalculate(src []string) (work int64, err error) {
	var sum int64
	if len(src) == 0 {
		err = errors.New("Pack file list is empty.")
		return work, err
	}
	for _, v := range src {
		var size int64
		err = filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			size = info.Size()
			return err
		})
		if err != nil {
			log.Println("Error calculate work:", err)
			return work, err
		}
		if size%AESBufferSize != 0 {
			padding := AESBufferSize - size%AESBufferSize
			size += padding
		}
		sum += size
	}
	work = sum
	return work, err
}

// PackAESOneGo function
// input source file, return value pointer and wait group pointer
// it will pack one file through goroutine
// src file support both absolute and relative paths, like 'C:\\file.txt' or '../test/data/file.txt'
// r should input byte slice pointer and it will fill in return value
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func PackAESOneGo(src string, r *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*r, err = PackAESOne(src)
	if err != nil {
		log.Println("Error aes pack one file:", err)
		return err
	}
	return err
}

// PackAESOneConfineGo function
// it common with function PackAESOneGo, just restrict goroutine when running
func PackAESOneConfineGo(src string, r *[]byte, wg *sync.WaitGroup, ch chan interface{}) (err error) {
	defer wg.Done()
	*r, err = PackAESOneConfine(src)
	if err != nil {
		log.Println("Error aes pack one file:", err)
		<-ch
		return err
	}
	<-ch
	return err
}

// PackAESOne function
// it the base function of PackAESOneGo
func PackAESOne(src string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return r, err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return r, err
	}
	// third, generate random key
	key := make([]byte, 16)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, AESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call AESEncrypt function
	wg := &sync.WaitGroup{}
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		go AESEncryptGo(v, key, &rr[k], wg)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	if len(key) > 16 {
		log.Println("Error key length:", err)
		return
	}
	head := TPackAESOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 16)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Key), key)
	BytesCopy(&(head.OriginSize), IntToBytes(len(data)))
	BytesCopy(&(head.CryptSize), IntToBytes(len(dest)))
	/*// fourth, we can call AESEncrypt function
	dest, err := AESEncrypt(data, key)
	if err != nil {
		log.Println("Error AES Encrypt data:", err)
		return r, err
	}*/
	// finally, return result
	var s [][]byte
	s = append(s, head.Name)
	s = append(s, head.Key)
	s = append(s, head.OriginSize)
	s = append(s, head.CryptSize)
	s = append(s, dest)
	r = bytes.Join(s, []byte(""))
	return r, err
}

// PackAESOneConfineGo function
// it the base function of PackAESOneConfineGo
func PackAESOneConfine(src string) (r []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return r, err
	}
	defer file.Close()
	// second, read file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error read file:", err)
		return r, err
	}
	// third, generate random key
	key := make([]byte, 16)
	_, err = rand.Read(key)
	if err != nil {
		log.Println("Error generate random key:", err)
		return r, err
	}
	// fourth, split the data slice
	ss, err := SplitByte(data, AESBufferSize)
	if err != nil {
		log.Println("Error split bytes:", err)
		return r, err
	}
	// fifth, we can call AESEncrypt function
	wg := &sync.WaitGroup{}
	ch := make(chan interface{}, 8192)
	rr := make([][]byte, len(ss))
	for k, v := range ss {
		wg.Add(1)
		ch <- struct{}{}
		go AESEncryptConfineGo(v, key, &rr[k], wg, ch)
	}
	wg.Wait()
	dest := bytes.Join(rr, []byte(""))
	// sixth, fill the packet struct
	_, name := filepath.Split(src)
	if len([]byte(name)) > 32 {
		log.Println("Error source file name length:", err)
		return
	}
	if len(key) > 16 {
		log.Println("Error key length:", err)
		return
	}
	head := TPackAESOne{}
	head.Name = make([]byte, 32)
	head.Key = make([]byte, 16)
	head.OriginSize = make([]byte, 4)
	head.CryptSize = make([]byte, 4)
	BytesCopy(&(head.Name), []byte(name))
	BytesCopy(&(head.Key), key)
	BytesCopy(&(head.OriginSize), IntToBytes(len(data)))
	BytesCopy(&(head.CryptSize), IntToBytes(len(dest)))
	/*// fourth, we can call AESEncrypt function
	dest, err := AESEncrypt(data, key)
	if err != nil {
		log.Println("Error AES Encrypt data:", err)
		return r, err
	}*/
	// finally, return result
	var s [][]byte
	s = append(s, head.Name)
	s = append(s, head.Key)
	s = append(s, head.OriginSize)
	s = append(s, head.CryptSize)
	s = append(s, dest)
	r = bytes.Join(s, []byte(""))
	return r, err
}

// AESEncryptGo function
// input source file, encrypt key, return value pointer and wait group pointer
// it will encrypt one file through goroutine
// src file must be byte slice, you can open file and read it through io
// key is a 128bit number which used by aes, here is 16 bit byte slice
// dest should input byte slice pointer and it will fill in return value after encrypt
// wg is a flag to control different goroutine sync
// return err indicate the success or failure function execute
func AESEncryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	*dest, err = AESEncrypt(src, key)
	if err != nil {
		log.Println("Error aes encrypt data:", err)
		return err
	}
	atomic.AddInt64(&Done, 1)
	return err
}

// AESEncryptConfineGo function
// it common with function AESEncryptGo, just restrict goroutine when running
// it will called at the confine mode
func AESEncryptConfineGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup, ch chan interface{}) (err error) {
	defer wg.Done()
	*dest, err = AESEncrypt(src, key)
	if err != nil {
		log.Println("Error aes encrypt data:", err)
		<-ch
		return err
	}
	atomic.AddInt64(&Done, 1)
	<-ch
	return err
}

// AESEncrypt function
// original function of encrypt
func AESEncrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 16, 24, 32
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// fill block data
	src = PKCS7Padding(src, blockSize)
	// encrypt mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// encrypt
	blockMode.CryptBlocks(dest, src)
	return dest, err
}

// PKCS7Padding function
// AESEncrypt will call this function
func PKCS7Padding(src []byte, size int) []byte {
	var padding int
	if len(src)%size != 0 {
		padding = size - len(src)%size
	}
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, text...)
}
