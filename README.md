# weather
This is a demo CLI that uses the gopackages/app-cli framework. This particular command has no subcommands, so it
only requires an option and an action routine to call when the option is processed.

The CLI help output looks like this:

  (C) Copyright Tom Cole 2020

  Usage:
     weather [options] [command]    view weather for a given location, 1.0-0

  Commands:
    help                           Display help text            
    profile                        Manage the default profile   
  
  Options:
    --debug, -d                    Are we debugging? [CLI_DEBUG]                                   
    --help, -h                     Show this help text                                             
    --location <list>              The location (city, state) for which the weather is displayed   
    --output-format <string>       Specify text or json output format [CLI_OUTPUT_FORMAT]          
    --profile, -p <string>         Name of profile to use [CLI_PROFILE]                            
    --quiet, -q                    If specified, suppress extra messaging [CLI_QUIET]              
    --version, -v                  Show version number of command line tool                        

Note that the actual command grammar only specifies the --location option, which is used to pass a city and state/country
designation. The other options are added automatically by the app-cli framework, as are the help and profile subcommands.

This CLI creates a profile, and uses it to store the REST API key and the last city,state designation used which becomes
the default if a new `--location` is not given on a subsequent command line.

The `main.go` module only creates the application context and runs the application. The `weather.go` file contains both
the grammar definition and the single action routine that uses OpenWeather.com to access a weather forcast.
