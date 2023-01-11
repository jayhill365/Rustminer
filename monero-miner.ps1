# Download and install Visual C++ Redistributable
$vcredist = "https://aka.ms/vs/16/release/vc_redist.x64.exe"
Invoke-WebRequest -Uri $vcredist -OutFile vc_redist.exe
Start-Process -FilePath .\vc_redist.exe -ArgumentList '/install', '/quiet', '/norestart' -Wait

# Download Monero miner
$minerUrl = "https://github.com/xmrig/xmrig/releases/download/v6.7.1/xmrig-6.7.1-win64.zip"
Invoke-WebRequest -Uri $minerUrl -OutFile xmrig.zip
Expand-Archive xmrig.zip

# Set sending address
$config = Get-Content ".\xmrig-6.7.1\config.json"
$config = $config -replace '"YOUR_WALLET_ADDRESS"', ('"' + "YOUR_MONERO_WALLET_ADDRESS" + '"')
$config | Set-Content ".\xmrig-6.7.1\config.json"

# Start miner
Start-Process -FilePath ".\xmrig-6.7.1\xmrig.exe" -ArgumentList "--config=.\xmrig-6.7.1\config.json" -NoNewWindow -Wait
