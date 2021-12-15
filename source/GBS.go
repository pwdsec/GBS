package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	setConsoleTitle()
	fmt.Println("[+] Made by @pwdsec")
	fmt.Println("[+] Watching folder for changes...")
	for {
		deleteFile()
	}
}

func deleteFile() {
	a, _ := os.UserCacheDir()
	appDataLocalRobloxPath := a + "\\Roblox\\GlobalBasicSettings_13.xml"
	if _, err := os.Stat(appDataLocalRobloxPath); err == nil {
		os.Remove(appDataLocalRobloxPath)
		fmt.Println("GlobalBasicSettings_13.xml file deleted")
	}
}

func setConsoleTitle() {
	err := exec.Command("cmd", "/C", "title GBS").Run()
	if err != nil {
		fmt.Println(err)
	}
}

