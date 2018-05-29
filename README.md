<h1 align="center">
  <img src="https://raw.githubusercontent.com/palash25/foaas-cli/master/assets/logo.png" alt="Logo">
  <h2 align="center">FOaaS CLI</h2>
  <p align="center">
  A simple CLI tool to interact with the FOaaS (Fuck Off as a Service) API made
  in an attempt to learn Go.
  </p>
</h1>

<p align="center">
  <img src="https://raw.githubusercontent.com/palash25/foaas-cli/master/assets/fuck.gif" />
</p>


## Installation
`go get github.com/palash25/foaas-cli`

Make sure that your `GOPATH` and `GOROOT` are set properly.

## Usage
Every command starts with the app name and then the command followed by
arguments and flags (if any)

There are 3 commands in total:
1. `help`: `foaas-cli help`
2. `version`: `foaas-cli version`. This will return the version of the FOaaS API
3. `fuck`: `foaas-cli fuck [argument] [flags]`. The argument is the kind of fuck
    (the message) that you want the API to return and the flags substitute the
    values provided to them in the message returned by the API. For example
    ```
      $ foaas-cli fuck bucket --from Borat

      Gives the following output:

      Please choke on a bucket of cocks. - Borat
    ```
    Here `bucket` is the argument and from is the `flag` with `Borat` as its
    value.

**Like what you see? Then consider giving it a :star: star**
