package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "sort"
)

func maximum(x, y int) int {
  if x < y {
    return y
  }
  return x
}

func part1() {
   file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var max, curr int = -1, 0;
    for scanner.Scan() {
        var line = scanner.Text();
        if len(line) == 0 {
            max = maximum(max, curr)
            curr = 0
        } else {
          intValue, err := strconv.Atoi(line);
          if err != nil {
            log.Fatal(err)
          }
          curr += intValue
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Part 1: " + strconv.Itoa(max))
}

func part2() {
   file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var curr = 0;
    list := make([]int, 100)
    for scanner.Scan() {
        var line = scanner.Text();
        if len(line) == 0 {
            list = append(list, curr)
            curr = 0
        } else {
          intValue, err := strconv.Atoi(line);
          if err != nil {
            log.Fatal(err)
          }
          curr += intValue
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    sort.Sort(sort.IntSlice(list))
    res := list[len(list) - 1] + list[len(list) - 2] + list[len(list) - 3]
    fmt.Println(res)
}

func main() {
  part2()
}
