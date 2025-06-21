package adbutil

import (
	"fmt"
	"path/filepath"

	"github.com/electricbubble/gadb"
)

func OpenDevice(serial string) (*gadb.Device, error) {
	client, err := gadb.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ADB server: %w", err)
	}

	devices, err := client.DeviceList()
	if err != nil {
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}

	var device *gadb.Device
	if len(devices) == 0 {
		return nil, fmt.Errorf("no devices found")
	}
	// If no serial is provided, return the first device
	if serial == "" {
		device = &devices[0]
	} else {
		for _, d := range devices {
			if d.Serial() == serial {
				device = &d
				break
			}
		}
		if device == nil {
			return nil, fmt.Errorf("device with serial %s not found", serial)
		}
	}

	return device, nil
}

func GetRemoteFileSize(device *gadb.Device, path string) (int64, error) {
	infos, err := device.List(filepath.Dir(path))
	if err != nil {
		return 0, fmt.Errorf("failed to list directory %s: %w", filepath.Dir(path), err)
	}

	for _, info := range infos {
		if info.Name == filepath.Base(path) {
			return int64(info.Size), nil
		}
	}

	return 0, fmt.Errorf("file %s not found on device", path)
}
