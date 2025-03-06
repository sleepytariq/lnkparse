package lnk

import (
	"encoding/binary"
	"lnkparse/util"
	"os"
)

var showWindowValues = map[int]string{
	0:   "SW_HIDE",
	1:   "SW_NORMAL + SW_SHOWNORMAL",
	2:   "SW_SHOWMINIMIZED",
	3:   "SW_MAXIMIZE + SW_SHOWMAXIMIZED",
	4:   "SW_SHOWNOACTIVATE",
	5:   "SW_SHOW",
	6:   "SW_MINIMIZE",
	7:   "SW_SHOWMINNOACTIVE",
	8:   "SW_SHOWNA",
	9:   "SW_RESTORE",
	10:  "SW_SHOWDEFAULT",
	11:  "SW_FORCEMINIMIZE",
	204: "SW_NORMALNA",
}

var hotKeyValues = map[byte]string{
	0x30: "Numeric key 0",
	0x31: "Numeric key 1",
	0x32: "Numeric key 2",
	0x33: "Numeric key 3",
	0x34: "Numeric key 4",
	0x35: "Numeric key 5",
	0x36: "Numeric key 6",
	0x37: "Numeric key 7",
	0x38: "Numeric key 8",
	0x39: "Numeric key 9",
	0x41: "Upper case A",
	0x42: "Upper case B",
	0x43: "Upper case C",
	0x44: "Upper case D",
	0x45: "Upper case E",
	0x46: "Upper case F",
	0x47: "Upper case G",
	0x48: "Upper case H",
	0x49: "Upper case I",
	0x4a: "Upper case J",
	0x4b: "Upper case K",
	0x4c: "Upper case L",
	0x4d: "Upper case M",
	0x4e: "Upper case N",
	0x4f: "Upper case O",
	0x50: "Upper case P",
	0x51: "Upper case Q",
	0x52: "Upper case R",
	0x53: "Upper case S",
	0x54: "Upper case T",
	0x55: "Upper case U",
	0x56: "Upper case V",
	0x57: "Upper case W",
	0x58: "Upper case X",
	0x59: "Upper case Y",
	0x5a: "Upper case Z",
	0x70: "VK_F1",
	0x71: "VK_F2",
	0x72: "VK_F3",
	0x73: "VK_F4",
	0x74: "VK_F5",
	0x75: "VK_F6",
	0x76: "VK_F7",
	0x77: "VK_F8",
	0x78: "VK_F9",
	0x79: "VK_F10",
	0x7a: "VK_F11",
	0x7b: "VK_F12",
	0x7c: "VK_F13",
	0x7d: "VK_F14",
	0x7e: "VK_F15",
	0x7f: "VK_F16",
	0x80: "VK_F17",
	0x81: "VK_F18",
	0x82: "VK_F19",
	0x83: "VK_F20",
	0x84: "VK_F21",
	0x85: "VK_F22",
	0x86: "VK_F23",
	0x87: "VK_F24",
	0x90: "VK_NUMLOCK",
	0x91: "VK_SCROLL",
}

var hotKeyModifiers = map[byte]string{
	0x01: "HOTKEYF_SHIFT",
	0x02: "HOTKEYF_CONTROL",
	0x33: "HOTKEYF_ALT",
}

