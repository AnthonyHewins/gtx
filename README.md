# GTX

Very lightweight golang contexts for your project's configuration

## What does it do

When you have a bigger project that has lots of configuration variables and environments like this:

```go
type Database struct {
    Host string `yaml:"host"`
    Port uint16 `yaml:"port"`
    User string `yaml:"user"`
}

dev := Database{"dev.host.com", 5432, "dev-user"}
stage := Database{"stage.host.com", 5432, "stage-user"}
prod := Database{"prod.host.com", 5432, "prod"}
```

It can be annoying to set those vars and edit config files back and forth in .env files or yaml:

```yaml
# this works when i'm using dev, but then I have to switch it to stage/prod
host: dev.host.com
port: 5432
user: dev-user
```

The goal is with `gtx` to just make all configs for each env and then you can pick the one
you want to use, and load your config into it

```shell
gtx create repo_name
gtx edit repo_name dev # create each one
gtx edit repo_name stage
gtx edit repo_name prod
gtx select repo_name prod # now it's using prod's env
```

Then in your code:

```go
var configObj config
if err := gtx.ReadInto(&configObj);err!=nil{
    return err
}
```

That's it. Makes it easier

## Usage

```shell
usage: gtx COMMANDS

GTX helps create contexts for Go applications to save config
for different environments

Commands:
         ls     List information about contexts
     create     Create a new context
     select     Select current active context
       edit     Edit a context in a repo
       help     Display help text
         rm     Delete a context
        cat     Cat out the contents of a context
```