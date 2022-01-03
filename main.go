package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"github.com/jessevdk/go-flags"
)

type PcapServerCLI struct {
}

func main() {
	plugin.Start(new(PcapServerCLI))
}

func (cli *PcapServerCLI) Run(cliConnection plugin.CliConnection, args []string) {
	// Initialize flags
	type positional struct {
		AppName string `positional-arg-name:"app" description:"The app to capture." required:"true"`
	}

	var opts struct {
		File       string     `short:"o" long:"file" description:"The output file. Written in binary pcap format." required:"true"`
		Filter     string     `short:"f" long:"filter" description:"Allows to provide a filter expression in pcap filter format." required:"false"`
		Positional positional `positional-args:"true" required:"true"`
	}

	args, err := flags.ParseArgs(&opts, args[1:])

	if err != nil {
		return
	}

	loggedIn, err := cliConnection.IsLoggedIn()

	if !loggedIn || err != nil {
		fmt.Println("Please log in first.")
		return
	}

	fmt.Println("app " + opts.Positional.AppName)
	fmt.Println("file " + opts.File)
	fmt.Println("filter " + opts.Filter)

}

func (cli *PcapServerCLI) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "PcapServerCLI",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "pcap",
				Alias:    "tcpdump",
				HelpText: "Pcap captures network traffic of your apps. To obtain more information use --help",
				UsageDetails: plugin.Usage{
					Usage: "pcap - stream pcap data from your app to disk\n   cf pcap <app> --file <file.pcap> [--filter <expression>]",
					Options: map[string]string{
						"file":   "The output file. Written in binary pcap format.",
						"filter": "Allows to provide a filter expression in pcap filter format. See https://linux.die.net/man/7/pcap-filter",
					},
				},
			},
		},
	}
}
