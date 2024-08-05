// package wifi

// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"strings"
// )

// type AccessPoint struct {
// 	SSID           string `json:"ssid"`
// 	BSSID          string `json:"bssid"`
// 	SignalStrength string `json:"signalStrength"`
// }

// func ScanAccessPoints() ([]AccessPoint, error) {
// 	log.Println("Scanning for access points...")
// 	cmd := exec.Command("nmcli", "-t", "-f", "SSID,BSSID,SIGNAL", "device", "wifi", "list")
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Printf("Error executing nmcli command: %v", err)
// 		return nil, fmt.Errorf("failed to scan: %v - %s", err, string(output))
// 	}

// 	log.Printf("Raw nmcli output: %s", string(output))
// 	return parseNmcliOutput(string(output)), nil
// }

// func parseNmcliOutput(output string) []AccessPoint {
// 	var aps []AccessPoint
// 	lines := strings.Split(output, "\n")

// 	log.Printf("Parsing %d lines of nmcli output", len(lines))
// 	for _, line := range lines {
// 		if line == "" {
// 			continue
// 		}
// 		// Split the line at the last colon
// 		lastColon := strings.LastIndex(line, ":")
// 		if lastColon == -1 {
// 			log.Printf("Unexpected format in line: %s", line)
// 			continue
// 		}

// 		ssidBssid := line[:lastColon]
// 		signal := line[lastColon+1:]

// 		// Split SSID and BSSID
// 		colonIndex := strings.LastIndex(ssidBssid, ":")
// 		if colonIndex == -1 {
// 			log.Printf("Unexpected format in SSID:BSSID part: %s", ssidBssid)
// 			continue
// 		}

// 		ssid := strings.TrimSpace(ssidBssid[:colonIndex])
// 		bssid := strings.ReplaceAll(ssidBssid[colonIndex+1:], "\\:", ":")

// 		ap := AccessPoint{
// 			SSID:           ssid,
// 			BSSID:          bssid,
// 			SignalStrength: signal,
// 		}
// 		aps = append(aps, ap)
// 		log.Printf("Parsed AP: %+v", ap)
// 	}
// 	return aps
// }

// func ConnectToAP(ssid, password string) error {
// 	log.Printf("Attempting to connect to SSID: %s", ssid)
// 	cmd := exec.Command("nmcli", "device", "wifi", "connect", ssid, "password", password)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Printf("Error connecting to AP: %v - %s", err, string(output))
// 		return fmt.Errorf("failed to connect: %v - %s", err, string(output))
// 	}
// 	log.Printf("Connection output: %s", string(output))
// 	return nil
// }

package wifi

import (
	"fmt"
	"os/exec"
	"strings"
)

type AccessPoint struct {
	SSID           string `json:"ssid"`
	BSSID          string `json:"bssid"`
	SignalStrength string `json:"signalStrength"`
}

func ScanAccessPoints() ([]AccessPoint, error) {
	cmd := exec.Command("nmcli", "-t", "-f", "SSID,BSSID,SIGNAL", "device", "wifi", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %v", err)
	}

	return parseNmcliOutput(string(output)), nil
}

func parseNmcliOutput(output string) []AccessPoint {
	var aps []AccessPoint
	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) >= 3 {
			ap := AccessPoint{
				SSID:           parts[0],
				BSSID:          parts[1],
				SignalStrength: parts[2],
			}
			aps = append(aps, ap)
		}
	}
	return aps
}
