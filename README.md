# geasy

`geasy` is a small wrapper cli tool that use the gcp api's to make it little more usable especially for more than one project. 

## Requirements: 

- go
- GOPATH env must be set

## Install

```bash
make install # this will build the cli and copy it into go bin path
```

## Authentication

In this example, Google Application Default Credentials are used for authentication.

For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.

## Autocompletion

As a developer/user you can enable autocompletion via a registered command 

```bash
geasy completion zsh > "${fpath[1]}/_geasy" #zsh example
```

more infos here you can find here https://github.com/spf13/cobra/blob/master/shell_completions.md#creating-your-own-completion-command