// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package zeroheavy

import (
	"bytes"
	"errors"
	"io"
	"math/rand"
	"testing"
)

func TestZeroheavyNullInput(t *testing.T) {
	inBuf := []byte{}
	source := bytes.NewReader(inBuf)
	enc := NewZeroheavyEncoder(source)
	dec := NewZeroheavyDecoder(enc)

	var buf [256]byte
	n, err := dec.Read(buf[:])
	if !errors.Is(err, io.EOF) {
		t.Fatal()
	}
	if n != 0 {
		t.Fatal(n, buf[0])
	}

	mid, err := ZeroheavyCompress(inBuf)
	if err != nil {
		t.Error(err)
	}
	res, err := ZeroheavyDecompress(mid, 1000)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(inBuf, res) {
		t.Fatal()
	}
}

func TestZeroHeavyOneByte(t *testing.T) {
	for i := 0; i < 256; i++ {
		inBuf := []byte{byte(i)}
		source := bytes.NewReader(inBuf)
		enc := NewZeroheavyEncoder(source)
		dec := NewZeroheavyDecoder(enc)

		buf, err := io.ReadAll(dec)
		if err != nil {
			t.Error(err)
		}
		if len(buf) != 1 {
			t.Fatal(i, len(buf))
		}
		if buf[0] != byte(i) {
			t.Fatal(buf[0], i)
		}

		mid, err := ZeroheavyCompress(inBuf)
		if err != nil {
			t.Error(err)
		}
		res, err := ZeroheavyDecompress(mid, 1000)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(inBuf, res) {
			t.Fatal()
		}
	}
}

func TestZeroHeavyRandomData(t *testing.T) {
	rand.Seed(0)
	for i := 0; i < 1024; i++ {
		size := rand.Uint64() % 4096
		inBuf := make([]byte, size)
		_, _ = rand.Read(inBuf)
		dec := NewZeroheavyDecoder(NewZeroheavyEncoder(bytes.NewReader(inBuf)))
		res, err := io.ReadAll(dec)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(inBuf, res) {
			t.Fatal()
		}
		mid, err := ZeroheavyCompress(inBuf)
		if err != nil {
			t.Error(err)
		}
		res, err = ZeroheavyDecompress(mid, 5000)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(inBuf, res) {
			t.Fatal()
		}
	}
}
