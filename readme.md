# Probe

Given a list of urls, Probe will determine if they are online or offline. 

```sh
$ ./probe -help

NAME:
   Probe - A new cli application

USAGE:
   probe [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output value      Where do you want to results written to? If left blank they will only be printed.
   --user-agent value  Custom user agent. Default 'probe/{version}'. (default: "probe/0.0.0")
   --filtered          When turned off all domains checked will be in the output/logs. If turned on only the domains that return with a successful status. (default: true)
   --verbose           Turns on verbose looking. (default: false)
   --help, -h          show help (default: false)
   --version, -v       print the version (default: false)
```

## Example usage

```sh
cat domains.txt | ./probe --filtered=false --verbose=true --output=res.txt --user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
```

## Todos

- [] Save html
- [] Follow redirects 