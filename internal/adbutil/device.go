package adbutil

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/electricbubble/gadb"
	"github.com/fatih/color"
)

func parseADBAddress(adbAddress string) (string, int, error) {
	if adbAddress == "" {
		return "localhost", 5037, nil // default
	}

	parts := strings.SplitN(adbAddress, ":", 2)
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid ADB address format, expected 'host:port': %s", adbAddress)
	}

	host := parts[0]

	portStr := parts[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port number %s: %w", portStr, err)
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
	output, err := device.RunShellCommand(fmt.Sprintf("stat -c %%s %s", filepath.ToSlash(path)))
	if err != nil {
		return 0, fmt.Errorf("failed to stat %s: %w", path, err)
	}

	output = strings.TrimSpace(output)

	size, err := strconv.ParseInt(output, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse size from output %s: %w", output, err)
	}

	return size, nil
}

func PrintDeviceInfo(device *gadb.Device) {
	bold := color.New(color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("Device Serial:", cyan(device.Serial()))
	fmt.Println("Device Information:")
	for key, value := range device.DeviceInfo() {
		fmt.Printf("\t%s = %s\n", bold(key), cyan(value))
	}
}
