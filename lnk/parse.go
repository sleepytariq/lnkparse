package lnk

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"lnkparse/util"
	"os"
	"path/filepath"
	"strings"
)

func ParseFromFile(path string, trimSpaces bool) (*LnkFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s", path)
	}
	defer f.Close()

	l := LnkFile{}

	// set file name
	l.FileName = filepath.Base(path)

	// parse header size
	data, err := util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse header size for %s", path)
	}
	binary.Decode(data, binary.LittleEndian, &l.HeaderSize)
	if l.HeaderSize != 76 {
		return nil, fmt.Errorf("invalid file header for %s", path)
	}

	// parse the clsid
	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CLSID for %s", path)
	}
	l.CLSID = hex.EncodeToString(data)

	// parse data flags
	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse data flags for %s", path)
	}
	l.ParseDataFlags(data)

	// parse file attribute flags
	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file attribute flags for %s", path)
	}
	l.ParseFileAttrFlags(data)

	// parse creation timestamp
	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse creation timestamp for %s", path)
	}
	l.CreationTimestamp = util.ConvertFILETIMEToUTCString(data)

	// parse last access timestamp
	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse last access timestamp for %s", path)
	}
	l.LastAccessTimestamp = util.ConvertFILETIMEToUTCString(data)

	// parse modification timestamp
	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse modification timestamp for %s", path)
	}
	l.ModificationTimestamp = util.ConvertFILETIMEToUTCString(data)

	// parse file size
	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file size for %s", path)
	}
	binary.Decode(data, binary.LittleEndian, &l.FileSize)

	// parse icon index
	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse icon index for %s", path)
	}
	binary.Decode(data, binary.LittleEndian, &l.IconIndex)

	// parse show window
	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return nil, fmt.Errorf("failed to parse show window for %s", path)
	}
	var ShowWindowIntVal int32
	binary.Decode(data, binary.LittleEndian, &ShowWindowIntVal)
	l.ShowWindow = showWindowValues[int(ShowWindowIntVal)]

	// parse hot key
	data, err = util.ReadBytes(f, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to parse hot key for %s", path)
	}
	if l.HotKey = hotKeyValues[data[0]]; l.HotKey == "" {
		l.HotKey = "INVALID"
	}

	// parse hot key modifier
	data, err = util.ReadBytes(f, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to parse hot key modifier for %s", path)
	}
	if l.HotKeyModifier = hotKeyModifiers[data[0]]; l.HotKeyModifier == "" {
		l.HotKeyModifier = "INVALID"
	}

	// skip remaining 10 reserved bytes
	f.Seek(10, 1)

	if l.DataFlags.HasTargetIDList {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IDListSize for %s", path)
		}
		binary.Decode(data, binary.LittleEndian, &l.IDListSize)
		f.Seek(int64(l.IDListSize), 1)
	}

	if l.DataFlags.HasLinkInfo {
		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return nil, fmt.Errorf("failed to parse LinkInfoSize for %s", path)
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfoSize)
		f.Seek(int64(l.LinkInfoSize)-4, 1)
	}

	if l.DataFlags.HasName {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Name size for %s", path)
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.Name, err = util.ReadString(f, int(size), l.DataFlags.IsUnicode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Name for %s", path)
		}
	}

	if l.DataFlags.HasRelativePath {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RelativePath size for %s", path)
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.RelativePath, err = util.ReadString(f, int(size), l.DataFlags.IsUnicode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RelativePath for %s", path)
		}
	}

	if l.DataFlags.HasWorkingDir {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse WorkingDir size for %s", path)
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.WorkingDir, err = util.ReadString(f, int(size), l.DataFlags.IsUnicode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse WorkingDir for %s", path)
		}
	}

	if l.DataFlags.HasArguments {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Arguments size for %s", path)
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.Arguments, err = util.ReadString(f, int(size), l.DataFlags.IsUnicode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Arguments for %s", path)
		}
		if trimSpaces {
			l.Arguments = strings.Trim(l.Arguments, " ")
		}
	}

	if l.DataFlags.HasIconLocation {
		data, err = util.ReadBytes(f, 2)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IconLocation size for %s <%s>", path, err)
		}
		var size int16
		binary.Decode(data, binary.LittleEndian, &size)
		l.IconLocation, err = util.ReadString(f, int(size), l.DataFlags.IsUnicode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IconLocation for %s", path)
		}
	}

	return &l, nil
}

