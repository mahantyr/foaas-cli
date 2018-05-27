package fuck

import (
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "strings"
)

func makeRequest(queries ...string) string {

  const baseURL = "http://foaas.com/"

  buildURL := func() string{
    endpoint := baseURL + strings.Join(queries, "/")
    return endpoint
  }

  client := &http.Client{}
  req, err := http.NewRequest("GET", buildURL(), nil)

  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("Accept", "text/plain")
  resp, err := client.Do(req)

  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  return string(body)
}

func GetVersion() string {
  return makeRequest("version")
}

func GetOperations() {
  fmt.Println(makeRequest("operations"))
}

func Anyway(comapany, from string) string {
    return makeRequest("anyway", comapany, from)
}

func Asshole(from string) string {
  return makeRequest("asshole", from)
}

func Awesome(from string) string {
  return makeRequest("awesome", from)
}

func Back(name, from string) string {
  return makeRequest("back", name, from)
}

func Bag(from string) string {
  return makeRequest("bag", from)
}

func Ballmer(name, company, from string) string {
  return makeRequest("ballmer", name, company, from)
}

func Bday(name, from string) string {
  return makeRequest("bday", name, from)
}

func Because(from string) string {
  return makeRequest("because", from)
}

func Blackadder(name, from string) string {
  return makeRequest("blackadder", name, from)
}

func Bm(name, from string) string {
  return makeRequest("bm", name, from)
}


func Bucket(from string) string {
  return makeRequest("bucket", from)
}

func Bus(name, from  string) string {
  return makeRequest("bus", name, from)
}

func Bye(from string) string {
  return makeRequest("bye", from)
}

func Caniuse(tool, from string) string {
  return makeRequest("caniuse", tool, from)
}

func Chainsaw(name, from string) string {
  return makeRequest("chainsaw", name, from)
}

func Cocksplat(name, from string) string {
  return makeRequest("cocksplat", name, from)
}

func Cool(from string) string {
  return makeRequest("cool", from)
}

func Cup(from string) string {
  return makeRequest("cup", from)
}

func Dalton(name, from string) string {
  return makeRequest("dalton", name, from)
}

func Deraadt(name, from string) string {
  return makeRequest("deraadt", name, from)
}

func Diabetes(from string) string {
  return makeRequest("diabetes", from)
}

func Donut(name, from string) string {
  return makeRequest("donut", name, from)
}

func DoSomething(do, something, from string) string {
  return makeRequest("dosomething", do, something, from)
}

func Everyone(from string) string {
  return makeRequest("everyone", from)
}

func Everything(from string) string {
  return makeRequest("everything", from)
}

func Family(from string) string {
  return makeRequest("family", from)
}


func Fascinating(from string) string {
  return makeRequest("fascinating", from)
}


func Flying(from string) string {
  return makeRequest("flying", from)
}

func Fyyff(from string) string {
  return makeRequest("fyyf", from)
}


func Give(from string) string {
  return makeRequest("give", from)
}

func Field(name, from, reference string) string {
  return makeRequest("field", name, from, reference)
}

func Greed(noun, from string) string {
  return makeRequest("greed", noun, from)
}

func Horse(from string) string {
  return makeRequest("horse", from)
}

func Immensity(from string) string {
  return makeRequest("immensity", from)
}


func Ing(name, from string) string {
  return makeRequest("ing", name, from)
}

func Keep(name, from string) string {
  return makeRequest("keep", name, from)
}

func KeepCalm(reaction, from string) string {
  return makeRequest("keepcalm", reaction, from)
}

func King(name, from string) string {
  return makeRequest("king", name, from)
}

func Life(from string) string {
  return makeRequest("life", from)
}

func Linus(name, from string) string {
  return makeRequest("linus", name, from)
}

func Look(name, from string) string {
  return makeRequest("look", name, from)
}

func Looking(from string) string {
  return makeRequest("looking", from)
}

func Madison(name, from string) string {
  return makeRequest("madison", name, from)
}

func Maybe(from string) string {
  return makeRequest("maybe", from)
}

func Me(from string) string {
  return makeRequest("me", from)
}

func Mornin(from string) string {
  return makeRequest("mornin", from)
}

func No(from string) string {
  return makeRequest("no", from)
}

func Nugget(name, from string) string {
  return makeRequest("nugget", name, from)
}

func Off(name, from string) string {
  return makeRequest("off", name, from)
}

func OffWith(behaviour, from string) string {
  return makeRequest("off-with", behaviour, from)
}

func Outside(name, from string) string {
  return makeRequest("outside", name, from)
}

func Particular(thing, from string) string {
  return makeRequest("particular", thing, from)
}

func Pink(from string) string {
  return makeRequest("pink", from)
}

func Problem(name, from string) string {
  return makeRequest("problem", name, from)
}

func Programmer(from string) string {
  return makeRequest("programmer", from)
}

func Pulp(language, from string) string {
  return makeRequest("pulp", language, from)
}

func Question(from string) string {
  return makeRequest("question", from)
}

func Retard(from string) string {
  return makeRequest("retard", from)
}

func Ridiculous(from string) string {
  return makeRequest("ridiculous", from)
}

func Rtfm(from string) string {
  return makeRequest("rtfm", from)
}

func Sake(from string) string {
  return makeRequest("sake", from)
}

func Shit(from string) string {
  return makeRequest("shit", from)
}

func Single(from string) string {
  return makeRequest("single", from)
}

func Thanks(from string) string {
  return makeRequest("thanks", from)
}

func That(from string) string {
  return makeRequest("that", from)
}

func Shakespeare(name, from string) string {
  return makeRequest("shakespeare", name, from)
}

func Think(name, from string) string {
  return makeRequest("think", name, from)
}

func Thinking(name, from string) string {
  return makeRequest("thinking", name, from)
}

func Thumbs(name, from string) string {
  return makeRequest("thumbs", name, from)
}

func Xmas(name, from string) string {
  return makeRequest("xmas", name, from)
}

func Yoda(name, from string) string {
  return makeRequest("yoda", name, from)
}

func You(name, from string) string {
  return makeRequest("you", name, from)
}

func This(from string) string {
  return makeRequest("this", from)
}

func Too(from string) string {
  return makeRequest("too", from)
}

func Tucker(from string) string {
  return makeRequest("tucker", from)
}

func What(from string) string {
  return makeRequest("what", from)
}

func Zayn(from string) string {
  return makeRequest("zayn", from)
}

func Zero(from string) string {
  return makeRequest("zero", from)
}
