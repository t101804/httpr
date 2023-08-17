package runner

import (
	"github.com/projectdiscovery/gologger"
)

const banner = `
    __    __  __            
   / /_  / /_/ /_____  _____
  / __ \/ __/ __/ __ \/ ___/
 / / / / /_/ /_/ /_/ / /    
/_/ /_/\__/\__/ .___/_/     
             /_/           
`

// Version is the current version of httpx
const version = `v1.0.0`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\trepcyber.com\n\n")
}
