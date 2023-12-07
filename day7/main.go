package main;

import (
  "fmt"
  "os"
  "strings"
  "cmp"
  "strconv"
  "slices"
)

type hand struct {
  cards string
  val int
}

func main() {
  file, _ := os.ReadFile("input.txt")

  var arr []hand;
  for _, x := range strings.Split(string(file), "\n") {
    row := strings.Split(x," ")
    if len(row) == 2{

      val ,_ := strconv.ParseInt(row[1],10,32)
      arr = append(arr, hand{row[0],int(val)})
    }
  }


  slices.SortFunc(arr,
    func(a, b hand) int {
      ra,rb:=get_rank(a.cards,false),get_rank(b.cards,false)
      if ra == rb{
        return compareHands(a.cards,b.cards)
      }
      return cmp.Compare(ra,rb)
    })

  sum :=0
  for i,x := range arr {
    fmt.Printf("%s %d\n",x.cards,x.val)
    sum+= (i+1) * x.val

  }
  fmt.Println(sum)

}

func compareHands(hand1, hand2 string) int {
  ranks := "AKQT98765432J"


  for i := 0; i < 5; i++ {
    card1Rank := strings.Index(ranks, string(hand1[i]))
    card2Rank := strings.Index(ranks, string(hand2[i]))


    if card1Rank < card2Rank {
      return 1
    } else if card1Rank > card2Rank {
      return -1
    }
  }

  return 0
}

func get_rank(s string, incl_joker bool) int {
  m := make(map[rune]int);

  jcount :=0;
  for _,chr := range s {
    if !incl_joker && chr != 'J'{
      m[chr] += 1
    }else{
      jcount +=1
    }
  }

  counts := make([]int,0,len(m))

  for _, value := range m {
    counts = append(counts, value)
  }

  slices.Sort(counts)

  fmt.Println(s)
  if len(counts) == 0 {
    return 6
  }
  counts[len(counts)-1] +=  jcount

  switch fmt.Sprint(counts) {
  case "[5]":
    return 6
  case "[1 4]":
    return 5
  case "[2 3]":
    return 4
  case "[1 1 3]":
    return 3
  case "[1 2 2]":
    return 2
  case "[1 1 1 2]":
    return 1
  case "[1 1 1 1 1]":
    return 0
  default:
    return -999
}
}
