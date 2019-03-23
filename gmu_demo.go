package main

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/ini.v1"
	"./utils"
)

const VERSION = "0.0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var infoFlag *bool = flag.Bool("info", false,
	"Get current git configuration information.")
var saveFlag *bool = flag.Bool("save", false,
	"Save the current user's git configuration.")
var saveFlag2 *bool = flag.Bool("save2", false,
	"Save the current user's git configuration.")
var restoreFlag1 *bool = flag.Bool("r1", false,
	"Restore user's git configuration.")
var restoreFlag2 *bool = flag.Bool("r2", false,
	"Restore user's git configuration.")

func get_git_config_info() bool {
	fmt.Println("get_git_config_info()")
	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	fmt.Println("USERPROFILE:", USERPROFILE)

	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"
	fmt.Println(".gitconfig:", git_config_file)

	if utils.FileExist(git_config_file) {
		fmt.Printf("\"%s\" exists\n", git_config_file)
	} else {
		fmt.Printf("\"%s\" does not exist\n", git_config_file)
		return false;
	}

	cfg, err := ini.Load(git_config_file)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return false;
	}

	var user_name string = cfg.Section("user").Key("name").String()
	var user_email string = cfg.Section("user").Key("email").String()

	fmt.Printf("user.name: %s, user.email: %s\n", user_name, user_email)

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh"
	fmt.Println(".ssh:", ssh_dir)

	if utils.FileExist(ssh_dir) {
		fmt.Printf("\"%s\" exists\n", ssh_dir)
	} else {
		fmt.Printf("\"%s\" does not exist\n", ssh_dir)
		return false;
	}

	return true;
}

func save_git_config() bool {
	if !get_git_config_info() {
		fmt.Println("The current git configuration may be incomplete.")
		return false;
	}

	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"
	var git_config_file_new string = USERPROFILE
	git_config_file_new += "\\.gitconfig.1"

	ret, _ := utils.CopyFile(git_config_file_new, git_config_file)
	if ret == 0 {
		// return false;
	}

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh"
	var ssh_dir_new string = USERPROFILE
	ssh_dir_new += "\\.ssh.1"

	fmt.Printf("%s\n", ssh_dir)
	fmt.Printf("%s\n", ssh_dir_new)

	err := os.MkdirAll(ssh_dir_new, os.ModePerm)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("succeed")
	}

	files, err := ioutil.ReadDir(ssh_dir)
	if err != nil {
		fmt.Println("ReadDir failed")
	}

	for _, f := range files {
		var f_old string = ssh_dir + "\\" + f.Name()
		var f_new string = ssh_dir_new + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		utils.CopyFile(f_new, f_old)
	}

	return true;
}

func save_git_config2() bool {
	if !get_git_config_info() {
		fmt.Println("The current git configuration may be incomplete.")
		return false;
	}

	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"
	var git_config_file_new string = USERPROFILE
	git_config_file_new += "\\.gitconfig.2"

	ret, _ := utils.CopyFile(git_config_file_new, git_config_file)
	if ret == 0 {
		// return false;
	}

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh"
	var ssh_dir_new string = USERPROFILE
	ssh_dir_new += "\\.ssh.2"

	fmt.Printf("%s\n", ssh_dir)
	fmt.Printf("%s\n", ssh_dir_new)

	err := os.MkdirAll(ssh_dir_new, os.ModePerm)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("succeed")
	}

	files, err := ioutil.ReadDir(ssh_dir)
	if err != nil {
		fmt.Println("ReadDir failed")
	}

	for _, f := range files {
		var f_old string = ssh_dir + "\\" + f.Name()
		var f_new string = ssh_dir_new + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		utils.CopyFile(f_new, f_old)
	}

	return true;
}

func restore_config1() bool {
	if !get_git_config_info() {
		fmt.Println("The current git configuration may be incomplete.")
		return false;
	}

	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig.1"
	var git_config_file_new string = USERPROFILE
	git_config_file_new += "\\.gitconfig"

	ret, _ := utils.CopyFile(git_config_file_new, git_config_file)
	if ret == 0 {
		// return false;
	}

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh.1"
	var ssh_dir_new string = USERPROFILE
	ssh_dir_new += "\\.ssh"

	fmt.Printf("%s\n", ssh_dir)
	fmt.Printf("%s\n", ssh_dir_new)

	err := os.MkdirAll(ssh_dir_new, os.ModePerm)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("succeed")
	}

	files, err := ioutil.ReadDir(ssh_dir)
	if err != nil {
		fmt.Println("ReadDir failed")
	}

	for _, f := range files {
		var f_old string = ssh_dir + "\\" + f.Name()
		var f_new string = ssh_dir_new + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		utils.CopyFile(f_new, f_old)
	}

	return true;
}

func restore_config2() bool {
	if !get_git_config_info() {
		fmt.Println("The current git configuration may be incomplete.")
		return false;
	}

	var USERPROFILE string = os.Getenv(`USERPROFILE`)
	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig.2"
	var git_config_file_new string = USERPROFILE
	git_config_file_new += "\\.gitconfig"

	ret, _ := utils.CopyFile(git_config_file_new, git_config_file)
	if ret == 0 {
		// return false;
	}

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh.2"
	var ssh_dir_new string = USERPROFILE
	ssh_dir_new += "\\.ssh"

	fmt.Printf("%s\n", ssh_dir)
	fmt.Printf("%s\n", ssh_dir_new)

	err := os.MkdirAll(ssh_dir_new, os.ModePerm)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("succeed")
	}

	files, err := ioutil.ReadDir(ssh_dir)
	if err != nil {
		fmt.Println("ReadDir failed")
	}

	for _, f := range files {
		var f_old string = ssh_dir + "\\" + f.Name()
		var f_new string = ssh_dir_new + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		utils.CopyFile(f_new, f_old)
	}

	return true;
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("version:", VERSION)
	} else if *infoFlag {
		get_git_config_info()
	} else if *saveFlag {
		if !save_git_config() {
			fmt.Println("save git config failed")
		} else {
			fmt.Println("save git config succeeded")
		}
	} else if *saveFlag2 {
		if !save_git_config2() {
			fmt.Println("save git config failed")
		} else {
			fmt.Println("save git config succeeded")
		}
	} else if *restoreFlag1 {
		if !restore_config1() {
			fmt.Println("restore git config failed")
		} else {
			fmt.Println("restore git config succeeded")
		}
	} else if *restoreFlag2 {
		if !restore_config2() {
			fmt.Println("restore git config failed")
		} else {
			fmt.Println("restore git config succeeded")
		}
	}
}
