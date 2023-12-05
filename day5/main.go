package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(content), "\n\n")

    var maps [][][3]int

    for _, m := range lines[1:] {
        rows := strings.Split(m, "\n")[1:]
        var currentMap [][3]int
        for _, row := range rows {
            values := strings.Fields(row)
            var ints [3]int
            for i, v := range values {
                num, err := strconv.Atoi(v)
                if err != nil {
                    panic(err)
                }
                ints[i] = num
            }
            currentMap = append(currentMap, ints)
        }
        maps = append(maps, currentMap)
    }

    maps = maps[:len(maps)-1]

    var res []int
    // Taken from elixir code
    ranges := []struct{ start, end int }{
        {858905075, 915841667},
        {947763189, 1214782614},
        {206349064, 458758537},
        {660226451, 752787537},
        {752930744, 777092798},
        {75704321, 139305268},
        {3866217991, 4189695523},
        {3356941271, 3411310160},
        {1755537789, 2231075088},
        {1327269841, 1754929574},
    }

    for _, r := range ranges{
        mnSoFar:= 9999999999;
        for seedNum:= r.start; seedNum <= r.end; seedNum++{

            result := seedNum

            for _, m := range maps {
                s_start := result
                for _, s := range m {
                    dest, src, length := s[0], s[1], s[2]
                    if result >= src && result < src+length && result == s_start {
                        result = dest + result - src
                        break
                    }
                }
            }
            if result < mnSoFar {
                mnSoFar = result
            }

        }
        
        fmt.Println(mnSoFar)
        res = append(res,mnSoFar)
    }

    minResult := res[0]
    for _, r := range res {
        if r < minResult {
            minResult = r
        }
    }

    fmt.Println(minResult)
}

