package main

import (
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"gopkg.in/ini.v1"
	"./utils"
)

const VERSION = "0.0.1"
const GMU_DEMO_RC = ".gmu_democonfig"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var infoFlag *bool = flag.Bool("info", false,
	"Get current git configuration information.")
var initFlag *bool = flag.Bool("init", false, "Initialization")
var updateFlag *bool = flag.Bool("update", false, "Update")
var usersFlag *bool = flag.Bool("users", false, "Users")
var checkoutFlag string

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

func contains(arr []string, str string) bool {
	fmt.Println("arr len:", len(arr))
	for _, ele := range arr {
		if ele == str {
			return true
		}
	}

	return false
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

	// handle [current]
	curr_sec := cfg.Section("current")
	if !curr_sec.HasKey("name") {
		fmt.Println("has no name key")
		curr_sec.NewKey("name", "");
	} else {
		fmt.Println("has name key")
	}

	if !curr_sec.HasKey("gitconfig") {
		curr_sec.NewKey("gitconfig", home + "\\.gitconfig");
	}

	if !curr_sec.HasKey("sshconfig") {
		curr_sec.NewKey("sshconfig", home + "\\.ssh")
	}

	// update current git user
	var user string = get_current_git_user()
	curr_sec.Key("name").SetValue(user)

	// handle [users]
	users_sec := cfg.Section("users")
	if !users_sec.HasKey("name") {
		users_sec.NewKey("name", user)
	}

	var users string = users_sec.Key("name").String()
	if !strings.Contains(users, user) {
		// new git user
		users = users + " " + user
		users_sec.Key("name").SetValue(users)
	}

	// handle [%git user%]
	new_user_sec := cfg.Section(user)

	var gitconfig string = home
	gitconfig = gitconfig + "\\.gitconfig." + user
	if utils.FileExist(gitconfig) {
		new_user_sec.NewKey("gitconfig", gitconfig)
	}

	var sshconfig string = home
	sshconfig = sshconfig + "\\.ssh." + user
	if utils.FileExist(sshconfig) {
		new_user_sec.NewKey("sshconfig", sshconfig)
	}

	cfg.SaveTo(ini_config)

	return true
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

func users() bool {
	var home string = os.Getenv(`USERPROFILE`)

	ini_config := home + "\\" + GMU_DEMO_RC;
	if !utils.FileExist(ini_config) {
		return false
	}

	cfg, err := ini.Load(ini_config)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return false
	}

	curr_user := cfg.Section("current").Key("name").String()
	users := cfg.Section("users").Key("name").String()

	user_arr := strings.Fields(users)
	for _, ele := range user_arr {
		if ele == curr_user {
			fmt.Println("*", ele)
		} else {
			fmt.Println(" ", ele)
		}
	}

	return true
}

func checkout_user(user string) bool {
	var home string = os.Getenv(`USERPROFILE`)

	ini_config := home + "\\" + GMU_DEMO_RC;
	if !utils.FileExist(ini_config) {
		return false
	}

	cfg, err := ini.Load(ini_config)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return false
	}

	curr_user := cfg.Section("current").Key("name").String()
	if curr_user == user {
		fmt.Printf("%s is already the current user\n", user)
		return true
	}

	users := cfg.Section("users").Key("name").String()
	if !strings.Contains(users, user) {
		fmt.Printf("%s does not exist, can't checkout", user)
		return false
	}

	var user_gitconfig = cfg.Section(user).Key("gitconfig").String()
	fmt.Println(user_gitconfig)

	if !utils.FileExist(user_gitconfig) {
		fmt.Printf("not find %s's .gitconfig", user)
		return false
	}

	var curr_gitconfig = cfg.Section("current").Key("gitconfig").String()

	ret, _ := utils.CopyFile(curr_gitconfig, user_gitconfig)
	if ret == 0 {
		fmt.Println("checkout .gitconfig failed")
		return false
	}

	var user_sshconfig = cfg.Section(user).Key("sshconfig").String()
	fmt.Println(user_sshconfig)

	var curr_sshconfig = cfg.Section("current").Key("sshconfig").String()

	files, err := ioutil.ReadDir(curr_sshconfig)
	if err != nil {
		fmt.Println("ReadDir failed")
		return false
	}

	for _, f := range files {
		var f_old string = user_sshconfig + "\\" + f.Name()
		var f_new string = curr_sshconfig + "\\" + f.Name()
		fmt.Println(f_old)
		fmt.Println(f_new)
		ret, _ = utils.CopyFile(f_new, f_old)
		if ret == 0 {
			fmt.Println("checkout .ssh failed")
			return false
		}
	}

	update_env()

	return true
}

func main() {
	flag.StringVar(&checkoutFlag, "c", "", "Set `user` as the current user")
	flag.Parse()

	if *versionFlag {
		fmt.Println("version:", VERSION)
	} else if *infoFlag {
		get_git_config_info()
	} else if *initFlag {
		fmt.Println("Init...")
		init_env()
	} else if *updateFlag {
		fmt.Println("Update...")
		update_env()
	} else if *usersFlag {
		users()
	} else if checkoutFlag != "" {
		fmt.Println("Checkout...")
		checkout_user(checkoutFlag)
	} else {
		fmt.Println("Print usage.")
	}
}
