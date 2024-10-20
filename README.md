# SOLID CLI

SOLID CLI is a command-line interface for managing SOLID accounts, files, and Pods.

## Usage

```sh
solid-cli [command]
```

## Available Commands

- `account` - Manage SOLID accounts
- `files` - Manage account files
- `help` - Help about any command
- `pod` - Manage SOLID Pods
- `version` - Print the version number of SOLID CLI

### Account Commands

```sh
solid-cli account [command]
```

#### Available Commands

- `ls` - List all SOLID Pods
- `rm` - Delete a SOLID account

##### Account `rm` Command

```sh
solid-cli account rm [flags]
```

#### Flags

- `-e, --email string` - Email
- `-h, --help` - Help for `rm`
- `-i, --id string` - Account ID

### Pod Commands

```sh
solid-cli pod [command]
```

#### Available Commands

- `ls` - List all SOLID Pods

### Files Commands

```sh
solid-cli files [command]
```

#### Available Commands

- `ls` - List all SOLID Files

##### Files `ls` Command

```sh
solid-cli files ls [flags]
```

#### Flags

- `-e, --email string` - Email
- `-h, --help` - Help for `rm`

## Getting Help

For more information on a specific command, you can use the `help` command:

```sh
solid-cli help [command]
```

## Version

To print the version number of SOLID CLI, use:

```sh
solid-cli version
```
