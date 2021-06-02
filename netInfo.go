//The purpose of this GO application is to list information on local network devices.
//It primarily lists info on device name, MAC/Hardware addresses, and applicable IPs.
//Additional functionality can be engaged by uncommenting some of the code below.
//Coded by: Matt Ready

package main

import (
	"fmt" //imported to enable Println functionality
	"net" //imported to enable access to networking information/functionality
)

func main() {
	networkInter, _ := net.Interfaces()        //"networkInter" variable created to interact with, and glean data from "net.Interfaces".
	for _, networkAddr := range networkInter { //Only interested in the data returned from the "networkAddr" variable, hence "_". "range" allows for more than one result to be returned, will cycle through them.
		fmt.Println("Name  :", networkAddr.Name)         //Takes the "Name": value returned from "networkAddr" variable (from "range networkInter"), printing the device name on its own line.
		fmt.Println("MAC   :", networkAddr.HardwareAddr) //Named "MAC" for the purpose of this exercise, it takes the "HardwareAddr": value returned from "networkAddr" variable (from "range networkInter"), printing the hardware address on it's own line.
		ip, _ := networkAddr.Addrs()                     //"ip" variable created to interact with "networkAddr.Addrs", which returns IPv4/6 addresses for each interface.
		for _, ip := range ip {                          //as there may be more than one IP per device returned, "range" is used to cycle through them. "ip" variable is used to interact with that data.
			fmt.Println("IP    :", ip) //Takes the value(s) returned from the "range ip" and prints them on their own lines
		}
		//Additional network device info, uncomment if needed.

		//fmt.Println("Index :", networkAddr.Index)
		//fmt.Println("MTU   :", networkAddr.MTU)
		//fmt.Println("Flags :", networkAddr.Flags)
		fmt.Println("===============================================") //Prints out a spacer line to easily differentiate between network devices.
	}
}

//==========================================================================================================================================================================================================
//	Attributions
//==========================================================================================================================================================================================================
//net - pkg.go.dev
//URL:	https://pkg.go.dev/net

//fmt - pkg.go.dev
//URL:	https://pkg.go.dev/fmt

//Download Visual Studio Code - Mac, Linux, Windows
//URL:	https://code.visualstudio.com/download
//* used to download a GO-supported coding environment.

//Go with Visual Studio Code
//URL: https://code.visualstudio.com/docs/languages/go
//* used to install/get started with GO in Visual Studio Code.

//Managing network devices with Golang using Netrasp
//Author:	Patrick Ogenstad
//URL:	https://networklore.com/hello-netrasp/

//Linux Show / Display Available Network Interfaces
//Author:	Vivek Gite
//URL:	https://www.cyberciti.biz/faq/linux-list-network-interfaces-names-command/

//System network interfaces
//URL:	https://ispycode.com/GO/Network/System-network-interfaces
//	*By far the most useful resource I came across, though in truth it's the last one I came across. This was exactly what I envisioned building from the start, and it helped me sort out my existing syntax issues, add additional formatting, and resolve some pesky errors.

//Extracting network interface type and name in GoLang on CentOS machine
//Contributor:	curiousengineer	https://stackoverflow.com/users/4260303/curiousengineer
//URL:	https://stackoverflow.com/questions/60938206/extracting-network-interface-type-and-name-in-golang-on-centos-machine

//networking - How to get all addresses and masks from local interfaces in go?
//Contributor: 	Everton	https://stackoverflow.com/users/1011695/everton
//URL:	https://stackoverflow.com/questions/23529663/how-to-get-all-addresses-and-masks-from-local-interfaces-in-go

//The Go Playground
//URL:	https://play.golang.org/p/fe-F2k6prlA

//get all IP address from CIDR in golang
//Contributor:	kotakanbe	https://gist.github.com/kotakanbe
//URL:	https://gist.github.com/kotakanbe/d3059af990252ba89a82

//net: re-implement Interfaces and InterfaceAddrs for IPNet, IPv6 on Windows #5395
//Contributor:	tom.walsh@expresshosting.net
//URL: https://github.com/golang/go/issues/5395

