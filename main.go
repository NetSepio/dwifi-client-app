// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	contract "github.com/NetSepio/dwifi-client/callcontract"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/fatih/color"
// 	"golang.org/x/net/websocket"
// )

// type DeviceInfo struct {
// 	MACAddress         string        `json:"macAddress"`
// 	IPAddress          string        `json:"ipAddress"`
// 	ConnectedAt        time.Time     `json:"connectedAt"`
// 	TotalConnectedTime time.Duration `json:"totalConnectedTime"`
// 	Connected          bool          `json:"connected"`
// 	LastChecked        time.Time     `json:"lastChecked"`
// 	DefaultGateway     string        `json:"defaultGateway"`
// 	Manufacturer       string        `json:"manufacturer"`
// 	InterfaceName      string        `json:"interfaceName"`
// 	HostSSID           string        `json:"hostSSID"`
// 	Chain              string        `json:"chain_name"`
// }

// type WiFiData struct {
// 	ID          uint         `json:"id" gorm:"primaryKey"`
// 	Gateway     string       `json:"gateway"`
// 	Status      []DeviceInfo `json:"status"`
// 	CreatedAt   time.Time    `json:"created_at"`
// 	UpdatedAt   time.Time    `json:"updated_at"`
// 	Password    string       `json:"password"`
// 	PricePerMin string       `json:"price_per_min"`
// }

// type NearbyNetwork struct {
// 	SSID           string `json:"ssid"`
// 	SignalStrength int    `json:"signalStrength"`
// 	BSSID          string `json:"bssid"`
// 	Channel        string `json:"channel"`
// 	Frequency      string `json:"frequency"`
// 	Rate           string `json:"rate"`
// 	Security       string `json:"security"`
// }

// type CurrentWiFi struct {
// 	SSID      string `json:"ssid"`
// 	BSSID     string `json:"bssid"`
// 	Signal    string `json:"signal"`
// 	Frequency string `json:"frequency"`
// 	Rate      string `json:"rate"`
// 	Security  string `json:"security"`
// }

// var wifiDataList []WiFiData
// var mu sync.Mutex

// func ConnectToWiFi(ssid, password string) error {
// 	cmd := exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to WiFi network: %s, %w", string(output), err)
// 	}
// 	return nil
// }

// func ScanNearbyNetworks() ([]NearbyNetwork, error) {
// 	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SIGNAL,BSSID,CHAN,FREQ,RATE,SECURITY", "dev", "wifi")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var networks []NearbyNetwork
// 	lines := strings.Split(string(output), "\n")
// 	for _, line := range lines {
// 		if line == "" {
// 			continue
// 		}
// 		parts := strings.Split(line, ":")
// 		if len(parts) < 7 {
// 			continue
// 		}
// 		signalStrength, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			continue
// 		}
// 		networks = append(networks, NearbyNetwork{
// 			SSID:           parts[0],
// 			SignalStrength: signalStrength,
// 			BSSID:          parts[2],
// 			Channel:        parts[3],
// 			Frequency:      parts[4],
// 			Rate:           parts[5],
// 			Security:       parts[6],
// 		})
// 	}
// 	return networks, nil
// }

// func getCurrentWiFi(w http.ResponseWriter, r *http.Request) {
// 	cmd := exec.Command("nmcli", "-t", "-f", "NAME", "connection", "show", "--active")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		log.Printf("Failed to get current WiFi connection: %v", err)
// 		http.Error(w, "Failed to get current WiFi connection", http.StatusInternalServerError)
// 		return
// 	}

// 	ssid := strings.TrimSpace(string(output))
// 	response := struct {
// 		SSID string `json:"ssid"`
// 	}{
// 		SSID: ssid,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func disconnectAndPay(w http.ResponseWriter, r *http.Request) {
// 	// var requestData struct {
// 	// 	SSID          string `json:"ssid"`
// 	// 	WalletAddress string `json:"walletAddress"`
// 	// 	PrivateKey    string `json:"privateKey"`
// 	// }
// 	// if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
// 	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 	// 	return
// 	// }

