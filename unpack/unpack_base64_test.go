package unpack

import (
	"bytes"
	"io/ioutil"
	. "qora/utils"
	"sync"
	"testing"
)

// TestUnpackBase64 function
func TestUnpackBase64(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	err := UnpackBase64(src, dest)
	if err != nil {
		t.Fatal("Error Unpack Base64:", err)
	}
}

// TestUnpackBase64Confine function
func TestUnpackBase64Confine(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	err := UnpackBase64Confine(src, dest)
	if err != nil {
		t.Fatal("Error Unpack Base64:", err)
	}
}

// TestUnpackBase64ToFile function
func TestUnpackBase64ToFile(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackBase64ToFile(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack Base64 To File:", err)
	}
}

// TestUnpackBase64ToFileConfine function
func TestUnpackBase64ToFileConfine(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	dest := "../test/data/unpack/"
	target := "file_1.txt"
	err := UnpackBase64ToFileConfine(src, target, dest)
	if err != nil {
		t.Fatal("Error Unpack Base64 To File:", err)
	}
}

// TestUnpackBase64ToMemory function
func TestUnpackBase64ToMemory(t *testing.T) {
	var dest []byte
	src := "../test/data/unpack/file_base64.txt"
	target := "file_1.txt"
	err := UnpackBase64ToMemory(src, target, &dest)
	if err != nil {
		t.Fatal("Error Unpack Base64 To Memory:", err)
	}
}

// TestBase64ExtractInfo function
func TestBase64ExtractInfo(t *testing.T) {
	var dest []string
	var sz []int
	src := "../test/data/unpack/file_base64.txt"
	err := UnpackBase64ExtractInfo(src, &dest, &sz)
	if err != nil {
		t.Fatal("Error Unpack Base64 Extract Information:", err)
	}
	if len(dest) != 5 {
		t.Fatal("Error Extract file number")
	}
}

// TestUnpackBase64WorkCalculate function
func TestUnpackBase64WorkCalculate(t *testing.T) {
	src := "../test/data/unpack/file_base64.txt"
	_, err := UnpackBase64WorkCalculate(src)
	if err != nil {
		t.Fatal("Error Unpack Base64 Work Calculate:", err)
	}
}

// TestUnpackBase64OneToMemory function
func TestUnpackBase64OneToMemory(t *testing.T) {
	var dest string
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	err = UnpackBase64OneToMemory(s, &dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
	if dest != "hello,world!" {
		t.Fatal("Error unpack content:", dest)
	}
}

// TestUnpackBase64OneGo function
func TestUnpackBase64OneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	dest := "../test/data/unpack/"
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	wg.Add(1)
	go UnpackBase64OneGo(s, h, dest, &wg)
	wg.Wait()
}

// TestUnpackBase64One function
func TestUnpackBase64One(t *testing.T) {
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	dest := "../test/data/unpack/"
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	err = UnpackBase64One(s, h, dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

// TestUnpackBase64OneConfine function
func TestUnpackBase64OneConfine(t *testing.T) {
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	dest := "../test/data/unpack/"
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	err = UnpackBase64OneConfine(s, h, dest)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

// TestBase64DecryptGo function
func TestBase64DecryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r string
	src := "aGVsbG8sd29ybGQh"
	dest := "hello,world!"
	wg.Add(1)
	go Base64DecryptGo(src, &r, &wg)
	wg.Wait()
	if r != dest {
		t.Errorf("Error Decrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

// TestBase64Decrypt function
func TestBase64Decrypt(t *testing.T) {
	src := "aGVsbG8sd29ybGQh"
	dest := "hello,world!"
	r := Base64Decrypt(src)
	if r != dest {
		t.Errorf("Error Decrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

// BenchmarkUnpackBase64 function
func BenchmarkUnpackBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		err := UnpackBase64(src, dest)
		if err != nil {
			b.Fatal("Error Unpack Base64:", err)
		}
	}
}

// BenchmarkUnpackBase64Confine function
func BenchmarkUnpackBase64Confine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		err := UnpackBase64Confine(src, dest)
		if err != nil {
			b.Fatal("Error Unpack Base64:", err)
		}
	}
}

// BenchmarkUnpackBase64ToFile function
func BenchmarkUnpackBase64ToFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackBase64ToFile(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack Base64 To File:", err)
		}
	}
}

// BenchmarkUnpackBase64ToFileConfine function
func BenchmarkUnpackBase64ToFileConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		dest := "../test/data/unpack/"
		target := "file_1.txt"
		err := UnpackBase64ToFileConfine(src, target, dest)
		if err != nil {
			b.Fatal("Error Unpack Base64 To File:", err)
		}
	}
}

// BenchmarkUnpackBase64ToMemory function
func BenchmarkUnpackBase64ToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		src := "../test/data/unpack/file_base64.txt"
		target := "file_1.txt"
		err := UnpackBase64ToMemory(src, target, &dest)
		if err != nil {
			b.Fatal("Error Unpack Base64 To Memory:", err)
		}
	}
}

// BenchmarkUnpackBase64ExtractInfo function
func BenchmarkUnpackBase64ExtractInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []string
		var sz []int
		src := "../test/data/unpack/file_base64.txt"
		err := UnpackBase64ExtractInfo(src, &dest, &sz)
		if err != nil {
			b.Fatal("Error Unpack Base64 Extract Information:", err)
		}
		if len(dest) != 5 {
			b.Fatal("Error Extract file number")
		}
	}
}

// BenchmarkUnpackBase64WorkCalculate function
func BenchmarkUnpackBase64WorkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/unpack/file_base64.txt"
		_, err := UnpackBase64WorkCalculate(src)
		if err != nil {
			b.Fatal("Error Unpack Base64 Work Calculate:", err)
		}
	}
}

// BenchmarkUnpackBase64OneToMemory function
func BenchmarkUnpackBase64OneToMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest string
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		err = UnpackBase64OneToMemory(s, &dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
		if dest != "hello,world!" {
			b.Fatal("Error unpack content:", dest)
		}
	}
}

// BenchmarkUnpackBase64OneGo function
func BenchmarkUnpackBase64OneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		dest := "../test/data/unpack/"
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		wg.Add(1)
		go UnpackBase64OneGo(s, h, dest, &wg)
		wg.Wait()
	}
}

// BenchmarkUnpackBase64One function
func BenchmarkUnpackBase64One(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		dest := "../test/data/unpack/"
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		err = UnpackBase64One(s, h, dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

// BenchmarkUnpackBase64One function
func BenchmarkUnpackBase64OneConfine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		dest := "../test/data/unpack/"
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		err = UnpackBase64OneConfine(s, h, dest)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

// BenchmarkBase64DecryptGo function
func BenchmarkBase64DecryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r string
		src := "aGVsbG8sd29ybGQh"
		dest := "hello,world!"
		wg.Add(1)
		go Base64DecryptGo(src, &r, &wg)
		wg.Wait()
		if r != dest {
			b.Errorf("Error Decrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}

// BenchmarkBase64Decrypt function
func BenchmarkBase64Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "aGVsbG8sd29ybGQh"
		dest := "hello,world!"
		r := Base64Decrypt(src)
		if r != dest {
			b.Errorf("Error Decrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}
