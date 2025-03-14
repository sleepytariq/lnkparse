package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
	"unicode/utf16"
)

func ReadBytes(f *os.File, n int) ([]byte, error) {
	if n < 1 {
		return nil, fmt.Errorf("size cannot be less than 1")
	}
	buf := make([]byte, n)
	_, err := f.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ReadStringNull(f *os.File) (string, error) {
	var buf []byte
	for {
		b, err := ReadBytes(f, 1)
		if err != nil {
			return "", err
		}
		if bytes.Compare(b, []byte{0x00}) == 0 {
			break
		}
		buf = append(buf, b...)
	}
	return string(buf), nil
}

func ReadString(f *os.File, n int, IsUnicode bool) (string, error) {
	if IsUnicode {
		n = n * 2
	}

	data, err := ReadBytes(f, n)
	if err != nil {
		return "", err
	}

	var s string
	if IsUnicode {
		s = DecodeUTF16(data)
	} else {
		s = string(data)
	}

	return s, nil
}

func ConvertFILETIMEToUTC(data []byte) time.Time {
	timestamp := binary.LittleEndian.Uint64(data)
	const fileTimeEpoch = 116444736000000000
	seconds := (timestamp - fileTimeEpoch) / 10000000
	nanos := (timestamp - fileTimeEpoch) % 10000000 * 100
	convertedTime := time.Unix(int64(seconds), int64(nanos)).UTC()
	if time.Since(convertedTime) < 0 {
		convertedTime = time.Unix(0, 0).UTC()
	}
	return convertedTime
}

func ConvertBytesToHumanReadableForm(size uint32) string {
	if size == 0 {
		return "0 B"
	}
	units := []string{"B", "KiB", "MiB", "GiB"}
	sizeInFloat := float64(size)
	var i int
	for sizeInFloat >= 1024 && i < len(units)-1 {
		sizeInFloat /= 1024
		i++
	}
	return fmt.Sprintf("%.1f %s", sizeInFloat, units[i])
}

func DecodeUTF16(data []byte) string {
	utf16ByteArray := make([]uint16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		utf16ByteArray[i/2] = uint16(data[i]) | uint16(data[i+1])<<8
	}
	return string(utf16.Decode(utf16ByteArray))
}

func BytesToGUID(data []byte) (string, error) {
	if len(data) != 16 {
		return "", fmt.Errorf("invalid GUID length %d", len(data))
	}

	d1 := binary.LittleEndian.Uint32(data[0:4])
	d2 := binary.LittleEndian.Uint16(data[4:6])
	d3 := binary.LittleEndian.Uint16(data[6:8])

	return fmt.Sprintf("%08x-%04x-%04x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		d1, d2, d3, data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15]), nil
}