// 	// Set up Ethereum client
// 	client, err := ethclient.Dial(os.Getenv("YOUR_ETHEREUM_NODE_URL"))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to connect to the Ethereum client: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Convert wallet address string to common.Address
// 	userWalletAddress := common.HexToAddress("0x1F1e15634BaD10E1105A0CA9cbBA26E6f986fB28")

// 	// // Call the smart contract
// 	// err = contract.MintAndPay(client, userAddress)
// 	// if err != nil {
// 	// 	http.Error(w, fmt.Sprintf("Error in payment process: %v", err), http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// // Disconnect from WiFi
// 	// cmd := exec.Command("nmcli", "device", "disconnect", "wifi")
// 	// err = cmd.Run()
// 	// if err != nil {
// 	// 	http.Error(w, fmt.Sprintf("Error disconnecting from WiFi: %v", err), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	color.Yellow("Initiating disconnection and payment process...")

// 	err0 := contract.MintAndPay(client, userWalletAddress)
// 	if err0 != nil {
// 		color.Red("Error in payment process: %v", err)
// 		return
// 	}

// 	// Disconnect from WiFi
// 	cmd := exec.Command("nmcli", "device", "disconnect", "wifi")
// 	err = cmd.Run()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error disconnecting from WiFi: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// fmt.Fprintf(w, "Successfully disconnected from WiFi and completed payment")
// }

// func handleWebSocket(ws *websocket.Conn) {
// 	var buffer string

// 	for {
// 		var msg string
// 		if err := websocket.Message.Receive(ws, &msg); err != nil {
// 			color.Red("Failed to receive message: %v", err)
// 			return
// 		}

// 		// log.Printf("Received message: %s\n", msg)
// 		buffer += msg

// 		if json.Valid([]byte(buffer)) {
// 			var newWiFiData WiFiData
// 			if err := json.Unmarshal([]byte(buffer), &newWiFiData); err != nil {
// 				color.Red("Failed to unmarshal message: %v", err)
// 				buffer = ""
// 				continue
// 			}

// 			mu.Lock()
// 			found := false
// 			for i, data := range wifiDataList {
// 				if data.ID == newWiFiData.ID {
// 					wifiDataList[i] = newWiFiData
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				wifiDataList = append(wifiDataList, newWiFiData)
// 			}
// 			mu.Unlock()

// 			// color.Yellow("Received WiFi data for ID: %d, SSID: %s, Gateway: %s,Price_per_min: %s, Chain: %s",
// 			// 	newWiFiData.ID, newWiFiData.Status[0].HostSSID, newWiFiData.Gateway, newWiFiData.PricePerMin, newWiFiData.Status[0].Chain)

// 			buffer = ""
// 		}
// 	}
// }

// func main() {
// 	color.Cyan("Connecting to WebSocket...")
// 	origin := "http://localhost/"
// 	url := "wss://dev.gateway.erebrus.io/api/v1.0/nodedwifi/stream"
// 	ws, err := websocket.Dial(url, "", origin)
// 	if err != nil {
// 		color.Red("Failed to connect to WebSocket: %v", err)
// 		return
// 	}
// 	defer ws.Close()

// 	fmt.Println(color.HiMagentaString("__________________________________________________________________________________________________\n"))

// 	go handleWebSocket(ws)

// 	http.HandleFunc("/wifi-networks", func(w http.ResponseWriter, r *http.Request) {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(wifiDataList)
// 	})

// 	http.HandleFunc("/scan-nearby-networks", func(w http.ResponseWriter, r *http.Request) {
// 		cmd := exec.Command("nmcli", "-t", "-f", "SSID,SIGNAL", "device", "wifi", "list")
// 		output, err := cmd.Output()
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("failed to scan nearby networks: %v", err), http.StatusInternalServerError)
// 			return
// 		}

// 		lines := strings.Split(string(output), "\n")
// 		var networks []map[string]interface{}
// 		for _, line := range lines {
// 			fields := strings.Split(line, ":")
// 			if len(fields) < 2 {
// 				continue
// 			}
// 			networks = append(networks, map[string]interface{}{
// 				"SSID":           fields[0],
// 				"SignalStrength": fields[1],
// 			})
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(networks)
// 	})

