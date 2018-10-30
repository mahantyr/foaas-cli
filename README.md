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

### Arguments

Here is a list of various arguments along with their flags generated using
go doc. The function names are the arguments (should be in small letters when
typed on the terminal) and the parameters passed to them are the flags. So
if you want to give a fuck, say "Anyway" then type
`foaas-cli fuck anyway --company DoofenshmirtzEvilInc --name Ferb`

```
func Anyway(comapany, from string) string
func Asshole(from string) string
func Awesome(from string) string
func Back(name, from string) string
func Bag(from string) string
func Ballmer(name, company, from string) string
func Bday(name, from string) string
func Because(from string) string
func Blackadder(name, from string) string
func Bm(name, from string) string
func Bucket(from string) string
func Bus(name, from string) string
func Bye(from string) string
func Caniuse(tool, from string) string
func Chainsaw(name, from string) string
func Cocksplat(name, from string) string
func Cool(from string) string
func Cup(from string) string
func Dalton(name, from string) string
func Deraadt(name, from string) string
func Diabetes(from string) string
func DoSomething(do, something, from string) string
func Donut(name, from string) string
func Everyone(from string) string
func Everything(from string) string
func Family(from string) string
func Fascinating(from string) string
func Field(name, from, reference string) string
func Flying(from string) string
func Fyyff(from string) string
func GetOperations() string
func GetVersion() string
func Give(from string) string
func Greed(noun, from string) string
func Horse(from string) string
func Immensity(from string) string
func Ing(name, from string) string
func Keep(name, from string) string
func KeepCalm(reaction, from string) string
func King(name, from string) string
func Life(from string) string
func Linus(name, from string) string
func Look(name, from string) string
func Looking(from string) string
func Madison(name, from string) string
func Maybe(from string) string
func Me(from string) string
func Mornin(from string) string
func No(from string) string
func Nugget(name, from string) string
func Off(name, from string) string
func OffWith(behaviour, from string) string
func Outside(name, from string) string
func Particular(thing, from string) string
func Pink(from string) string
func Problem(name, from string) string
func Programmer(from string) string
func Pulp(language, from string) string
func Question(from string) string
func Retard(from string) string
func Ridiculous(from string) string
func Rtfm(from string) string
func Sake(from string) string
func Shakespeare(name, from string) string
func Shit(from string) string
func Single(from string) string
func Thanks(from string) string
func That(from string) string
func Think(name, from string) string
func Thinking(name, from string) string
func This(from string) string
func Thumbs(name, from string) string
func Too(from string) string
func Tucker(from string) string
func What(from string) string
func Xmas(name, from string) string
func Yoda(name, from string) string
func You(name, from string) string
func Zayn(from string) string
```

**Like what you see? Then consider giving it a :star: star**
