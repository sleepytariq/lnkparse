package lnk

type LnkFile struct {
	FileName              string        `json:"FileName"`
	HeaderSize            int32         `json:"-"`
	CLSID                 string        `json:"-"`
	DataFlags             DataFlags     `json:"DataFlags"`
	FileAttrFlags         FileAttrFlags `json:"FileAttrFlags"`
	CreationTimestamp     uint64        `json:"CreationTimestamp"`
	LastAccessTimestamp   uint64        `json:"LastAccessTimestamp"`
	ModificationTimestamp uint64        `json:"ModificationTimestamp"`
	FileSize              uint32        `json:"FileSize"`
	IconIndex             int32         `json:"IconIndex"`
	ShowWindow            int32         `json:"ShowWindow"`
	HotKey                byte          `json:"HotKey"`
	HotKeyModifier        byte          `json:"HotKeyModifier"`
	IDListSize            int16         `json:"-"`
	LinkInfoSize          int32         `json:"-"`
	Name                  string        `json:"Name"`
	RelativePath          string        `json:"RelativePath"`
	WorkingDir            string        `json:"WorkingDir"`
	Arguments             string        `json:"Arguments"`
	IconLocation          string        `json:"IconLocation"`
}

type DataFlags struct {
	HasTargetIDList             bool `json:"HasTargetIDList"`
	HasLinkInfo                 bool `json:"HasLinkInfo"`
	HasName                     bool `json:"HasLinkName"`
	HasRelativePath             bool `json:"HasRelativePath"`
	HasWorkingDir               bool `json:"HasWorkingDir"`
	HasArguments                bool `json:"HasArguments"`
	HasIconLocation             bool `json:"HasIconLocation"`
	IsUnicode                   bool `json:"IsUnicode"`
	ForceNoLinkInfo             bool `json:"ForceNoLinkInfo"`
	HasExpString                bool `json:"HasExpString"`
	RunInSeparateProcess        bool `json:"RunInSeparateProcess"`
	HasDarwinID                 bool `json:"HasDarwinID"`
	RunAsUser                   bool `json:"RunAsUser"`
	HasExpIcon                  bool `json:"HasExpIcon"`
	NoPidlAlias                 bool `json:"NoPidlAlias"`
	RunWithShimLayer            bool `json:"RunWithShimLayer"`
	ForceNoLinkTrack            bool `json:"ForceNoLinkTrack"`
	EnableTargetMetadata        bool `json:"EnableTargetMetadata"`
	DisableLinkPathTracking     bool `json:"DisableLinkPathTracking"`
	DisableKnownFolderTracking  bool `json:"DisableKnownFolderTracking"`
	DisableKnownFolderAlias     bool `json:"DisableKnownFolderAlias"`
	AllowLinkToLink             bool `json:"AllowLinkToLink"`
	UnaliasOnSave               bool `json:"UnaliasOnSave"`
	PreferEnvironmentPath       bool `json:"PreferEnvironmentPath"`
	KeepLocalIDListForUNCTarget bool `json:"KeepLocalIDListForUNCTarget"`
}

type FileAttrFlags struct {
	FILE_ATTRIBUTE_READONLY            bool `json:"FILE_ATTRIBUTE_READONLY"`
	FILE_ATTRIBUTE_HIDDEN              bool `json:"FILE_ATTRIBUTE_HIDDEN"`
	FILE_ATTRIBUTE_SYSTEM              bool `json:"FILE_ATTRIBUTE_SYSTEM"`
	FILE_ATTRIBUTE_DIRECTORY           bool `json:"FILE_ATTRIBUTE_DIRECTORY"`
	FILE_ATTRIBUTE_ARCHIVE             bool `json:"FILE_ATTRIBUTE_ARCHIVE"`
	FILE_ATTRIBUTE_DEVICE              bool `json:"FILE_ATTRIBUTE_DEVICE"`
	FILE_ATTRIBUTE_NORMAL              bool `json:"FILE_ATTRIBUTE_NORMAL"`
	FILE_ATTRIBUTE_TEMPORARY           bool `json:"FILE_ATTRIBUTE_TEMPORARY"`
	FILE_ATTRIBUTE_SPARSE_FILE         bool `json:"FILE_ATTRIBUTE_SPARSE_FILE"`
	FILE_ATTRIBUTE_REPARSE_POINT       bool `json:"FILE_ATTRIBUTE_REPARSE_POINT"`
	FILE_ATTRIBUTE_COMPRESSED          bool `json:"FILE_ATTRIBUTE_COMPRESSED"`
	FILE_ATTRIBUTE_OFFLINE             bool `json:"FILE_ATTRIBUTE_OFFLINE"`
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED bool `json:"FILE_ATTRIBUTE_NOT_CONTENT_INDEXED"`
	FILE_ATTRIBUTE_ENCRYPTED           bool `json:"FILE_ATTRIBUTE_ENCRYPTED"`
	FILE_ATTRIBUTE_VIRTUAL             bool `json:"FILE_ATTRIBUTE_VIRTUAL"`
}

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
	0x00: "NONE",
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
	0x00: "NONE",
	0x01: "HOTKEYF_SHIFT",
	0x02: "HOTKEYF_CONTROL",
	0x33: "HOTKEYF_ALT",
}
