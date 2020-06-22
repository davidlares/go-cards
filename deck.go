package main

import (
  "fmt"
  "strings"
  "math/rand"
  "io/ioutil"
  "time"
  "os"
)

// similarities = deck = []string
type deck []string

// functions
func newDeck() deck {
  cards := deck{} // empty
  // suites
  cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
  // values
  cardValues := []string{"Ace", "Two", "Three", "Four"}
  // looping
  for _, suit := range cardSuits {
      for _, value := range cardValues {
        cards = append(cards, value + " of " + suit)
      }
  }
  return cards
}

// (d deck) is the receiver "entry param" - any variable inside of the app of type d can access
func (d deck) print() {
  // calling elements
  for i, card := range d {
    fmt.Println(i, card)
  }
}

func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

// slice to string
func (d deck) toString() string {
  return strings.Join([]string(d),",")
  // ["one", "two"] to "one,two" - byteSlice, and strings.Join
}

// ioutil for writing a file to the HDD
func (d deck) saveToFile(filename string) error {
  // string to byteSlice
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// reading from HDD
func newDeckFromFile(filename string) deck {
  bs, err := ioutil.ReadFile(filename)
  if(err != nil) {
    fmt.Println("Error: ", err) // writing the log
    os.Exit(1) // exit the program
  }
  // byteSlice to String and splitted with the commas (,)
  s := strings.Split(string(bs),",")
  // then, to deck -- if you check - type deck string[]
  return deck(s)
}

func (d deck) shuffe() {
  // seed value
  source := rand.NewSource(time.Now().UnixNano()) // int64 value expected - time is used for unique
  r := rand.New(source)

  for i := range d {
    // newPosition := rand.Intn(len(d) - 1) // pseudoRandom generator - uses the same data seed
    newPosition := r.Intn(len(d) - 1) // new source - r has a Intn method as well
    // ordered assignation
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
