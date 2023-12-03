package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "unicode"
  "strconv"
)


func get_entire_number(i,j int, grid []string) string {
  num := ""
  if ! unicode.IsDigit(rune(grid[i][j])) {
    return num
  }
  var cur = grid[i][j];
  for  jj := j ; jj < len(grid[i]) && unicode.IsDigit(rune(grid[i][jj])) ;jj++{
    if jj < len(grid[i]){

      cur = grid[i][jj]
      num = num + string(cur)
    }
  }

  for  jj := j-1 ; jj >=0 && unicode.IsDigit(rune(grid[i][jj])) ;jj--{
    cur = grid[i][jj]
    num = string(cur) + num
  }
  return num
}

func get_neighboring_numbers(x,y int, grid []string) []string {
  var numbers []string 

  for dy := -1; dy <= 1; dy++ {
    for dx := -1; dx <= 1; dx++ {
      if dy == 0 && dx == 0 {
        continue
      }
      num := get_entire_number(x+dx,y+dy,grid) 
      if num != "" {
        numbers = append(numbers,num)
      }

    }
  }
  return numbers
}


func main(){
  content, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  input := string(content)

  lines := strings.Split(input, "\n")

  res :=0

  for i, line := range lines {
    for j, char := range line {
      if char == '*' {
        l := get_neighboring_numbers(i,j,lines)
        uniqueElements := make(map[string]bool)
        result := []string{}

        for _, num := range l {
          if !uniqueElements[num] {
            uniqueElements[num] = true
            result = append(result, num)
          }
        }
        if len(result) == 2 {
          fmt.Println(l)
          n1, _ := strconv.Atoi(result[0])
          n2, _ := strconv.Atoi(result[1])
          res += n1 * n2
        }
      }
    }
  }
  println(res)
}
