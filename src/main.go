package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/zserge/lorca"
)

//go:embed www
var fs embed.FS

func getAppPath() string {
	ex, err := os.Executable()
  if err != nil {
    panic(err)
	}
  exPath := filepath.Dir(ex)
	return exPath
}

func formatPath(paths []string) string {
	return strings.Join(paths[:], string(os.PathSeparator))
}

func isPorkInstalled() bool {
	os.Chdir(getAppPath())

	_, err := os.Stat(getAppPath() + string(os.PathSeparator) + "powercord")
	return !os.IsNotExist(err)
}

func canDo() string {
	node := exec.Command("node", "-v")
	node.Stdout = os.Stdout
	node.Stderr = os.Stderr
	node.Stdin = os.Stdin
	err := node.Run()
	if err != nil {
		return "{\"text\": \"Install node js\", \"link\": \"https://nodejs.org/en/\"}"
	}

	npm := exec.Command("npm", "-v")
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr
	npm.Stdin = os.Stdin
	npmErr := npm.Run()
	if npmErr != nil {
		return "{\"text\": \"Install npm\", \"link\": \"https://nodejs.org/en/\"}"
	}

	git := exec.Command("git", "--version")
	git.Stdout = os.Stdout
	git.Stderr = os.Stderr
	git.Stdin = os.Stdin
	gitErr := git.Run()
	if gitErr != nil {
		return "{\"text\": \"Install git\", \"link\": \"https://git-scm.com/downloads\"}"
	}
	return "{\"ok\": true}"
}

func installPC() string {
  if !isPorkInstalled() {
  	cmd := exec.Command("git", "clone", "https://github.com/powercord-org/powercord")
	
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
  		log.Fatal(err)
		}
		return "Installed powercord"
  } else {
		return "Powercord is already installed"
	}
}

func uninstallPC() string {
  if !isPorkInstalled() {
		return "Powercord isn't installed"
  } else {
		err := os.RemoveAll("powercord")
    if err != nil {
      log.Fatal(err)
    }
		return "Uninstalled powercord"
	}
}

func unplugThePork() string {
  if !isPorkInstalled() {
		return "Powercord isn't installed"
	} else {
		os.Chdir(getAppPath() + string(os.PathSeparator) + "powercord")
		cmd := exec.Command("npm", "run", "unplug")
	
  	cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
  		log.Fatal(err)
		}

		return "Unplugged powercord"
	}
}

func plugThePork() string {
  if !isPorkInstalled() {
		return "Powercord isn't installed"
	} else {
		os.Chdir(getAppPath() + string(os.PathSeparator) + "powercord")
		cmd := exec.Command("npm", "i")
	
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
  	cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
  		log.Fatal(err)
		}

		plug := exec.Command("npm", "run", "plug")
	
		plug.Stdin = os.Stdin
		plug.Stderr = os.Stderr
  	plug.Stdout = os.Stdout
		errPlug := plug.Run()
		if errPlug != nil {
  		log.Fatal(errPlug)
		}
		return "Plugged powercord"
	}
}

func getThemeDownloader() string {
	if isPorkInstalled() {
		_, error := os.Stat(getAppPath() + string(os.PathSeparator) + formatPath([]string {"powercord", "src", "Powercord", "plugins", "PowercordThemeDownloader"}))
		if os.IsNotExist(error) {
			os.Chdir(getAppPath() + string(os.PathSeparator) + formatPath([]string {"powercord", "src", "Powercord", "plugins"}))
  		cmd := exec.Command("git", "clone", "https://github.com/ploogins/PowercordThemeDownloader")
	
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			err := cmd.Run()
			if err != nil {
  			log.Fatal(err)
			}
			return "Installed theme downloader"
		} else {
			return "Theme downloader is already installed"
		}
  } else {
		return "Powercord is already installed"
	}
}

func getPluginDownloader() string {
	if isPorkInstalled() {
		_, error := os.Stat(getAppPath() + string(os.PathSeparator) + formatPath([]string {"powercord", "src", "Powercord", "plugins", "PowercordPluginDownloader"}))
		if os.IsNotExist(error) {
			os.Chdir(getAppPath() + string(os.PathSeparator) + formatPath([]string {"powercord", "src", "Powercord", "plugins"}))
  		cmd := exec.Command("git", "clone", "https://github.com/LandenStephenss/PowercordPluginDownloader")
	
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			err := cmd.Run()
			if err != nil {
  			log.Fatal(err)
			}
			return "Installed plugin downloader"
		} else {
			return "Plugin downloader is already installed"
		}
  } else {
		return "Powercord is already installed"
	}
}

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	ui.Bind("installPC", installPC)
	ui.Bind("uninstallPC", uninstallPC)
	ui.Bind("plugPowercord", plugThePork)
	ui.Bind("unplugPowercord", unplugThePork)
	ui.Bind("canDo", canDo)
	ui.Bind("isInstalled", isPorkInstalled)
	ui.Bind("downloadThemePlugin", getThemeDownloader)
	ui.Bind("downloadPluginDownloader", getPluginDownloader)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
