# Simple Example Reading From Config File

This example just prints the string provided.

Precedence is:

1. `--name` or `-n` from the CLI
2. Env Var of `MYCLI_NAME`
3. Config file found `~/.mycli.yaml`

# Example

Provide via CLI

```
$ go run main.go saymyname --name Christian
Christian
```

Provide via config file

```
$ echo "name: Ricardo" > ~/.mycli.yaml
$ go run main.go saymyname 
Ricardo
```

Provide via an env var

```
export MYCLI_NAME=Bob
$ go run main.go saymyname 
Bob
```

Prompt and ask, write to configfile to use for later (unset env var for effect)
```
$ unset MYCLI_NAME 
$ cat ~/.mycli.yaml 
name: Ricardo

$ go run main.go saymyname --ask
Your name: Fred

$ go run main.go saymyname 
Fred
$ go run main.go saymyname 
Fred
$ go run main.go saymyname 
Fred

$ cat ~/.mycli.yaml 
name: Fred
```