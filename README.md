# testport

Small utility to test ports, no matter what OS you're using.

# Installation

```
go install github.com/gkawamoto/testport@latest
```

# Usage

```bash
$ testport -help
USAGE: testport host port [-revert] [-timeout milliseconds; default=500] [-help]
```

# Examples

```bash
$ testport localhost 22 && echo "ssh server running on localhost"
```

```bash
$ testport -revert 192.168.1.2 3389 && echo "remote desktop not available at 192.168.1.2"
```

```bash
$ # slow network, lets leave some breathing room
$ testport -timeout 2000 201.48.33.80 3389 && echo "remote desktop not available at 192.168.1.2"
```
