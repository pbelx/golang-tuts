package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	for {
		fmt.Println("Welcome to NetSolutions")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Press 1 to get external IP")
		fmt.Println("Press 2 to check if online")
		fmt.Println("Press 3 to get local ip")
		fmt.Println("Press x to Exit")
		scanner.Scan()
		msg := scanner.Text()
		ans := msg
		switch ans {
		case "1":
			fmt.Println("external ip is ...")
			cmd, _ := exec.Command("curl", "ifconfig.io").Output()

			fmt.Println(string(cmd))
			fmt.Println("*****")
		case "2":
			_, err := exec.Command("ping", "-c 3", "8.8.8.8").Output()
			if err != nil {
				fmt.Println("*****")
				fmt.Println("you offline ...")
				fmt.Println("******")
			} else {
				fmt.Println("*****")
				fmt.Println("online")
				fmt.Println("*****")
			}
		case "3":
			fmt.Println("local ip is ...")
			c2 := "/sbin/ifconfig | grep inet | grep -v inet6"
			cmd, _ := exec.Command("bash", "-c", c2).Output()
			fmt.Println(string(cmd))
			fmt.Println("******")
		case "x":
			fmt.Println("bye bye ...")
			os.Exit(0)
		default:
			fmt.Println("no option")
		}
	}
}
