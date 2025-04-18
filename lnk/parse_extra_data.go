package lnk

import (
	"bytes"
	"encoding/binary"
	"lnkparse/util"
	"os"
)

func (l *Lnk) parseExtraData(f *os.File) error {
	for {
		currOffset, _ := f.Seek(0, 1)

		data, err := util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		if bytes.Equal(data, []byte{0x00, 0x00, 0x00, 0x00}) {
			break
		}
		var blockSize uint32
		binary.Decode(data, binary.LittleEndian, &blockSize)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		var blockSignature uint32
		binary.Decode(data, binary.LittleEndian, &blockSignature)

		switch blockSignature {
		case 0xa0000006:
			if err := l.parseDarwinID(f); err != nil {
				return err
			}
		case 0xa0000001:
			if err := l.parseEnvironmentVariableLocation(f); err != nil {
				return err
			}
		case 0xa0000007:
			if err := l.parseIconLocation(f); err != nil {
				return err
			}
		case 0xa000000b:
			if err := l.parseKnownFolderLocation(f); err != nil {
				return err
			}
		case 0xa0000008:
			if err := l.parseShimLayer(f, blockSize); err != nil {
				return err
			}
		case 0xa0000005:
			if err := l.parseSpecialFolderLocation(f); err != nil {
				return err
			}
		case 0xa0000003:
			if err := l.parseTracker(f); err != nil {
				return err
			}
		}

		f.Seek(currOffset+int64(blockSize), 0)
	}

	return nil
}

func (l *Lnk) parseDarwinID(f *os.File) error {
	data, err := util.ReadBytes(f, 260)
	if err != nil {
		return err
	}
	l.ExtraData.DarwinID = string(data)
	return nil
}

func (l *Lnk) parseEnvironmentVariableLocation(f *os.File) error {
	data, err := util.ReadBytes(f, 260)
	if err != nil {
		return err
	}
	l.ExtraData.ExpString = string(data)
	return nil
}

func (l *Lnk) parseIconLocation(f *os.File) error {
	data, err := util.ReadBytes(f, 260)
	if err != nil {
		return err
	}
	l.ExtraData.ExpIcon = string(bytes.TrimRight(data, "\x00"))
	return nil
}

func (l *Lnk) parseKnownFolderLocation(f *os.File) error {
	data, err := util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.KnownFolderLocation, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lnk) parseShimLayer(f *os.File, blockSize uint32) error {
	dataSize := blockSize - 8
	data, err := util.ReadBytes(f, int(dataSize))
	if err != nil {
		return err
	}
	l.ExtraData.ShimLayer = util.DecodeUTF16(data)
	return nil
}

func (l *Lnk) parseSpecialFolderLocation(f *os.File) error {
	data, err := util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.ExtraData.SpecialFolderID)
	return nil
}

func (l *Lnk) parseTracker(f *os.File) error {
	f.Seek(8, 1)

	data, err := util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.Tracker.MachineID = string(bytes.TrimRight(data, "\x00"))

	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.Tracker.DroidVolumeID, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}

	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.Tracker.DroidFileID, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}

	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.Tracker.DroidVolumeBirth, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}

	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.ExtraData.Tracker.DroidFileBirth, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}

	return nil
}
