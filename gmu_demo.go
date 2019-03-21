package main

import (
	"flag"
	"os"
	"fmt"
	"gopkg.in/ini.v1"
)

const VERSION = "0.0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var infoFlag *bool = flag.Bool("info", false,
	"Get current git configuration information.")

func file_exist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		// file exists
		return true;
	} else if os.IsNotExist(err) {
		// file does *not* exist
		return false;
	} else {
		// Schrodinger: file may or may not exist.
		// See err for details

		// Therefore, do *NOT* use !os.IsNotExist(err)
		// to test for file existence
		return false;
	}
}

func get_git_config_info() {
	fmt.Println("get_git_config_info()")
	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	fmt.Println("USERPROFILE:", USERPROFILE)

	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"
	fmt.Println(".gitconfig:", git_config_file)

	if file_exist(git_config_file) {
		fmt.Printf("\"%s\" exists\n", git_config_file)
	} else {
		fmt.Printf("\"%s\" does not exist\n", git_config_file)
		return;
	}

	cfg, err := ini.Load(git_config_file)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return;
	}

	var user_name string = cfg.Section("user").Key("name").String()
	var user_email string = cfg.Section("user").Key("email").String()

	fmt.Println("user.name:", user_name)
	fmt.Println("user.email:", user_email)

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh"
	fmt.Println(".ssh:", ssh_dir)

	if file_exist(ssh_dir) {
		fmt.Printf("\"%s\" exists\n", ssh_dir)
	} else {
		fmt.Printf("\"%s\" does not exist\n", ssh_dir)
		return;
	}
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("version:", VERSION)
	} else if *infoFlag {
		get_git_config_info()
	}
}