func (l *Lnk) parseHeader(f *os.File) error {
	data, err := util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.Header.size)

	data, err = util.ReadBytes(f, 16)
	if err != nil {
		return err
	}
	l.Header.class, err = util.BytesToGUID(data)
	if err != nil {
		return err
	}

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	DataFlagsInt := binary.LittleEndian.Uint32(data)

	if DataFlagsInt&0x00000001 != 0 {
		l.Header.DataFlags.HasTargetIDList = true
	}

	if DataFlagsInt&0x00000002 != 0 {
		l.Header.DataFlags.HasLinkInfo = true
	}

	if DataFlagsInt&0x00000004 != 0 {
		l.Header.DataFlags.HasName = true
	}

	if DataFlagsInt&0x00000008 != 0 {
		l.Header.DataFlags.HasRelativePath = true
	}

	if DataFlagsInt&0x00000010 != 0 {
		l.Header.DataFlags.HasWorkingDir = true
	}

	if DataFlagsInt&0x00000020 != 0 {
		l.Header.DataFlags.HasArguments = true
	}

	if DataFlagsInt&0x00000040 != 0 {
		l.Header.DataFlags.HasIconLocation = true
	}

	if DataFlagsInt&0x00000080 != 0 {
		l.Header.DataFlags.IsUnicode = true
	}

	if DataFlagsInt&0x00000100 != 0 {
		l.Header.DataFlags.ForceNoLinkInfo = true
	}

	if DataFlagsInt&0x00000200 != 0 {
		l.Header.DataFlags.HasExpString = true
	}

	if DataFlagsInt&0x00000400 != 0 {
		l.Header.DataFlags.RunInSeparateProcess = true
	}

	if DataFlagsInt&0x00001000 != 0 {
		l.Header.DataFlags.HasDarwinID = true
	}

	if DataFlagsInt&0x00002000 != 0 {
		l.Header.DataFlags.RunAsUser = true
	}

	if DataFlagsInt&0x00004000 != 0 {
		l.Header.DataFlags.HasExpIcon = true
	}

	if DataFlagsInt&0x00008000 != 0 {
		l.Header.DataFlags.NoPidlAlias = true
	}

	if DataFlagsInt&0x00020000 != 0 {
		l.Header.DataFlags.RunWithShimLayer = true
	}

	if DataFlagsInt&0x00040000 != 0 {
		l.Header.DataFlags.ForceNoLinkTrack = true
	}

	if DataFlagsInt&0x00080000 != 0 {
		l.Header.DataFlags.EnableTargetMetadata = true
	}

	if DataFlagsInt&0x00100000 != 0 {
		l.Header.DataFlags.DisableLinkPathTracking = true
	}

	if DataFlagsInt&0x00200000 != 0 {
		l.Header.DataFlags.DisableKnownFolderTracking = true
	}

	if DataFlagsInt&0x00400000 != 0 {
		l.Header.DataFlags.DisableKnownFolderAlias = true
	}

	if DataFlagsInt&0x00800000 != 0 {
		l.Header.DataFlags.AllowLinkToLink = true
	}

	if DataFlagsInt&0x01000000 != 0 {
		l.Header.DataFlags.UnaliasOnSave = true
	}

	if DataFlagsInt&0x02000000 != 0 {
		l.Header.DataFlags.PreferEnvironmentPath = true
	}

	if DataFlagsInt&0x04000000 != 0 {
		l.Header.DataFlags.PreferEnvironmentPath = true
	}

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	FileAttrFlagsInt := binary.LittleEndian.Uint32(data)

	if FileAttrFlagsInt&0x00000001 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_READONLY = true
	}

	if FileAttrFlagsInt&0x00000002 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_HIDDEN = true
	}

	if FileAttrFlagsInt&0x00000004 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_SYSTEM = true
	}

	if FileAttrFlagsInt&0x00000010 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_DIRECTORY = true
	}

	if FileAttrFlagsInt&0x00000020 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_ARCHIVE = true
	}

	if FileAttrFlagsInt&0x00000040 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_DEVICE = true
	}

	if FileAttrFlagsInt&0x00000080 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_NORMAL = true
	}

	if FileAttrFlagsInt&0x00000100 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_TEMPORARY = true
	}

	if FileAttrFlagsInt&0x00000200 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_SPARSE_FILE = true
	}

	if FileAttrFlagsInt&0x00000400 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_REPARSE_POINT = true
	}

	if FileAttrFlagsInt&0x00000800 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_COMPRESSED = true
	}

	if FileAttrFlagsInt&0x00001000 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_OFFLINE = true
	}

	if FileAttrFlagsInt&0x00002000 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_NOT_CONTENT_INDEXED = true
	}

	if FileAttrFlagsInt&0x00004000 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_ENCRYPTED = true
	}

	if FileAttrFlagsInt&0x00010000 != 0 {
		l.Header.FileAttributeFlags.FILE_ATTRIBUTE_VIRTUAL = true
	}

	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return err
	}
	l.Header.CreationTimestamp = util.ConvertFILETIMEToUTC(data)

	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return err
	}
	l.Header.LastAccessTimestamp = util.ConvertFILETIMEToUTC(data)

	data, err = util.ReadBytes(f, 8)
	if err != nil {
		return err
	}
	l.Header.ModificationTimestamp = util.ConvertFILETIMEToUTC(data)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.Header.FileSize)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.Header.IconIndex)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	var ShowWindowIntVal int32
	binary.Decode(data, binary.LittleEndian, &ShowWindowIntVal)
	l.Header.ShowWindow = showWindowValues[int(ShowWindowIntVal)]

	data, err = util.ReadBytes(f, 1)
	if err != nil {
		return err
	}
	l.Header.HotKey = hotKeyValues[data[0]]

	data, err = util.ReadBytes(f, 1)
	if err != nil {
		return err
	}
	l.Header.HotKeyModifier = hotKeyModifiers[data[0]]

	return nil
}
