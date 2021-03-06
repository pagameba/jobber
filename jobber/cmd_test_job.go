package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/dshearer/jobber/ipc"
)

func doTestCmd(args []string) int {
	// parse flags
	flagSet := flag.NewFlagSet(TestCmdStr, flag.ExitOnError)
	flagSet.Usage = subcmdUsage(TestCmdStr, "JOB", flagSet)
	var help_p *bool = flagSet.Bool("h", false, "help")
	//	var jobUser_p *string = flagSet.String("u", user.Username, "user")
	flagSet.Parse(args)

	if *help_p {
		flagSet.Usage()
		return 0
	}

	// get job to test
	if len(flagSet.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a job.\n")
		return 1
	}
	var job string = flagSet.Args()[0]

	// get current user
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintf(
			os.Stderr, "Failed to get current user: %v\n", err,
		)
		return 1
	}

	// send command
	var resp ipc.TestCmdResp
	fmt.Printf("Running job \"%v\"...\n", job)
	err = CallDaemon(
		"IpcService.Test",
		ipc.TestCmd{Job: job},
		&resp,
		usr,
		nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	// handle response
	fmt.Printf("%v\n", resp.Result)
	return 0
}
