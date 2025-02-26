package lnk

import (
	"fmt"
	"lnkparse/util"
	"reflect"
	"strings"
)

func (l *LnkFile) String() string {
	var s []string

	s = append(s, fmt.Sprintf("%-24s: %s", "File Name", l.FileName))

	dataVal := reflect.ValueOf(l.DataFlags)
	var EnabledDataFlags []string
	for i := range dataVal.NumField() {
		flag := dataVal.Type().Field(i).Name
		enabled := dataVal.Field(i).Bool()
		if enabled {
			EnabledDataFlags = append(EnabledDataFlags, flag)
		}
	}

	if len(EnabledDataFlags) == 0 {
		EnabledDataFlags = append(EnabledDataFlags, "NONE")
	}

	s = append(s, fmt.Sprintf("%-24s: %s", "Data Flags", strings.Join(EnabledDataFlags, ", ")))

	fileVal := reflect.ValueOf(l.FileAttrFlags)
	var EnabledFileAttrFlags []string
	for i := range fileVal.NumField() {
		flag := fileVal.Type().Field(i).Name
		enabled := fileVal.Field(i).Bool()
		if enabled {
			EnabledFileAttrFlags = append(EnabledFileAttrFlags, flag)
		}
	}

	if len(EnabledFileAttrFlags) == 0 {
		EnabledFileAttrFlags = append(EnabledFileAttrFlags, "NONE")
	}

	s = append(s, fmt.Sprintf("%-24s: %s", "File Attributes Flags", strings.Join(EnabledFileAttrFlags, ", ")))

	s = append(s, fmt.Sprintf("%-24s: %s", "Creation Timestamp", util.ConvertFILETIMEToUTCString(l.CreationTimestamp)))
	s = append(s, fmt.Sprintf("%-24s: %s", "Last Access Timestamp", util.ConvertFILETIMEToUTCString(l.LastAccessTimestamp)))
	s = append(s, fmt.Sprintf("%-24s: %s", "Modification Timestamp", util.ConvertFILETIMEToUTCString(l.ModificationTimestamp)))
	s = append(s, fmt.Sprintf("%-24s: %s", "File Size", util.ConvertBytesToHumanReadableForm(l.FileSize)))
	s = append(s, fmt.Sprintf("%-24s: %d", "Icon Index", l.IconIndex))
	s = append(s, fmt.Sprintf("%-24s: %s", "Show Window", showWindowValues[int(l.ShowWindow)]))
	s = append(s, fmt.Sprintf("%-24s: %s", "Hot Key", hotKeyValues[l.HotKey]))
	s = append(s, fmt.Sprintf("%-24s: %s", "Hot Key Modifier", hotKeyModifiers[l.HotKeyModifier]))

	if l.DataFlags.HasRelativePath {
		s = append(s, fmt.Sprintf("%-24s: %s", "Relative Path", l.RelativePath))
	}

	if l.DataFlags.HasWorkingDir {
		s = append(s, fmt.Sprintf("%-24s: %s", "Working Directory", l.WorkingDir))
	}

	if l.DataFlags.HasArguments {
		s = append(s, fmt.Sprintf("%-24s: %s", "Command Line", l.Arguments))
	}

	if l.DataFlags.HasIconLocation {
		s = append(s, fmt.Sprintf("%-24s: %s", "Icon Location", l.IconLocation))
	}

	return strings.Join(s, "\n")
}