// 	http.HandleFunc("/current-wifi", getCurrentWiFi)

// 	http.HandleFunc("/connect-wifi", func(w http.ResponseWriter, r *http.Request) {
// 		var requestData struct {
// 			SSID string `json:"ssid"`
// 		}
// 		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
// 			http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 			return
// 		}

// 		mu.Lock()
// 		var password string
// 		for _, data := range wifiDataList {
// 			for _, status := range data.Status {
// 				if status.HostSSID == requestData.SSID {
// 					password = data.Password
// 					break
// 				}
// 			}
// 		}
// 		mu.Unlock()

// 		err := ConnectToWiFi(requestData.SSID, password)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Fprintln(w, "WiFi connection attempt completed. Please check your network status.")
// 	})

// 	http.HandleFunc("/disconnect-and-pay", disconnectAndPay)

// 	fs := http.FileServer(http.Dir("./frontend"))
// 	http.Handle("/", fs)

// 	fmt.Println("Server is running on port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	contract "github.com/NetSepio/dwifi-client/callcontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"golang.org/x/net/websocket"
)

type DeviceInfo struct {
	MACAddress         string        `json:"macAddress"`
	IPAddress          string        `json:"ipAddress"`
	ConnectedAt        time.Time     `json:"connectedAt"`
	TotalConnectedTime time.Duration `json:"totalConnectedTime"`
	Connected          bool          `json:"connected"`
	LastChecked        time.Time     `json:"lastChecked"`
	DefaultGateway     string        `json:"defaultGateway"`
	Manufacturer       string        `json:"manufacturer"`
	InterfaceName      string        `json:"interfaceName"`
	HostSSID           string        `json:"hostSSID"`
	ChainName          string        `json:"chain_name"`
}

