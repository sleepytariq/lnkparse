package lnk

import (
	"encoding/binary"
	"lnkparse/util"
	"os"
)

var driveTypeValues = map[int]string{
	0: "DRIVE_UNKNOWN",
	1: "DRIVE_NO_ROOT_DIR",
	2: "DRIVE_REMOVABLE",
	3: "DRIVE_FIXED",
	4: "DRIVE_REMOTE",
	5: "DRIVE_CDROM",
	6: "DRIVE_RAMDISK",
}

var providerTypeValues = map[uint32]string{
	0x001a0000: "WNNC_NET_AVID",
	0x001b0000: "WNNC_NET_DOCUSPACE",
	0x001c0000: "WNNC_NET_MANGOSOFT",
	0x001d0000: "WNNC_NET_SERNET",
	0x001e0000: "WNNC_NET_RIVERFRONT1",
	0x001f0000: "WNNC_NET_RIVERFRONT2",
	0x00200000: "WNNC_NET_DECORB",
	0x00210000: "WNNC_NET_PROTSTOR",
	0x00220000: "WNNC_NET_FJ_REDIR",
	0x00230000: "WNNC_NET_DISTINCT",
	0x00240000: "WNNC_NET_TWINS",
	0x00250000: "WNNC_NET_RDR2SAMPLE",
	0x00260000: "WNNC_NET_CSC",
	0x00270000: "WNNC_NET_3IN1",
	0x00290000: "WNNC_NET_EXTENDNET",
	0x002a0000: "WNNC_NET_STAC",
	0x002b0000: "WNNC_NET_FOXBAT",
	0x002c0000: "WNNC_NET_YAHOO",
	0x002d0000: "WNNC_NET_EXIFS",
	0x002e0000: "WNNC_NET_DAV",
	0x002f0000: "WNNC_NET_KNOWARE",
	0x00300000: "WNNC_NET_OBJECT_DIRE",
	0x00310000: "WNNC_NET_MASFAX",
	0x00320000: "WNNC_NET_HOB_NFS",
	0x00330000: "WNNC_NET_SHIVA",
	0x00340000: "WNNC_NET_IBMAL",
	0x00350000: "WNNC_NET_LOCK",
	0x00360000: "WNNC_NET_TERMSRV",
	0x00370000: "WNNC_NET_SRT",
	0x00380000: "WNNC_NET_QUINCY",
	0x00390000: "WNNC_NET_OPENAFS",
	0x003a0000: "WNNC_NET_AVID1",
	0x003b0000: "WNNC_NET_DFS",
	0x003c0000: "WNNC_NET_KWNP",
	0x003d0000: "WNNC_NET_ZENWORKS",
	0x003e0000: "WNNC_NET_DRIVEONWEB",
	0x003f0000: "WNNC_NET_VMWARE",
	0x00400000: "WNNC_NET_RSFX",
	0x00410000: "WNNC_NET_MFILES",
	0x00420000: "WNNC_NET_MS_NFS",
	0x00430000: "WNNC_NET_GOOGLE",
}

func (l *Lnk) parseLinkInfo(f *os.File) error {
	baseOffset, _ := f.Seek(0, 1)

	data, err := util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.size)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.headerSize)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	var flagsInt int32
	binary.Decode(data, binary.LittleEndian, &flagsInt)
	if flagsInt&0x0001 != 0 {
		l.LinkInfo.Flags.VolumeIDAndLocalBasePath = true
	}
	if flagsInt&0x0002 != 0 {
		l.LinkInfo.Flags.CommonNetworkRelativeLinkAndPathSuffix = true
	}

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.volumeInfoOffset)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.localPathOffset)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.networkShareInfoOffset)

	data, err = util.ReadBytes(f, 4)
	if err != nil {
		return err
	}
	binary.Decode(data, binary.LittleEndian, &l.LinkInfo.commonPathSuffixOffset)

	if l.LinkInfo.Flags.VolumeIDAndLocalBasePath {
		localBaseOffset, _ := f.Seek(baseOffset+int64(l.LinkInfo.volumeInfoOffset), 0)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.VolumeInfo.size)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		var driveTypeValue int32
		binary.Decode(data, binary.LittleEndian, &driveTypeValue)
		l.LinkInfo.VolumeInfo.DriveType = driveTypeValues[int(driveTypeValue)]

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.VolumeInfo.DriveSerialNumber)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.VolumeInfo.labelOffset)

		f.Seek(localBaseOffset+int64(l.LinkInfo.VolumeInfo.labelOffset), 0)
		l.LinkInfo.VolumeInfo.DriveLabel, err = util.ReadStringNull(f)
		if err != nil {
			return nil
		}

		f.Seek(baseOffset+int64(l.LinkInfo.localPathOffset), 0)
		l.LinkInfo.LocalBasePath, err = util.ReadStringNull(f)
		if err != nil {
			return nil
		}
	}

	if l.LinkInfo.Flags.CommonNetworkRelativeLinkAndPathSuffix {
		localBaseOffset, _ := f.Seek(baseOffset+int64(l.LinkInfo.networkShareInfoOffset), 0)

		data, err := util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.NetworkShareInfo.size)

		var flagsInt int32
		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &flagsInt)
		if flagsInt&0x0001 != 0 {
			l.LinkInfo.NetworkShareInfo.Flags.ValidDevice = true
		}
		if flagsInt&0x0002 != 0 {
			l.LinkInfo.NetworkShareInfo.Flags.ValidNetType = true
		}

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.NetworkShareInfo.networkShareNameOffset)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		binary.Decode(data, binary.LittleEndian, &l.LinkInfo.NetworkShareInfo.deviceNameOffset)

		data, err = util.ReadBytes(f, 4)
		if err != nil {
			return err
		}
		var ProviderTypeInt uint32
		binary.Decode(data, binary.LittleEndian, &ProviderTypeInt)
		l.LinkInfo.NetworkShareInfo.ProviderType = providerTypeValues[ProviderTypeInt]

		f.Seek(localBaseOffset+int64(l.LinkInfo.NetworkShareInfo.networkShareNameOffset), 0)
		l.LinkInfo.NetworkShareInfo.NetworkShareName, err = util.ReadStringNull(f)
		if err != nil {
			return nil
		}

		f.Seek(localBaseOffset+int64(l.LinkInfo.NetworkShareInfo.deviceNameOffset), 0)
		l.LinkInfo.NetworkShareInfo.DeviceName, err = util.ReadStringNull(f)
		if err != nil {
			return nil
		}
	}

	f.Seek(baseOffset+int64(l.LinkInfo.size), 0)

	return nil
}
