package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	gitHubHost := flag.String("gc", "gh", "host")
	// datePtr := flag.String("d", "NH", "Date")
	// timePtr := flag.String("t", "4H", "Time duration")

	flag.Parse()

	for i := range flag.Args() {
		doGitClone(*gitHubHost, flag.Arg(i))
	}

	fmt.Println("Done Done.")
}

func doGitClone(host string, repo string) {
	configMap := processConfigMap()

	gitHostDir := configMap[host + "_dir"]
	gitPath := getGitPath(repo)
	gitURI := getGitURL(repo)

	folderPath := filepath.Join(gitHostDir, gitPath)
	gitClone(gitURI, folderPath)
	fmt.Println(repo, "Done.")
}

func gitClone(gitURI string, folderPath string) {
	cmd := exec.Command("git", "clone", gitURI, folderPath)
	cmd.Run()
}

func processConfigMap() map[string]string {
	configFile := filepath.Join(os.Getenv("systemdrive")+os.Getenv("homepath"),".dvclirc")
	content, err1 := ioutil.ReadFile(configFile)
	if err1 != nil {
		log.Fatal(err1)
	}

	lines := strings.Split(string(content), "\n")
	configMap := make(map[string]string)

	for i := range lines {
		line := lines[i];
		if strings.Contains(line, "=") {
			config := strings.Split(line, "=")
			value := config[1]
			if strings.Contains(value, "\n") {

			} else {
				configMap[config[0]] = value
			}
		}
	}
	return configMap;
}

func getGitPath(gitArg string) string {
	if strings.Contains(gitArg, "/") {
		return gitArg
	}
	return gitArg + "/" + gitArg
}

func getGitURL(gitArg string) string {
	return "https://github.com/" + getGitPath(gitArg) + ".git"
}