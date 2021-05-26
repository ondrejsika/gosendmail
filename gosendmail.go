// gosendmail
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/gosendmail
//

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ondrejsika/gosendmail/lib"
)

func main() {
	from := flag.String("from", "", "")
	to := flag.String("to", "", "")
	smtpHost := flag.String("smtp-host", "", "")
	smtpPort := flag.String("smtp-port", "587", "optional")
	password := flag.String("password", "", "")
	subject := flag.String("subject", "", "")
	message := flag.String("message", "", "")

	flag.Parse()

	if *from == "" {
		fmt.Fprintf(os.Stderr, "-from is not defined\n")
		os.Exit(1)
	}
	if *to == "" {
		fmt.Fprintf(os.Stderr, "-to is not defined\n")
		os.Exit(1)
	}
	if *smtpHost == "" {
		fmt.Fprintf(os.Stderr, "-smtp-host is not defined\n")
		os.Exit(1)
	}
	if *smtpPort == "" {
		fmt.Fprintf(os.Stderr, "-smtp-port is not defined\n")
		os.Exit(1)
	}
	if *subject == "" {
		fmt.Fprintf(os.Stderr, "-subject is not defined\n")
		os.Exit(1)
	}
	if *message == "" {
		fmt.Fprintf(os.Stderr, "-message is not defined\n")
		os.Exit(1)
	}

	err := lib.GoSendMail(*smtpHost, *smtpPort, *from, *password, *to, *subject, *message)
	if err != nil {
		panic(err)
	}
}
