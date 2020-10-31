# ec2meta

![build](https://github.com/invit/ec2meta/workflows/build/badge.svg)

Simple CLI to query AWS EC2 instance metadata.

## Usage

``` 
$ ec2meta
Display EC2 instance metadata

Usage:
  ec2meta [command]

Available Commands:
  get         Returns arbitrary metadata by path
  help        Help about any command
  version     Print the version number of ec2meta

Flags:
  -h, --help   help for ec2meta

Use "ec2meta [command] --help" for more information about a command.
```

### Query

``` 
$ ec2meta help get
Returns arbitrary metadata by path

Usage:
  ec2meta get <path> [flags]

Flags:
  -h, --help   help for get
```

`<path>` can either point to a directory path, like

``` 
$ ec2meta get /
ami-id
ami-launch-index
ami-manifest-path
block-device-mapping/
events/
hibernation/
hostname
iam/
identity-credentials/
instance-action
instance-id
instance-life-cycle
instance-type
local-hostname
local-ipv4
mac
metrics/
network/
placement/
profile
public-hostname
public-ipv4
reservation-id
security-groups
services/
```

a single element

```
$ ec2meta get /hostname
ip-10-0-0-248.eu-central-1.compute.internal
```

or reference either all `[*]` elements or a specific one (e.g. `[0]`) in a list:

```
$ ec2meta get /network/interfaces/macs/[0]/vpc-id
vpc-XXXXXXXX
```