type WiFiData struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Gateway     string       `json:"gateway"`
	Status      []DeviceInfo `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Password    string       `json:"password"`
	PricePerMin string       `json:"price_per_min"`
}

type NearbyNetwork struct {
	SSID           string `json:"ssid"`
	SignalStrength int    `json:"signalStrength"`
	BSSID          string `json:"bssid"`
	Channel        string `json:"channel"`
	Frequency      string `json:"frequency"`
	Rate           string `json:"rate"`
	Security       string `json:"security"`
	Gateway        string `json:"gateway"`
	PricePerMin    string `json:"price_per_min"`
	ChainName      string `json:"chain_name"`
}

type CurrentWiFi struct {
	SSID      string `json:"ssid"`
	BSSID     string `json:"bssid"`
	Signal    string `json:"signal"`
	Frequency string `json:"frequency"`
	Rate      string `json:"rate"`
	Security  string `json:"security"`
}

var wifiDataList []WiFiData
var mu sync.Mutex

func ConnectToWiFi(ssid, password string) error {
	cmd := exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to WiFi network: %s, %w", string(output), err)
	}
	return nil
}

func ScanNearbyNetworks() ([]NearbyNetwork, error) {
	cmd := exec.Command("nmcli", "-t", "-f", "SSID,SIGNAL,BSSID,CHAN,FREQ,RATE,SECURITY", "dev", "wifi")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var networks []NearbyNetwork
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) < 7 {
			continue
		}
		signalStrength, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		networks = append(networks, NearbyNetwork{
			SSID:           parts[0],
			SignalStrength: signalStrength,
			BSSID:          parts[2],
			Channel:        parts[3],
			Frequency:      parts[4],
			Rate:           parts[5],
			Security:       parts[6],
		})
	}

	mu.Lock()
	defer mu.Unlock()
	for i, network := range networks {
		for _, data := range wifiDataList {
			for _, status := range data.Status {
				if status.HostSSID == network.SSID {
					networks[i].Gateway = data.Gateway
					networks[i].PricePerMin = data.PricePerMin
					networks[i].ChainName = status.ChainName
				}
			}
		}
	}

	return networks, nil
}

func getCurrentWiFi(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("nmcli", "-t", "-f", "NAME", "connection", "show", "--active")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to get current WiFi connection: %v", err)
		http.Error(w, "Failed to get current WiFi connection", http.StatusInternalServerError)
		return
	}

	ssid := strings.TrimSpace(string(output))
	response := struct {
		SSID string `json:"ssid"`
	}{
		SSID: ssid,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func disconnectAndPay(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SSID          string `json:"ssid"`
		WalletAddress string `json:"walletAddress"`
		PrivateKey    string `json:"privateKey"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	client, err := ethclient.Dial("https://rpcpc1-qa.agung.peaq.network")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to the Ethereum client: %v", err), http.StatusInternalServerError)
		return
	}

	privateKey, err := crypto.HexToECDSA(requestData.PrivateKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading private key: %v", err), http.StatusInternalServerError)
		return
	}

	userWalletAddress := common.HexToAddress(requestData.WalletAddress)
	txHash, minedHash, err := contract.MintAndPay(client, userWalletAddress, privateKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in payment process: %v", err), http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE", "device")
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting device list: %v", err), http.StatusInternalServerError)
		return
	}

	var wifiDevice string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 && parts[1] == "wifi" {
			wifiDevice = parts[0]
			break
		}
	}

	if wifiDevice == "" {
		http.Error(w, "No WiFi device found", http.StatusInternalServerError)
		return
	}

	cmd = exec.Command("nmcli", "device", "disconnect", wifiDevice)
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error disconnecting from WiFi: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message   string `json:"message"`
		TxHash    string `json:"txHash"`
		MinedHash string `json:"minedHash"`
	}{
		Message:   "Successfully disconnected from WiFi and completed payment",
		TxHash:    txHash,
		MinedHash: minedHash,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleWebSocket(ws *websocket.Conn) {
	var buffer string

	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			color.Red("Failed to receive message: %v", err)
			return
		}

		buffer += msg

		if json.Valid([]byte(buffer)) {
			var newWiFiData WiFiData
			if err := json.Unmarshal([]byte(buffer), &newWiFiData); err != nil {
				color.Red("Failed to unmarshal message: %v", err)
				buffer = ""
				continue
			}

			mu.Lock()
			found := false
			for i, data := range wifiDataList {
				if data.ID == newWiFiData.ID {
					wifiDataList[i] = newWiFiData
					found = true
					break
				}
			}
			if !found {
				wifiDataList = append(wifiDataList, newWiFiData)
			}
			color.Yellow("Received WiFi data for ID: %d, SSID: %s, Gateway: %s,Price_per_min: %s, Chain: %s",
				newWiFiData.ID, newWiFiData.Status[0].HostSSID, newWiFiData.Gateway, newWiFiData.PricePerMin, newWiFiData.Status[0].ChainName)
			mu.Unlock()

			buffer = ""
		}
	}
}

func main() {
	color.Cyan("Connecting to WebSocket...")
	origin := "http://localhost/"
	url := "wss://dev.gateway.erebrus.io/api/v1.0/nodedwifi/stream"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		color.Red("Failed to connect to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	fmt.Println(color.HiMagentaString("__________________________________________________________________________________________________\n"))

	go handleWebSocket(ws)

	http.HandleFunc("/wifi-networks", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wifiDataList)
	})

	http.HandleFunc("/scan-nearby-networks", func(w http.ResponseWriter, r *http.Request) {
		networks, err := ScanNearbyNetworks()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to scan nearby networks: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(networks)
	})

	http.HandleFunc("/connect-wifi", func(w http.ResponseWriter, r *http.Request) {
		var requestData struct {
			SSID string `json:"ssid"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		mu.Lock()
		var password string
		for _, data := range wifiDataList {
			for _, status := range data.Status {
				if status.HostSSID == requestData.SSID {
					password = data.Password
					break
				}
			}
		}
		mu.Unlock()

		err := ConnectToWiFi(requestData.SSID, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "WiFi connection attempt completed. Please check your network status.")
	})
	http.HandleFunc("/current-wifi", getCurrentWiFi)
	http.HandleFunc("/disconnect-and-pay", disconnectAndPay)

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