//A Tour of GO - Range
//URL: https://tour.golang.org/moretypes/16

//The GO Blog - Using GO Modules
//Author(s): Tyler Bui-Palsulich, Eno Compton
//URL: https://blog.golang.org/using-go-modules

//Package net
//URL: https://golang.org/pkg/net/

//Go by Example: Range
//Author:	Mark McGranaghan
//URL:	https://gobyexample.com/range

//codecademy Learn GO
//URL:	https://www.codecademy.com/learn/learn-go
//* Helpful in introducing me to basic syntax, variables, etc. in GO. Very similar to other programming languages, so I was able to grasp the basics quickly.

//Npcap: WinPcap for Windows 10
//URL:	https://nmap.org/npcap/windows-10.html

//Print details of network interface ( IP address, mac, status etc ) using go language program
//URL:	https://lynxbee.com/print-details-of-network-interface-ip-address-mac-status-etc-using-go-language-program/

//syscall - pkg.go.dev
//URL: https://pkg.go.dev/syscall

//Golang in net package usage (i)
//Author:	User
//URL:	https://topic.alibabacloud.com/a/golang-in-net-package-usage-i_1_38_30925771.html

//GO language Get network card information
//URL:	https://www.programmersought.com/article/24248046952/
//* Another very useful resource that helped me better understand getting network info via GO, along with good examples.

//Interfaces in Go (part I)
//Author: Michał Łowicki https://medium.com/@mlowicki
//URL: https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c

//Golang Examples (Go Examples)
//URL:	https://golang-examples.tumblr.com/post/99458329439/get-local-ip-addresses

//[Solved] $GOPATH/go.mod exists but should not
//URL:	https://programmerah.com/solved-gopath-go-mod-exists-but-should-not-28581/

//Go Modules Reference - The Go Programming Language
//URL:	https://golang.org/ref/mod#go-mod-init
//* used as a reference for creating go.mod files.

//Documentation - The Go Programming Language
//URL:	https://golang.org/doc/#getting-started

//Tutorial: Create a GO module - The Go Programming Language
//URL:	https://golang.org/doc/tutorial/create-module

//How to Write GO Code - The Go Programming Language
//URL:	https://golang.org/doc/code

//src/net/mac.go - The Go Programming Language
//URL:	https://golang.org/src/net/mac.go

//How do I get the local IP address in Go?
//Contributor:	Jerry YY Rain https://stackoverflow.com/users/1337961/jerry-yy-rain
//URL: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go

//Downloads - The Go Programming Language
//URL:	https://golang.org/dl/
// *used to download/install GO

//Interface naming convention Golang
//Contributor:	Dale https://stackoverflow.com/users/1119810/dale
//URL:	https://stackoverflow.com/questions/38842457/interface-naming-convention-golang#:~:text=Effective%20Go%3A%20Interface%20names%3A,the%20function%20names%20they%20capture.

//Go and SSH for network devices
//Author:	Egor Krivosheev https://medium.com/@Vasya4k
//URL:	https://medium.com/@Vasya4k/go-and-ssh-for-network-devices-128937852ccb

//Go how to list all IPs in a network
//Contributor:	Jonathan https://stackoverflow.com/users/3176550/jonathan
//URL: https://stackoverflow.com/questions/60540465/go-how-to-list-all-ips-in-a-network

//What is “_,” (underscore comma) in a Go declaration?
//Contributor:	Kansuler https://stackoverflow.com/users/3656968/kansuler
//URL:	https://stackoverflow.com/questions/27764421/what-is-underscore-comma-in-a-go-declaration

//The GO Programming Language Specification - The GO Programming Language
//URL: https://golang.org/ref/spec

//how can I discover the network device type?
//Contributor:	Richard https://stackoverflow.com/users/99542/richard
//URL: https://stackoverflow.com/questions/47355429/how-can-i-discover-the-network-device-type

//pcap - pkg.go.dev
//URL:	https://pkg.go.dev/github.com/google/gopacket@v1.1.19/pcap

//==========================================================================================================================================================================================================
//
//==========================================================================================================================================================================================================
