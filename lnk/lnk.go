package lnk

import (
	"fmt"
	"lnkparse/util"
	"reflect"
	"strings"
	"time"
)

type Lnk struct {
	FileName string `json:"File Name"`
	Header   struct {
		size      uint32
		class     string
		DataFlags struct {
			HasTargetIDList             bool
			HasLinkInfo                 bool
			HasName                     bool
			HasRelativePath             bool
			HasWorkingDir               bool
			HasArguments                bool
			HasIconLocation             bool
			IsUnicode                   bool
			ForceNoLinkInfo             bool
			HasExpString                bool
			RunInSeparateProcess        bool
			HasDarwinID                 bool
			RunAsUser                   bool
			HasExpIcon                  bool
			NoPidlAlias                 bool
			RunWithShimLayer            bool
			ForceNoLinkTrack            bool
			EnableTargetMetadata        bool
			DisableLinkPathTracking     bool
			DisableKnownFolderTracking  bool
			DisableKnownFolderAlias     bool
			AllowLinkToLink             bool
			UnaliasOnSave               bool
			PreferEnvironmentPath       bool
			KeepLocalIDListForUNCTarget bool
		} `json:"Data Flags"`
		FileAttributeFlags struct {
			FILE_ATTRIBUTE_READONLY            bool
			FILE_ATTRIBUTE_HIDDEN              bool
			FILE_ATTRIBUTE_SYSTEM              bool
			FILE_ATTRIBUTE_DIRECTORY           bool
			FILE_ATTRIBUTE_ARCHIVE             bool
			FILE_ATTRIBUTE_DEVICE              bool
			FILE_ATTRIBUTE_NORMAL              bool
			FILE_ATTRIBUTE_TEMPORARY           bool
			FILE_ATTRIBUTE_SPARSE_FILE         bool
			FILE_ATTRIBUTE_REPARSE_POINT       bool
			FILE_ATTRIBUTE_COMPRESSED          bool
			FILE_ATTRIBUTE_OFFLINE             bool
			FILE_ATTRIBUTE_NOT_CONTENT_INDEXED bool
			FILE_ATTRIBUTE_ENCRYPTED           bool
			FILE_ATTRIBUTE_VIRTUAL             bool
		} `json:"File Attribute Flags"`
		CreationTimestamp     time.Time `json:"Creation Timestamp"`
		LastAccessTimestamp   time.Time `json:"Last Access Timestamp"`
		ModificationTimestamp time.Time `json:"Modification Timestamp"`
		FileSize              uint32    `json:"File Size"`
		IconIndex             int32     `json:"Icon Index"`
		ShowWindow            string    `json:"Show Window"`
		HotKey                string    `json:"Hot Key"`
		HotKeyModifier        string    `json:"Hot Key Modifier"`
	}
	LinkInfo struct {
		size       uint32
		headerSize uint32
		Flags      struct {
			VolumeIDAndLocalBasePath               bool
			CommonNetworkRelativeLinkAndPathSuffix bool
		}
		volumeInfoOffset       uint32
		localPathOffset        uint32
		networkShareInfoOffset uint32
		commonPathSuffixOffset uint32
		VolumeInfo             struct {
			size              uint32
			DriveType         string `json:"Drive Type"`
			DriveSerialNumber uint32 `json:"Drive Serial Number"`
			labelOffset       uint32
			DriveLabel        string `json:"Drive Label"`
		} `json:"Volume Info"`
		LocalBasePath    string `json:"Local Base Path"`
		NetworkShareInfo struct {
			size  uint32
			Flags struct {
				ValidDevice  bool
				ValidNetType bool
			}
			networkShareNameOffset uint32
			deviceNameOffset       uint32
			ProviderType           string `json:"Provider Type"`
			NetworkShareName       string `json:"Network Share Name"`
			DeviceName             string `json:"Device Name"`
		} `json:"Network Share Info"`
		CommonPathSuffix string `json:"Common Path Suffix"`
	} `json:"Link Info"`
	DataStrings struct {
		Name         string
		RelativePath string `json:"Relative Path"`
		WorkingDir   string `json:"Working Directory"`
		Arguments    string
		IconLocation string `json:"Icon Location"`
	} `json:"Data Strings"`
	ExtraData struct {
		DarwinID            string `json:"Darwin ID"`
		ExpString           string `json:"Exp String"`
		ExpIcon             string `json:"Exp Icon"`
		KnownFolderLocation string `json:"Known Folder Location"`
		ShimLayer           string `json:"Shim Layer"`
		SpecialFolderID     uint32 `json:"Special Folder ID"`
		Tracker             struct {
			MachineID        string `json:"Machine ID"`
			DroidVolumeID    string `json:"Droid Volume ID"`
			DroidFileID      string `json:"Droid File ID"`
			DroidVolumeBirth string `json:"Droid Volume Birth"`
			DroidFileBirth   string `json:"Droid File Birth"`
		}
	} `json:"Extra Data"`
}

