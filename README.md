# ChargePing ⚡

Simple battery monitor for Termux that detects charging progress and plays a sound when battery percentage increases.

---

## 🚀 Features

- Monitor battery percentage in real-time
- Detect charging status
- Play sound (mpv) when battery increases
- Timeout protection (avoid hang when Termux:API not responding)
- Embedded audio (no external file needed)

---

## 📦 Requirements

- Termux
- Termux:API app installed
- Packages:
  - golang
  - termux-api
  - jq
  - mpv

Install:
```bash
pkg install golang termux-api jq mpv -y
```
---

## ⚙️ Build
```
go build -o chargeping
```
---

## ▶️ Run
```
./chargeping
```
---

## 🔊 How it works
- Reads battery data via  `termux-battery-status`
- Uses  `timeout`  to prevent hanging
- Plays sound using  `mpv`  when battery increases
---

## ⚠️ Notes
- Make sure Termux:API app is installed and opened at least once
- Grant all required permissions
- Audio file is embedded into the binary
---

## 📌 Example Output
```
[*] Starting battery monitor...
Current: 95 | 13:20:11
Current: 95 -> 96 | 13:22:03
```
---
