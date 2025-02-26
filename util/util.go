package util

import (
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
	err := binary.Read(f, binary.LittleEndian, &buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ConvertFILETIMEToUTCString(timestamp uint64) string {
	const fileTimeEpoch = 116444736000000000
	seconds := (timestamp - fileTimeEpoch) / 10000000
	nanos := (timestamp - fileTimeEpoch) % 10000000 * 100
	convertedTime := time.Unix(int64(seconds), int64(nanos)).UTC()
	if time.Since(convertedTime) < 0 {
		convertedTime = time.Unix(0, 0).UTC()
	}
	return convertedTime.Format("2006-01-02 15:04:05 MST")
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
	if len(data)%2 != 0 {
		data = append(data, 0x00)
	}
	utf16ByteArray := make([]uint16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		utf16ByteArray[i/2] = uint16(data[i]) | uint16(data[i+1])<<8
	}
	return string(utf16.Decode(utf16ByteArray))
}
