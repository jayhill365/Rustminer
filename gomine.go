package main

import (
    "fmt"
    "os/exec"
    "runtime"
)

func main() {
    // Get current operating system
    os := runtime.GOOS

    var cmd *exec.Cmd
    var err error
    var output []byte

    // Choose appropriate miner script
    switch os {
    case "windows":
        cmd = exec.Command("powershell.exe", "-File", "monero-miner.ps1")
    case "linux":
        cmd = exec.Command("bash", "monero-miner.sh")
    case "darwin":
        cmd = exec.Command("bash", "monero-miner.sh")
    default:
        fmt.Println("Unsupported operating system.")
        return
    }

    // Run miner script
    output, err = cmd.CombinedOutput()
    if err != nil {
        fmt.Println(err)
        fmt.Println(string(output))
        return
    }
    fmt.Println(string(output))
}

       

