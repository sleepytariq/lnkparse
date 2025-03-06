package lnk

import (
	"encoding/binary"
	"fmt"
	"lnkparse/util"
	"os"
	"path/filepath"
)

func ParseFromFile(path string) (*Lnk, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s", path)
	}
	defer f.Close()

	l := Lnk{}

	l.FileName = filepath.Base(path)

	if err := l.parseHeader(f); err != nil {
		return nil, fmt.Errorf("failed to parse Header for %s", path)
	}
	if l.Header.size != 76 {
		return nil, fmt.Errorf("invalid file header for %s", path)
	}

	// skip remaining 10 reserved bytes
	f.Seek(10, 1)

	if l.Header.DataFlags.HasTargetIDList {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IDListSize for %s", path)
		}
		var idListSize uint16
		binary.Decode(data, binary.LittleEndian, &idListSize)
		f.Seek(int64(idListSize), 1)
	}

	if l.Header.DataFlags.HasLinkInfo {
		err := l.parseLinkInfo(f)
		if err != nil {
			return nil, fmt.Errorf("failed to parse LinkInfo for %s", path)
		}
	}

	if err = l.parseDataStrings(f); err != nil {
		return nil, fmt.Errorf("failed to parse DataStrings for %s", path)
	}

	if err = l.parseExtraData(f); err != nil {
		return nil, fmt.Errorf("failed to parse ExtraData for %s", path)
	}

	return &l, nil
}
