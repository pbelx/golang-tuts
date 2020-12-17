package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

//main area
func main() {

	//infinite loop
	for {
		menu := `
	  **BelNet Utils**
	  Choose options Below
	   **************
		
	1 to get external ip
	2 check if online
	3 gets local ip
	4 gets tcp connections
	5 gets open tcp ports
	0 Exit stage Left `
		print(menu)

		//konsole commands
		pingc := "ping -c 2 8.8.8.8"                        //2 check if online ?
		ifc := "/sbin/ifconfig | grep inet | grep -v inet6" //3 interface ips
		ntcp := "ss -ant"                                   //4 tcp connections
		getip := "curl ifconfig.io -v"                      //1 external ip
		getports := "netstat -antp | grep TEN"              //open port ss -ltp

		//getinput

		scanner := bufio.NewScanner(os.Stdin)
		print(`
	Wake up, Neo ::`)
		scanner.Scan()
		msg := scanner.Text()
		ans := msg
		switch ans {
		case "1":
			// clear()
			roll(getip)
		case "2":
			rollerr(pingc)
		case "3":
			roll(ifc)
		case "4":
			roll(ntcp)
		case "5":
			roll(getports)
		case "0":
			print(`
		Follow the white Rabbit`)
			os.Exit(0)

		default:
			print("Option not allowed")
			print("Follow the white rabbit")
		}
	}
}

//print function
func print(msg string) {
	fmt.Println(msg)
}

//clear screen

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//konsole commands exec no errors

func roll(command string) {
	clear()
	cmd, _ := exec.Command("bash", "-c", command).Output()
	fmt.Println(`
	*********
	Matrix REply
	************`)
	fmt.Println(string(cmd))
}

//kommands exec with err returned

func rollerr(command string) {
	_, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		print("Failed mission")
	} else {
		print("w00t success achieved")
	}
}
