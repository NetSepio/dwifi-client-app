document.addEventListener('DOMContentLoaded', () => {
    const ws = new WebSocket('wss://dev.gateway.erebrus.io/api/v1.0/nodedwifi/stream');

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        updateWiFiDataList(data);
    };

    function updateWiFiDataList(data) {
        const wifiDataList = document.getElementById('wifi-data-list');
        wifiDataList.textContent = JSON.stringify(data, null, 2);
    }

    async function fetchCurrentWiFi() {
        const response = await fetch('/current-wifi');
        const data = await response.json();
        document.getElementById('current-wifi-info').textContent = JSON.stringify(data, null, 2);
    }

    async function scanNearbyNetworks() {
        const response = await fetch('/scan-nearby-networks');
        const networks = await response.json();
        const networksList = document.getElementById('networks-list');
        networksList.innerHTML = networks.map(network => `
            <li>
                SSID: ${network.ssid} <br>
                Signal Strength: ${network.signalStrength} <br>
                BSSID: ${network.bssid} <br>
                Channel: ${network.channel} <br>
                Frequency: ${network.frequency} <br>
                Rate: ${network.rate} <br>
                Security: ${network.security}
            </li>
        `).join('');
    }

    async function connectToWiFi() {
        const ssid = document.getElementById('ssid').value;
        const response = await fetch('/connect-wifi', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ ssid })
        });
        const result = await response.text();
        alert(result);
    }

    fetchCurrentWiFi();

    window.scanNearbyNetworks = scanNearbyNetworks;
    window.connectToWiFi = connectToWiFi;
});
