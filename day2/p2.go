package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, _ := os.ReadFile("input.txt")
  res :=0
  for _, x := range strings.Split(string(file),"\n") {
    parts := strings.Split(x,":")
    if len(parts) <= 1 {
      break
    }
    _, rest := parts[0],parts[1]
    rest = strings.ReplaceAll(rest,";",",")

    m := make(map[string]int)
    m["red"] = 0
    m["green"] = 0
    m["blue"] = 0

    for _, part := range strings.Split(rest,","){
      z := strings.Split(strings.TrimSpace(part)," ")
      num,color := z[0],z[1]
      val,_ := strconv.Atoi(num)
      if val > m[color]{
        m[color] = val
      }

    }
    res += m["red"] * m["blue"] * m["green"]
  }
  fmt.Println(res)
}
