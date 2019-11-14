package main

import (
	"github.com/mysteriumnetwork/vend/cli"
	"github.com/mysteriumnetwork/vend/file"
)

func main() {

	options := cli.ParseOptions()

	if options.Help {
		options.PrintUsage()

	} else {
		cli.UpdateModule()
		json := cli.ReadDownloadJSON()
		deps := file.ParseDownloadJSON(json)

		if options.PkgOnly {
			json = cli.ReadModJSON()
			mod := file.ParseModJSON(json)

			file.CopyPkgDependencies(mod, deps)
		} else {
			file.CopyModuleDependencies(deps, options.ReplaceDeps)
		}
	}
}