func (l *LnkFile) ParseDataFlags(data []byte) {
	DataFlagsInt := binary.LittleEndian.Uint32(data)

	if DataFlagsInt&0x00000001 != 0 {
		l.DataFlags.HasTargetIDList = true
	}

	if DataFlagsInt&0x00000002 != 0 {
		l.DataFlags.HasLinkInfo = true
	}

	if DataFlagsInt&0x00000004 != 0 {
		l.DataFlags.HasName = true
	}

	if DataFlagsInt&0x00000008 != 0 {
		l.DataFlags.HasRelativePath = true
	}

	if DataFlagsInt&0x00000010 != 0 {
		l.DataFlags.HasWorkingDir = true
	}

	if DataFlagsInt&0x00000020 != 0 {
		l.DataFlags.HasArguments = true
	}

	if DataFlagsInt&0x00000040 != 0 {
		l.DataFlags.HasIconLocation = true
	}

	if DataFlagsInt&0x00000080 != 0 {
		l.DataFlags.IsUnicode = true
	}

	if DataFlagsInt&0x00000100 != 0 {
		l.DataFlags.ForceNoLinkInfo = true
	}

	if DataFlagsInt&0x00000200 != 0 {
		l.DataFlags.HasExpString = true
	}

	if DataFlagsInt&0x00000400 != 0 {
		l.DataFlags.RunInSeparateProcess = true
	}

	if DataFlagsInt&0x00001000 != 0 {
		l.DataFlags.HasDarwinID = true
	}

	if DataFlagsInt&0x00002000 != 0 {
		l.DataFlags.RunAsUser = true
	}

	if DataFlagsInt&0x00004000 != 0 {
		l.DataFlags.HasExpIcon = true
	}

	if DataFlagsInt&0x00008000 != 0 {
		l.DataFlags.NoPidlAlias = true
	}

	if DataFlagsInt&0x00020000 != 0 {
		l.DataFlags.RunWithShimLayer = true
	}

	if DataFlagsInt&0x00040000 != 0 {
		l.DataFlags.ForceNoLinkTrack = true
	}

	if DataFlagsInt&0x00080000 != 0 {
		l.DataFlags.EnableTargetMetadata = true
	}

	if DataFlagsInt&0x00100000 != 0 {
		l.DataFlags.DisableLinkPathTracking = true
	}

	if DataFlagsInt&0x00200000 != 0 {
		l.DataFlags.DisableKnownFolderTracking = true
	}

	if DataFlagsInt&0x00400000 != 0 {
		l.DataFlags.DisableKnownFolderAlias = true
	}

	if DataFlagsInt&0x00800000 != 0 {
		l.DataFlags.AllowLinkToLink = true
	}

	if DataFlagsInt&0x01000000 != 0 {
		l.DataFlags.UnaliasOnSave = true
	}

	if DataFlagsInt&0x02000000 != 0 {
		l.DataFlags.PreferEnvironmentPath = true
	}

	if DataFlagsInt&0x04000000 != 0 {
		l.DataFlags.PreferEnvironmentPath = true
	}
}

func (l *LnkFile) ParseFileAttrFlags(data []byte) {
	FileAttrFlagsInt := binary.LittleEndian.Uint32(data)

	if FileAttrFlagsInt&0x00000001 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_READONLY = true
	}

	if FileAttrFlagsInt&0x00000002 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_HIDDEN = true
	}

	if FileAttrFlagsInt&0x00000004 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_SYSTEM = true
	}

	if FileAttrFlagsInt&0x00000010 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_DIRECTORY = true
	}

	if FileAttrFlagsInt&0x00000020 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_ARCHIVE = true
	}

	if FileAttrFlagsInt&0x00000040 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_DEVICE = true
	}

	if FileAttrFlagsInt&0x00000080 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_NORMAL = true
	}

	if FileAttrFlagsInt&0x00000100 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_TEMPORARY = true
	}

	if FileAttrFlagsInt&0x00000200 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_SPARSE_FILE = true
	}

	if FileAttrFlagsInt&0x00000400 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_REPARSE_POINT = true
	}

	if FileAttrFlagsInt&0x00000800 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_COMPRESSED = true
	}

	if FileAttrFlagsInt&0x00001000 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_OFFLINE = true
	}

	if FileAttrFlagsInt&0x00002000 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_NOT_CONTENT_INDEXED = true
	}

	if FileAttrFlagsInt&0x00004000 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_ENCRYPTED = true
	}

	if FileAttrFlagsInt&0x00010000 != 0 {
		l.FileAttrFlags.FILE_ATTRIBUTE_VIRTUAL = true
	}
}
