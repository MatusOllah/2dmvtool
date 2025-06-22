package adbutil

import (
	"fmt"
	"path/filepath"

	"github.com/electricbubble/gadb"
)

func parseADBAddress(adbAddress string) (string, int, error) {
	if adbAddress == "" {
		return "localhost", 5037, nil // default
	}

	var host string
	var port int
	n, err := fmt.Sscanf(adbAddress, "%[^:]:%d", &host, &port)
	if err != nil || n != 2 {
		return "", 0, fmt.Errorf("invalid ADB address format: %s", adbAddress)
	}

	if port <= 0 || port > 65535 {
		return "", 0, fmt.Errorf("port must be between 1 and 65535: %d", port)
	}

	return host, port, nil
}

func OpenDevice(adbAddress string, serial string) (*gadb.Device, error) {
	host, port, err := parseADBAddress(adbAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ADB address %s: %w", adbAddress, err)
	}

	client, err := gadb.NewClientWith(host, port)
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
