package cmd

import (
  "fmt"
  "github.com/palash25/foaas-cli/fucks"
  "github.com/spf13/cobra"
)

func init() {
  fuckCmd.Flags().StringVarP(&from, "from", "f", "", "Who the fuck is this fuck from?")
  fuckCmd.Flags().StringVarP(&company, "company", "c", "", "")
  fuckCmd.Flags().StringVarP(&name, "name", "n", "", "")
  fuckCmd.Flags().StringVarP(&lang, "lang", "l", "", "")
  fuckCmd.Flags().StringVarP(&ref, "reference", "r", "", "")
  fuckCmd.Flags().StringVarP(&react, "reaction", "", "", "")
  fuckCmd.Flags().StringVarP(&thing, "thing", "t", "", "")
  fuckCmd.Flags().StringVarP(&behaviour, "behaviour", "b", "", "")
  fuckCmd.Flags().StringVarP(&tool, "tool", "", "", "")
  fuckCmd.Flags().StringVarP(&do, "do", "d", "", "")
  fuckCmd.Flags().StringVarP(&something, "something", "s", "", "")
  fuckCmd.Flags().StringVarP(&noun, "noun", "", "", "")
}

var from, name, noun, lang, company, ref, react, thing, behaviour, tool, do, something string

var fuckCmd = &cobra.Command{
  Use:   "fuck",
  Short: "Request various fucks from FOaaS",
  Run: func(cmd *cobra.Command, args []string) {
         switch args[0] {
         case "asshole":
                fmt.Println(fucks.Asshole(from))
         case "anyway":
                fmt.Println(fucks.Anyway(company, from))
         case "awesome":
                fmt.Println(fucks.Awesome(from))
         case "back":
                fmt.Println(fucks.Back(name, from))
         case "bag":
                fmt.Println(fucks.Bag(from))
         case "ballmer":
                fmt.Println(fucks.Ballmer(name, company, from))
         case "bday":
                fmt.Println(fucks.Bday(name, from))
         case "because":
                fmt.Println(fucks.Because(from))
         case "blackadder":
                fmt.Println(fucks.Blackadder(name, from))
         case "bm":
                fmt.Println(fucks.Bm(name, from))
         case "bucket":
                fmt.Println(fucks.Bucket(from))
         case "bus":
                fmt.Println(fucks.Bus(name, from))
         case "bye":
                fmt.Println(fucks.Bye(from))
         case "caniuse":
                fmt.Println(fucks.Caniuse(name, from))
         case "chainsaw":
                fmt.Println(fucks.Chainsaw(name, from))
         case "cocksplat":
                fmt.Println(fucks.Cocksplat(name, from))
         case "cool":
                fmt.Println(fucks.Cool(from))
         case "cup":
                fmt.Println(fucks.Cup(from))
         case "dalton":
                fmt.Println(fucks.Dalton(name, from))
         case "deraadt":
                fmt.Println(fucks.Deraadt(name, from))
         case "diabetes":
                fmt.Println(fucks.Diabetes(from))
         case "donut":
                fmt.Println(fucks.Donut(name, from))
         case "dosomething":
                fmt.Println(fucks.DoSomething(do, something, from))
         case "everyone":
                fmt.Println(fucks.Everyone(from))
         case "everything":
                fmt.Println(fucks.Everything(from))
         case "family":
                fmt.Println(fucks.Family(from))
         case "fascinating":
                fmt.Println(fucks.Fascinating(from))
         case "flying":
                fmt.Println(fucks.Flying(from))
         case "fyyff":
                fmt.Println(fucks.Fyyff(from))
         case "give":
                fmt.Println(fucks.Give(from))
         case "field":
                fmt.Println(fucks.Field(name, from, ref))
         case "greed":
                fmt.Println(fucks.Greed(noun, from))
         case "horse":
                fmt.Println(fucks.Horse(from))
         case "immensity":
                fmt.Println(fucks.Immensity(from))
         case "ing":
                fmt.Println(fucks.Ing(name, from))
         case "keep":
                fmt.Println(fucks.Keep(name, from))
         case "keepcalm":
                fmt.Println(fucks.KeepCalm(react, from))
         case "king":
                fmt.Println(fucks.King(name, from))
         case "life":
                fmt.Println(fucks.Life(from))
         case "linus":
                fmt.Println(fucks.Linus(name, from))
         case "look":
                fmt.Println(fucks.Look(name, from))
         case "looking":
                fmt.Println(fucks.Looking(from))
         case "madison":
                fmt.Println(fucks.Madison(name, from))
         case "maybe":
                fmt.Println(fucks.Maybe(from))
         case "me":
                fmt.Println(fucks.Me(from))
         case "mornin":
                fmt.Println(fucks.Mornin(from))
         case "no":
                fmt.Println(fucks.No(from))
         case "nugget":
                fmt.Println(fucks.Nugget(name, from))
         case "off":
                fmt.Println(fucks.Off(name, from))
         case "offwith":
                fmt.Println(fucks.OffWith(behaviour, from))

         default:
                panic("Invalid argument")
         }
  },
}
