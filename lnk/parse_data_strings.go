package lnk

import (
	"encoding/binary"
	"lnkparse/util"
	"os"
)

func (l *Lnk) parseDataStrings(f *os.File) error {
	if l.Header.DataFlags.HasName {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return err
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.DataStrings.Name, err = util.ReadString(f, int(size), l.Header.DataFlags.IsUnicode)
		if err != nil {
			return err
		}
	}

	if l.Header.DataFlags.HasRelativePath {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return err
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.DataStrings.RelativePath, err = util.ReadString(f, int(size), l.Header.DataFlags.IsUnicode)
		if err != nil {
			return err
		}
	}

	if l.Header.DataFlags.HasWorkingDir {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return err
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.DataStrings.WorkingDir, err = util.ReadString(f, int(size), l.Header.DataFlags.IsUnicode)
		if err != nil {
			return err
		}
	}

	if l.Header.DataFlags.HasArguments {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return err
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.DataStrings.Arguments, err = util.ReadString(f, int(size), l.Header.DataFlags.IsUnicode)
		if err != nil {
			return err
		}
	}

	if l.Header.DataFlags.HasIconLocation {
		data, err := util.ReadBytes(f, 2)
		if err != nil {
			return err
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.DataStrings.IconLocation, err = util.ReadString(f, int(size), l.Header.DataFlags.IsUnicode)
		if err != nil {
			return err
		}
	}

	return nil
}
