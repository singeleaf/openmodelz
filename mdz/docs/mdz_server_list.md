## mdz server list

List all servers in the cluster

### Synopsis

List all servers in the cluster

```
mdz server list [flags]
```

### Examples

```
  mdz server list
```

### Options

```
  -h, --help      help for list
  -q, --quiet     Quiet mode - print out only the server names
  -v, --verbose   Verbose mode - print out all server details
```

### Options inherited from parent commands

```
  -a, --agent string                URL of the OpenModelZ agent (MDZ_AGENT) (default http://localhost:8081)
      --debug                       Enable debug logging
  -n, --namespace string            Namespace to use for OpenModelZ inferences (default "default")
  -p, --polling-interval duration   Polling interval (default 3s)
```

### SEE ALSO

* [mdz server](mdz_server.md)	 - Manage the servers

###### Auto generated by spf13/cobra on 19-Jul-2023