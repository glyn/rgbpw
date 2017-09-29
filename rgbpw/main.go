package main

import (
	"os"
	"code.cloudfoundry.org/lager"

	"github.com/glyn/rgbpw/system"
	"fmt"
	"strings"
)


func main() {
	logger := lager.NewLogger("Logs")
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	if len(os.Args) < 2 || len(os.Args) > 4 || os.Args[1] == "help" {
		fmt.Printf("Prints the UAA admin credentials, or just the password if the user name is 'admin', of a specified PCF instance.\n\nUsage: \n  %s <color | hostname> <ops manager userid> <ops manager password>\n", os.Args[0])
		os.Exit(-1)
	}

	hostName := os.Args[1]
	if !strings.Contains(hostName, ".") {
		hostName = fmt.Sprintf("pcf.%s.springapps.io", hostName)
	}

	opsMgrUser := "pivotalcf"
	if len(os.Args) > 2 {
		opsMgrUser = os.Args[2]
	}

	opsMgrPassword := "pivotalcf"
	if len(os.Args) > 3 {
		opsMgrPassword = os.Args[3]
	}

	opsManager, err := system.NewOpsManagerClient(hostName, opsMgrUser, opsMgrPassword, logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create ops manager client: %s\n", err)
		os.Exit(-2)
	}

	adminUser, adminPassword, err := opsManager.GetAdminCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain admin credentials: %s\n", err)
		os.Exit(-3)
	}

	if adminUser != "admin" {
		fmt.Println("admin user: ", adminUser)
	}

	fmt.Println(adminPassword)
	os.Exit(0)
}
