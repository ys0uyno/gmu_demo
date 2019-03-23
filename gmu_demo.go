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
const GMU_DEMO_RC = ".gmu_demorc"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var infoFlag *bool = flag.Bool("info", false,
	"Get current git configuration information.")
var restoreFlag1 *bool = flag.Bool("r1", false,
	"Restore user's git configuration.")
var restoreFlag2 *bool = flag.Bool("r2", false,
	"Restore user's git configuration.")
var initFlag *bool = flag.Bool("init", false, "Initialization")
var updateFlag *bool = flag.Bool("update", false, "Update")

func get_git_config_info() bool {
	fmt.Println("get_git_config_info()")
	var USERPROFILE string = os.Getenv(`USERPROFILE`)

	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"

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

	if utils.FileExist(ssh_dir) {
		fmt.Printf("\"%s\" exists\n", ssh_dir)
	} else {
		fmt.Printf("\"%s\" does not exist\n", ssh_dir)
		return false;
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

func get_current_git_user() string {
	var USERPROFILE string = os.Getenv(`USERPROFILE`)

	var git_config string = USERPROFILE
	git_config += "\\.gitconfig"

	if !utils.FileExist(git_config) {
		fmt.Printf("\"%s\" does not exist\n", git_config)
		return "";
	}

	cfg, err := ini.Load(git_config)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return ""
	}

	var name string = cfg.Section("user").Key("name").String()
	// var email string = cfg.Section("user").Key("email").String()

	return name
}

func update_ini_config(home string) bool {
	ini_config := home + "\\" + GMU_DEMO_RC;
	fmt.Println("ini_confi:", ini_config)
	if !utils.FileExist(ini_config) {
		fmt.Println("create new", ini_config)
		file, err := os.Create(ini_config)
		if err != nil {
			fmt.Println("os.Create failed:", err)
			return false
		}
		defer file.Close()
	}

	cfg, err := ini.Load(ini_config)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return false
	}

	sec := cfg.Section("current")
	if !sec.HasKey("name") {
		fmt.Println("has no name key")
		sec.NewKey("name", "value");
	} else {
		fmt.Println("has name key")
	}

	return true;
}

func save_git_config(home, user string) bool {
	current_git_user := get_current_git_user()
	old_config_file := home + "\\.gitconfig"
	new_config_file := old_config_file + "." + user

	// if current user is equal to 'user' and .gitconfig.user exists,
	// no need save again
	if current_git_user == user && utils.FileExist(new_config_file) {
		fmt.Printf("\"%s\" already exists.\n", new_config_file)
		return true;
	}

	// save .gitconfig as .gitconfig.user_name
	ret, _ := utils.CopyFile(new_config_file, old_config_file)
	if ret == 0 {
		fmt.Println("utils.CopyFile failed")
		return false
	}

	return true
}

func save_ssh_config(home, user string) bool {
	current_git_user := get_current_git_user()
	old_config_file := home + "\\.ssh"
	new_config_file := old_config_file + "." + user

	// if current user is equal to 'user' and .ssh.user exists,
	// no need save again
	if current_git_user == user && utils.FileExist(new_config_file) {
		fmt.Printf("\"%s\" already exists.\n", new_config_file)
		return true;
	}

	err := os.MkdirAll(new_config_file, os.ModePerm)
	if err != nil {
		fmt.Println("os.MkdirAll failed")
		return false
	} else {
		fmt.Println("os.MkdirAll succeeded")
	}

	files, err := ioutil.ReadDir(old_config_file)
	if err != nil {
		fmt.Println("ReadDir failed")
		return false
	}

	for _, f := range files {
		var f_old string = old_config_file + "\\" + f.Name()
		var f_new string = new_config_file + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		utils.CopyFile(f_new, f_old)
	}

	return true
}

func init_env() bool {
	var USERPROFILE string = os.Getenv(`USERPROFILE`)

	var git_config_file string = USERPROFILE
	git_config_file += "\\.gitconfig"

	if !utils.FileExist(git_config_file) {
		fmt.Printf("\"%s\" does not exist\n", git_config_file)
		return false;
	}

	var ssh_dir string = USERPROFILE
	ssh_dir += "\\.ssh"

	if !utils.FileExist(ssh_dir) {
		fmt.Printf("\"%s\" does not exist\n", ssh_dir)
		return false;
	}

	cfg, err := ini.Load(git_config_file)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		return false;
	}

	var user_name string = cfg.Section("user").Key("name").String()
	var user_email string = cfg.Section("user").Key("email").String()

	fmt.Printf("user.name: %s, user.email: %s\n", user_name, user_email)

	if !save_git_config(USERPROFILE, user_name) {
		fmt.Println("save .gitconfig failed.")
		return false
	}

	if !save_ssh_config(USERPROFILE, user_name) {
		fmt.Println("save .ssh failed.")
		return false
	}

	update_ini_config(USERPROFILE)

	return true;
}

func update_env() bool {
	return init_env()
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("version:", VERSION)
	} else if *infoFlag {
		get_git_config_info()
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
	} else if *initFlag {
		fmt.Println("Init...")
		init_env()
	} else if *updateFlag {
		fmt.Println("Update...")
		update_env()
	} else {
		fmt.Println("Print usage.")
	}
}
