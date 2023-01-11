package main

import (
    "fmt"
    "os/exec"
    "runtime"
)

func checkWindowsCompatibility() {
    // Check for PowerShell
    _, err := exec.Command("powershell.exe", "-Command", "Get-Command powershell").Output()
    if err != nil {
        fmt.Println("PowerShell is not installed or not accessible. Please install it or add it to your PATH.")
        return
    }

    // Check for Visual C++ Redistributable
    output, err := exec.Command("powershell.exe", "-Command", "Get-Package -Name vcredist* -AllVersions").Output()
    if err != nil {
        fmt.Println("Visual C++ Redistributable is not installed. Please install it.")
        return
    }
    if string(output) == "" {
        fmt.Println("Visual C++ Redistributable is not installed. Please install it.")
        return
    }

    // Check for .NET Framework
    _, err = exec.Command("powershell.exe", "-Command", "Test-Path 'C:\Windows\Microsoft.NET\Framework\'").Output()
    if err != nil {
        fmt.Println(".NET Framework is not installed or not accessible. Please install it.")
        return
    }
}

func checkLinuxCompatibility() {
    // Check for package manager (apt or yum)
    _, errApt := exec.Command("apt-get").Output()
    _, errYum := exec.Command("yum").Output()
    if errApt != nil && errYum != nil {
        fmt.Println("No package manager (apt or yum) found. Please install one or check your PATH.")
        return
    }
}

func checkMacCompatibility() {
    // Check for Xcode Command Line Tools
    _, err := exec.Command("xcode-select", "-p").Output()
    if err != nil {
        fmt.Println("Xcode Command Line Tools are not installed. Please install them.")
        return
    }
}

func main() {
    // Get current operating system
    os := runtime.GOOS

    // Check OS compatibility
    switch os {
    case "windows":
        checkWindowsCompatibility()
    case "linux":
        checkLinuxCompatibility()
    case "darwin":
        checkMacCompatibility()
    default:
        fmt.Println("Unsupported operating system.")
        return
    }

    // Choose appropriate miner script
    var cmd *exec.Cmd
    if os == "windows" {
        cmd = exec.Command("powershell.exe", "-File", "monero-miner.ps1")
    } else if os == "linux" {
        cmd = exec.Command("./monero-miner.sh")
    } else if os == "darwin" {
        cmd = exec.Command("./monero-miner.sh")
    } else {
        fmt.Println("Unsupported operating system.")
        return
    }

    // Run miner script
    output, err := cmd.CombinedOutput()
    if err != nil {
       

