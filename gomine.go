package main

import (
    "fmt"
    "runtime"
    "os/exec"
)

func main() {
    // Get current operating system
    os := runtime.GOOS

    // Choose appropriate miner script
    var cmd *exec.Cmd
    if os == "windows" {
        cmd = exec.Command("powershell.exe", "-File", "monero-miner.ps1")
    } else if os == "linux" {
        cmd = exec.Command("./monero-miner.sh")
    } else {
        fmt.Println("Unsupported operating system.")
        return
    }

    // Run miner script
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    fmt.Printf("Output: %s\n", output)
}
