package main

import (
  "os"
 "fmt"
 "os/exec"
 "bytes"
 "time"
 "strconv"
 _ "embed"
)

//go:embed beep.mp3
var beepData []byte

func checkBins(binaries []string) bool {
  for _, bin := range binaries {
    if _, err := exec.LookPath(bin); err != nil {
      fmt.Printf("Error: %s not found!\n", bin)
      return false
    }
  }
  return true
}

func getStatus() bool {
  cmd := exec.Command("bash", "-c", "timeout 3s termux-battery-status | jq .status")
  
  out, err := cmd.Output()
  out = bytes.TrimSpace(out)
  if err != nil {
    return false
  }
  
   if string(out) == `"CHARGING"` {
    return true
   }
    return false
}

func getTime() string {
  out, _ := exec.Command("date", "+%H:%M").Output()
  return string(bytes.TrimSpace(out))
}

func getPercentage() (int, error) {
  cmd := exec.Command("bash", "-c", "timeout 3s termux-battery-status | jq .percentage")
  
  out, err := cmd.Output()
  if err != nil {
    return 0, err
  }
  
  out = bytes.TrimSpace(out)
  return strconv.Atoi(string(out))
}

func playSound() {
	f, err := os.CreateTemp("", "beep-*.mp3")
	if err != nil {
		return
	}
	defer f.Close()

	f.Write(beepData)

	exec.Command(
		"mpv",
		f.Name(),
		"--no-video",
		"--quiet",
	).Start()
}

func main() {
    required := []string{"termux-battery-status", "jq", "mpv"}
    if !checkBins(required) {
      return
    }
    
    if !getStatus() {
      fmt.Println("Battery not charging or Termux:API not running.")
      return
    }
    
    prev, err := getPercentage()
     if err != nil {
       fmt.Println("Failed to get battery percentage.")
        return
    }
    lastTime := getTime()
    fmt.Println("[*] Starting battery monitor...\n[+] Current:", prev, "|", lastTime)
    for {
      time.Sleep(1 * time.Second)
      curr, _ := getPercentage()
      if prev  == 0 || curr == 0 {
        fmt.Println("[-] Termux:API Stopped.")
        return
      } else if curr > prev {
        timeNow := getTime()
        playSound()
        fmt.Println("[+] Current:", curr, "|", timeNow)
      } else if !getStatus() || curr == 100 {
        fmt.Println("[-] Stopped Charging.")
        return
      }
      prev = curr
    }
    
}