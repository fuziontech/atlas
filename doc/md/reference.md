---
title: CLI Reference
id: cli-reference
slug: cli-reference
---
## Introduction

This document serves as reference documentation for all available commands in the Atlas CLI.
Similar information can be obtained by running any atlas command with the `-h` or `--help`
flags.

For a more detailed introduction to the CLI capabilities, head over to the
[Getting Started](/getting-started/) page.

## Distributed Binaries

The binaries distributed in official releases are released under the  [Ariga End User License](https://ariga.io/legal/atlas/eula).
If you would like to build Atlas from source follow the instructions [here](https://atlasgo.io/cli-reference#building-from-source).

### Building from Source

If you would like to build Atlas from source without the UI code run:
```shell
go get ariga.io/atlas/cmd/atlas
```

## atlas env

Print atlas environment variables.

#### Usage
```
atlas env
```

#### Details
'atlas env' prints atlas environment information.

Every set environment param will be printed in the form of NAME=VALUE.

List of supported environment parameters:
* ATLAS_NO_UPDATE_NOTIFIER: On any command, the CLI will check for new releases using the GitHub API.
  This check will happen at most once every 24 hours. To cancel this behavior, set the environment 
  variable "ATLAS_NO_UPDATE_NOTIFIER".


## atlas license

Display license information

#### Usage
```
atlas license
```


## atlas migrate

Manage versioned migration files

#### Usage
```
atlas migrate
```

#### Details
'atlas migrate' wraps several sub-commands for migration management.

#### Flags
```
      --dir string           select migration directory using URL format (default "file://migrations")
      --env string           set which env from the project file to use
      --var stringToString   input variables (default [])

```


### atlas migrate apply

Applies pending migration files on the connected database.

#### Usage
```
atlas migrate apply [flags] [count]
```

#### Details
'atlas migrate apply' reads the migration state of the connected database and computes what migrations are pending.
It then attempts to apply the pending migration files in the correct order onto the database. 
The first argument denotes the maximum number of migration files to apply.
As a safety measure 'atlas migrate apply' will abort with an error, if:
  - the migration directory is not in sync with the 'atlas.sum' file
  - the migration and database history do not match each other

If run with the "--dry-run" flag, atlas will not execute any SQL.

#### Example

```
  atlas migrate apply -u mysql://user:pass@localhost:3306/dbname
  atlas migrate apply --dir file:///path/to/migration/directory --url mysql://user:pass@localhost:3306/dbname 1
  atlas migrate apply --env dev 1
  atlas migrate apply --dry-run --env dev 1
```
#### Flags
```
      --log string                log format to use (default "tty")
      --revisions-schema string   schema name where the revisions table resides
      --dry-run                   do not actually execute any SQL but show it on screen
      --from string               calculate pending files from the given version (including it)
      --baseline string           start the first migration after the given baseline version
      --tx-mode string            set transaction mode [none, file, all] (default "file")
      --allow-dirty               allow start working on a non-clean database
  -u, --url string                [driver://username:password@address/dbname?param=value] select a database using the URL format

```


### atlas migrate diff

Compute the diff between the migration directory and a desired state and create a new migration file.

#### Usage
```
atlas migrate diff [flags] [name]
```

#### Details
'atlas migrate diff' uses the dev-database to re-run all migration files in the migration directory, compares
it to a given desired state and create a new migration file containing SQL statements to migrate the migration
directory state to the desired schema. The desired state can be another connected database or an HCL file.

#### Example

```
  atlas migrate diff --dev-url mysql://user:pass@localhost:3306/dev --to file://atlas.hcl
  atlas migrate diff --dev-url mysql://user:pass@localhost:3306/dev --to file://atlas.hcl add_users_table
  atlas migrate diff --dev-url mysql://user:pass@localhost:3306/dev --to mysql://user:pass@localhost:3306/dbname
  atlas migrate diff --env dev
```
#### Flags
```
      --dev-url string     [driver://username:password@address/dbname?param=value] select a database using the URL format
      --to strings         [driver://username:password@address/dbname?param=value ...] select a desired state using the URL format
      --qualifier string   qualify tables with custom qualifier when working on a single schema

```


### atlas migrate hash

Hash (re-)creates an integrity hash file for the migration directory.

#### Usage
```
atlas migrate hash
```

#### Details
'atlas migrate hash' computes the integrity hash sum of the migration directory and stores it in the atlas.sum file.
This command should be used whenever a manual change in the migration directory was made.

#### Example

```
  atlas migrate hash --force
```

### atlas migrate lint

Run analysis on the migration directory

#### Usage
```
atlas migrate lint [flags]
```

#### Example

```
  atlas migrate lint --env dev
  atlas migrate lint --dir file:///path/to/migration/directory --dev-url mysql://root:pass@localhost:3306 --latest 1
  atlas migrate lint --dir file:///path/to/migration/directory --dev-url mysql://root:pass@localhost:3306 --git-base master
  atlas migrate lint --dir file:///path/to/migration/directory --dev-url mysql://root:pass@localhost:3306 --log '{{ json .Files }}'
```
#### Flags
```
      --dev-url string    [driver://username:password@address/dbname?param=value] select a database using the URL format
      --git-base string   run analysis against the base Git branch
      --git-dir string    path to the repository working directory (default ".")
      --latest uint       run analysis on the latest N migration files
      --log string        custom logging using a Go template

```


### atlas migrate new

Creates a new empty migration file in the migration directory.

#### Usage
```
atlas migrate new [name]
```

#### Details
'atlas migrate new' creates a new migration according to the configured formatter without any statements in it.

#### Example

```
  atlas migrate new my-new-migration
```

### atlas migrate status

Get information about the current migration status.

#### Usage
```
atlas migrate status [flags]
```

#### Details
'atlas migrate status' reports information about the current status of a connected database compared to the migration directory.

#### Example

```
  atlas migrate status --url mysql://user:pass@localhost:3306/
  atlas migrate status --url mysql://user:pass@localhost:3306/ --dir file:///path/to/migration/directory
```
#### Flags
```
      --revisions-schema string   schema name where the revisions table resides
  -u, --url string                [driver://username:password@address/dbname?param=value] select a database using the URL format

```


### atlas migrate validate

Validates the migration directories checksum and SQL statements.

#### Usage
```
atlas migrate validate [flags]
```

#### Details
'atlas migrate validate' computes the integrity hash sum of the migration directory and compares it to the
atlas.sum file. If there is a mismatch it will be reported. If the --dev-url flag is given, the migration
files are executed on the connected database in order to validate SQL semantics.

#### Example

```
  atlas migrate validate
  atlas migrate validate --dir file:///path/to/migration/directory
  atlas migrate validate --dir file:///path/to/migration/directory --dev-url mysql://user:pass@localhost:3306/dev
  atlas migrate validate --env dev
```
#### Flags
```
      --dev-url string   [driver://username:password@address/dbname?param=value] select a database using the URL format

```


## atlas schema

Work with atlas schemas.

#### Usage
```
atlas schema
```

#### Details
The `atlas schema` command groups subcommands for working with Atlas schemas.

#### Flags
```
      --env string           set which env from the project file to use
      --var stringToString   input variables (default [])

```


### atlas schema apply

Apply an atlas schema to a target database.

#### Usage
```
atlas schema apply [flags]
```

#### Details
'atlas schema apply' plans and executes a database migration to bring a given
database to the state described in the provided Atlas schema. Before running the
migration, Atlas will print the migration plan and prompt the user for approval.

The schema is provided by one or more paths (to a file or directory) using the "-f" flag:
  atlas schema apply -u URL -f file1.hcl -f file2.hcl
  atlas schema apply -u URL -f schema/ -f override.hcl

As a convenience, schemas may also be provided via an environment definition in
the project file (see: https://atlasgo.io/cli/projects).

If run with the "--dry-run" flag, atlas will exit after printing out the planned
migration.

#### Example

```
  atlas schema apply -u "mysql://user:pass@localhost/dbname" -f atlas.hcl
  atlas schema apply -u "mysql://localhost" -f schema.hcl --schema prod --schema staging
  atlas schema apply -u "mysql://user:pass@localhost:3306/dbname" -f schema.hcl --dry-run
  atlas schema apply -u "mariadb://user:pass@localhost:3306/dbname" -f schema.hcl
  atlas schema apply --url "postgres://user:pass@host:port/dbname?sslmode=disable" -f schema.hcl
  atlas schema apply -u "sqlite://file:ex1.db?_fk=1" -f schema.hcl
```
#### Flags
```
  -f, --file strings      [paths...] file or directory containing the HCL files
  -u, --url string        URL to the database using the format:
                          [driver://username:password@address/dbname?param=value]
      --exclude strings   List of glob patterns used to filter resources from applying.
  -s, --schema strings    Set schema names.
      --dev-url string    URL for the dev database. Used to validate schemas and calculate diffs
                          before running migration.
      --dry-run           Dry-run. Print SQL plan without prompting for execution.
      --auto-approve      Auto approve. Apply the schema changes without prompting for approval.

```


### atlas schema clean

Removes all objects from the connected database.

#### Usage
```
atlas schema clean [flags]
```

#### Details
'atlas schema clean' drops all objects in the connected database and leaves it in an empty state.
As a safety feature, 'atlas schema clean' will ask for confirmation before attempting to execute any SQL.

#### Example

```
  atlas schema clean -u mysql://user:pass@localhost:3306/dbname
  atlas schema clean -u mysql://user:pass@localhost:3306/
```
#### Flags
```
      --auto-approve   Auto approve. Apply the schema changes without prompting for approval.
  -u, --url string     URL to the database using the format:
                       [driver://username:password@address/dbname?param=value]

```


### atlas schema diff

Calculate and print the diff between two schemas.

#### Usage
```
atlas schema diff [flags]
```

#### Details
'atlas schema diff' connects to two given databases, inspects
them, calculates the difference in their schemas, and prints a plan of
SQL statements to migrate the "from" database to the schema of the "to" database.

#### Flags
```
      --from string   [driver://username:password@protocol(address)/dbname?param=value] select a database using the URL format
      --to string     [driver://username:password@protocol(address)/dbname?param=value] select a database using the URL format

```


### atlas schema fmt

Formats Atlas HCL files

#### Usage
```
atlas schema fmt [path ...]
```

#### Details
'atlas schema fmt' formats all ".hcl" files under the given path using
canonical HCL layout style as defined by the github.com/hashicorp/hcl/v2/hclwrite package.
Unless stated otherwise, the fmt command will use the current directory.

After running, the command will print the names of the files it has formatted. If all
files in the directory are formatted, no input will be printed out.



### atlas schema inspect

Inspect a database and print its schema in Atlas DDL syntax.

#### Usage
```
atlas schema inspect [flags]
```

#### Details
'atlas schema inspect' connects to the given database and inspects its schema.
It then prints to the screen the schema of that database in Atlas DDL syntax. This output can be
saved to a file, commonly by redirecting the output to a file named with a ".hcl" suffix:

  atlas schema inspect -u "mysql://user:pass@localhost:3306/dbname" > schema.hcl

This file can then be edited and used with the `atlas schema apply` command to plan
and execute schema migrations against the given database. In cases where users wish to inspect
all multiple schemas in a given database (for instance a MySQL server may contain multiple named
databases), omit the relevant part from the url, e.g. "mysql://user:pass@localhost:3306/".
To select specific schemas from the databases, users may use the "--schema" (or "-s" shorthand)
flag.
	

#### Example

```
  atlas schema inspect -u "mysql://user:pass@localhost:3306/dbname"
  atlas schema inspect -u "mariadb://user:pass@localhost:3306/" --schema=schemaA,schemaB -s schemaC
  atlas schema inspect --url "postgres://user:pass@host:port/dbname?sslmode=disable"
  atlas schema inspect -u "sqlite://file:ex1.db?_fk=1"
```
#### Flags
```
      --exclude strings   List of glob patterns used to filter resources from inspection
  -s, --schema strings    Set schema name
  -u, --url string        [driver://username:password@protocol(address)/dbname?param=value] select a database using the URL format

```


## atlas version

Prints this Atlas CLI version information.

#### Usage
```
atlas version
```



