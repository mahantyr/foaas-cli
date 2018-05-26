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

func getOperations() {
  fmt.Println(makeRequest("operations"))
}

func anyway(comapany, from string) string {
    return makeRequest("anyway", comapany, from)
}

func asshole(from string) string {
  return makeRequest("asshole", from)
}


func awesome(from string) string {
  return makeRequest("awesome", from)
}


func back(name, from string) string {
  return makeRequest("back", name, from)
}


func bag(from string) string {
  return makeRequest("bag", from)
}


func ballmer(name, company, from string) string {
  return makeRequest("ballmer", name, company, from)
}


func bday(name, from string) string {
  return makeRequest("bday", name, from)
}


func because(from string) string {
  return makeRequest("because", from)
}


func blackadder(name, from string) string {
  return makeRequest("blackadder", name, from)
}


func bm(name, from string) string {
  return makeRequest("bm", name, from)
}


func bucket(from string) string {
  return makeRequest("bucket", from)
}


func bus(name, from  string) string {
  return makeRequest("bus", name, from)
}


func bye(from string) string {
  return makeRequest("bye", from)
}


func caniuse(tool, from string) string {
  return makeRequest("caniuse", tool, from)
}


func chainsaw(name, from string) string {
  return makeRequest("chainsaw", name, from)
}

func cocksplat(name, from string) string {
  return makeRequest("cocksplat", name, from)
}

func cool(from string) string {
  return makeRequest("cool", from)
}

func cup(from string) string {
  return makeRequest("cup", from)
}

func dalton(name, from string) string {
  return makeRequest("dalton", name, from)
}


func deraadt(name, from string) string {
  return makeRequest("deraadt", name, from)
}

func diabetes(from string) string {
  return makeRequest("diabetes", from)
}

func donut(name, from string) string {
  return makeRequest("donut", name, from)
}

func dosomething(do, something, from string) string {
  return makeRequest("dosomething", do, something, from)
}

func everyone(from string) string {
  return makeRequest("everyone", from)
}

func everything(from string) string {
  return makeRequest("everything", from)
}


func family(from string) string {
  return makeRequest("family", from)
}


func fascinating(from string) string {
  return makeRequest("fascinating", from)
}


func flying(from string) string {
  return makeRequest("flying", from)
}


func fyyff(from string) string {
  return makeRequest("fyyf", from)
}


func give(from string) string {
  return makeRequest("give", from)
}


func field(name, from, reference string) string {
  return makeRequest("field", name, from, reference)
}

func greed(noun, from string) string {
  return makeRequest("greed", noun, from)
}

func horse(from string) string {
  return makeRequest("horse", from)
}

func immensity(from string) string {
  return makeRequest("immensity", from)
}


func ing(name, from string) string {
  return makeRequest("ing", name, from)
}


func keep(name, from string) string {
  return makeRequest("keep", name, from)
}


func keepcalm(reaction, from string) string {
  return makeRequest("keepcalm", reaction, from)
}


func king(name, from string) string {
  return makeRequest("king", name, from)
}


func life(from string) string {
  return makeRequest("life", from)
}


func linus(name, from string) string {
  return makeRequest("linus", name, from)
}


func look(name, from string) string {
  return makeRequest("look", name, from)
}

func looking(from string) string {
  return makeRequest("looking", from)
}


func madison(name, from string) string {
  return makeRequest("madison", name, from)
}


func maybe(from string) string {
  return makeRequest("maybe", from)
}


func me(from string) string {
  return makeRequest("me", from)
}


func mornin(from string) string {
  return makeRequest("mornin", from)
}


func no(from string) string {
  return makeRequest("no", from)
}


func nugget(name, from string) string {
  return makeRequest("nugget", name, from)
}


func off(name, from string) string {
  return makeRequest("off", name, from)
}


func off_with(behaviour, from string) string {
  return makeRequest("off-with", behaviour, from)
}


func outside(name, from string) string {
  return makeRequest("outside", name, from)
}


func particular(thing, from string) string {
  return makeRequest("particular", thing, from)
}


func pink(from string) string {
  return makeRequest("pink", from)
}


func problem(name, from string) string {
  return makeRequest("problem", name, from)
}


func programmer(from string) string {
  return makeRequest("programmer", from)
}


func pulp(language, from string) string {
  return makeRequest("pulp", language, from)
}


func question(from string) string {
  return makeRequest("question", from)
}


func retard(from string) string {
  return makeRequest("retard", from)
}


func ridiculous(from string) string {
  return makeRequest("ridiculous", from)
}


func rtfm(from string) string {
  return makeRequest("rtfm", from)
}


func sake(from string) string {
  return makeRequest("sake", from)
}


func shit(from string) string {
  return makeRequest("shit", from)
}


func single(from string) string {
  return makeRequest("single", from)
}


func thanks(from string) string {
  return makeRequest("thanks", from)
}


func that(from string) string {
  return makeRequest("that", from)
}


func shakespeare(name, from string) string {
  return makeRequest("shakespeare", name, from)
}


func think(name, from string) string {
  return makeRequest("think", name, from)
}


func thinking(name, from string) string {
  return makeRequest("thinking", name, from)
}


func thumbs(name, from string) string {
  return makeRequest("thumbs", name, from)
}


func xmas(name, from string) string {
  return makeRequest("xmas", name, from)
}


func yoda(name, from string) string {
  return makeRequest("yoda", name, from)
}


func you(name, from string) string {
  return makeRequest("you", name, from)
}


func this(from string) string {
  return makeRequest("this", from)
}


func too(from string) string {
  return makeRequest("too", from)
}


func tucker(from string) string {
  return makeRequest("tucker", from)
}


func what(from string) string {
  return makeRequest("what", from)
}


func zayn(from string) string {
  return makeRequest("zayn", from)
}


func zero(from string) string {
  return makeRequest("zero", from)
}
