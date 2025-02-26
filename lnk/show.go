package lnk

import (
	"fmt"
	"lnkparse/util"
	"reflect"
	"strings"
)

func (l *LnkFile) ShowInfo() {
	fmt.Printf("%-24s: %s\n", "File Name", l.FileName)

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

	fmt.Printf("%-24s: %s\n", "Data Flags", strings.Join(EnabledDataFlags, ", "))

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

	fmt.Printf("%-24s: %s\n", "File Attributes Flags", strings.Join(EnabledFileAttrFlags, ", "))

	fmt.Printf("%-24s: %s\n", "Creation Timestamp", util.ConvertFILETIMEToUTCString(l.CreationTimestamp))

	fmt.Printf("%-24s: %s\n", "Last Access Timestamp", util.ConvertFILETIMEToUTCString(l.LastAccessTimestamp))

	fmt.Printf("%-24s: %s\n", "Modification Timestamp", util.ConvertFILETIMEToUTCString(l.ModificationTimestamp))

	fmt.Printf("%-24s: %s\n", "File Size", util.ConvertBytesToHumanReadableForm(l.FileSize))

	fmt.Printf("%-24s: %d\n", "Icon Index", l.IconIndex)

	fmt.Printf("%-24s: %s\n", "Show Window", showWindowValues[int(l.ShowWindow)])

	fmt.Printf("%-24s: %s\n", "Hot Key", hotKeyValues[l.HotKey])

	fmt.Printf("%-24s: %s\n", "Hot Key Modifier", hotKeyModifiers[l.HotKeyModifier])

	if l.DataFlags.HasRelativePath {
		fmt.Printf("%-24s: %s\n", "Relative Path", l.RelativePath)
	}

	if l.DataFlags.HasWorkingDir {
		fmt.Printf("%-24s: %s\n", "Working Directory", l.WorkingDir)
	}

	if l.DataFlags.HasArguments {
		fmt.Printf("%-24s: %s\n", "Command Line", l.Arguments)
	}

	if l.DataFlags.HasIconLocation {
		fmt.Printf("%-24s: %s\n", "Icon Location", l.IconLocation)
	}
}