func (l *Lnk) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%-24s: %s", "File Name", l.FileName))

	dataVal := reflect.ValueOf(l.Header.DataFlags)
	var EnabledDataFlags []string
	for i := range dataVal.NumField() {
		flag := dataVal.Type().Field(i).Name
		enabled := dataVal.Field(i).Bool()
		if enabled {
			EnabledDataFlags = append(EnabledDataFlags, flag)
		}
	}
	if len(EnabledDataFlags) != 0 {
		s = append(s, fmt.Sprintf("%-24s: %s", "Data Flags", strings.Join(EnabledDataFlags, ", ")))
	}

	fileVal := reflect.ValueOf(l.Header.FileAttributeFlags)
	var EnabledFileAttrFlags []string
	for i := range fileVal.NumField() {
		flag := fileVal.Type().Field(i).Name
		enabled := fileVal.Field(i).Bool()
		if enabled {
			EnabledFileAttrFlags = append(EnabledFileAttrFlags, flag)
		}
	}
	if len(EnabledFileAttrFlags) != 0 {
		s = append(s, fmt.Sprintf("%-24s: %s", "File Attribute Flags", strings.Join(EnabledFileAttrFlags, ", ")))
	}

	s = append(s, fmt.Sprintf("%-24s: %s", "Creation Timestamp", l.Header.CreationTimestamp))
	s = append(s, fmt.Sprintf("%-24s: %s", "Last Access Timestamp", l.Header.LastAccessTimestamp))
	s = append(s, fmt.Sprintf("%-24s: %s", "Modification Timestamp", l.Header.ModificationTimestamp))
	s = append(s, fmt.Sprintf("%-24s: %s", "File Size", util.ConvertBytesToHumanReadableForm(l.Header.FileSize)))
	s = append(s, fmt.Sprintf("%-24s: %d", "Icon Index", l.Header.IconIndex))
	s = append(s, fmt.Sprintf("%-24s: %s", "Show Window", l.Header.ShowWindow))

	if l.Header.HotKey != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Hot Key", l.Header.HotKey))
	}

	if l.Header.HotKeyModifier != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Hot Key Modifier", l.Header.HotKeyModifier))
	}

	if l.LinkInfo.Flags.VolumeIDAndLocalBasePath {
		s = append(s, fmt.Sprintf("%-24s: %s", "Drive Type", l.LinkInfo.VolumeInfo.DriveType))
		s = append(s, fmt.Sprintf("%-24s: %d", "Drive Serial Number", l.LinkInfo.VolumeInfo.DriveSerialNumber))
		if l.LinkInfo.VolumeInfo.DriveLabel != "" {
			s = append(s, fmt.Sprintf("%-24s: %s", "Drive Label", l.LinkInfo.VolumeInfo.DriveLabel))
		}
		s = append(s, fmt.Sprintf("%-24s: %s", "Local Base Path", l.LinkInfo.LocalBasePath))
	}

	if l.LinkInfo.Flags.CommonNetworkRelativeLinkAndPathSuffix {
		s = append(s, fmt.Sprintf("%-24s: %s", "Provider Type", l.LinkInfo.NetworkShareInfo.ProviderType))
		s = append(s, fmt.Sprintf("%-24s: %s", "Network Share Name", l.LinkInfo.NetworkShareInfo.NetworkShareName))
		s = append(s, fmt.Sprintf("%-24s: %s", "Device Name", l.LinkInfo.NetworkShareInfo.DeviceName))
	}

	if l.Header.DataFlags.HasName {
		s = append(s, fmt.Sprintf("%-24s: %s", "Name", l.DataStrings.Name))
	}

	if l.Header.DataFlags.HasRelativePath {
		s = append(s, fmt.Sprintf("%-24s: %s", "Relative Path", l.DataStrings.RelativePath))
	}

	if l.Header.DataFlags.HasWorkingDir {
		s = append(s, fmt.Sprintf("%-24s: %s", "Working Directory", l.DataStrings.WorkingDir))
	}

	if l.Header.DataFlags.HasArguments {
		s = append(s, fmt.Sprintf("%-24s: %s", "Arguments", l.DataStrings.Arguments))
	}

	if l.Header.DataFlags.HasIconLocation {
		s = append(s, fmt.Sprintf("%-24s: %s", "Icon Location", l.DataStrings.IconLocation))
	}

	if l.Header.DataFlags.HasDarwinID {
		s = append(s, fmt.Sprintf("%-24s: %s", "Darwin ID", l.ExtraData.DarwinID))
	}

	if l.Header.DataFlags.HasExpString {
		s = append(s, fmt.Sprintf("%-24s: %s", "Exp String", l.ExtraData.ExpString))
	}

	if l.Header.DataFlags.HasExpIcon {
		s = append(s, fmt.Sprintf("%-24s: %s", "Exp Icon", l.ExtraData.ExpIcon))
	}

	if l.ExtraData.KnownFolderLocation != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Known Folder Location", l.ExtraData.KnownFolderLocation))
	}

	if l.Header.DataFlags.RunWithShimLayer {
		s = append(s, fmt.Sprintf("%-24s: %s", "Shim Layer", l.ExtraData.ShimLayer))
	}

	if l.ExtraData.SpecialFolderID != 0 {
		s = append(s, fmt.Sprintf("%-24s: %d", "Special Folder ID", l.ExtraData.SpecialFolderID))
	}

	if l.ExtraData.Tracker.MachineID != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Machine ID", l.ExtraData.Tracker.MachineID))
	}

	if l.ExtraData.Tracker.DroidVolumeID != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Droid Volume ID", l.ExtraData.Tracker.DroidVolumeID))
	}

	if l.ExtraData.Tracker.DroidFileID != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Droid File ID", l.ExtraData.Tracker.DroidFileID))
	}

	if l.ExtraData.Tracker.DroidVolumeBirth != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Droid Volume Birth", l.ExtraData.Tracker.DroidVolumeBirth))
	}

	if l.ExtraData.Tracker.DroidFileBirth != "" {
		s = append(s, fmt.Sprintf("%-24s: %s", "Droid File Birth", l.ExtraData.Tracker.DroidFileBirth))
	}

	return strings.Join(s, "\n")
}
