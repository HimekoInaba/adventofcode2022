package main
import (
  "bufio"
  "os"
  "log"
  "fmt"
  "strings"
)

func main() {
  part2();
}

func part2() {
  file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var score = 0;
    for scanner.Scan() {
      var line = scanner.Text();
      words := strings.Fields(line);
      p1 := words[0];
      p2 := words[1];
      score += (getBaseScoreFromOutcome(p1, p2) + getOutcomeScore(p2));
    }
  fmt.Println(score);
}

func part1() {
  file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var score = 0;
    for scanner.Scan() {
      var line = scanner.Text();
      words := strings.Fields(line);
      p1 := words[0];
      p2 := words[1];
      score += (getBaseScore(p2) + getJudgementScore(p1, p2));
    }
  fmt.Println(score);
}

func getBaseScore(x string) int {
  if x == "X" {
    return 1;
  } else if x == "Y" {
    return 2;
  } else {
    return 3;
  }
}

func getJudgementScore(x, y string) int {
  if x == "A" && y == "X" {
    return 3;
  } else if x == "A" && y == "Y" {
    return 6;
  } else if x == "A" && y == "Z" {
    return 0;
  } else if x == "B" && y == "X" {
    return 0;
  } else if x == "B" && y == "Y" {
    return 3;
  } else if x == "B" && y == "Z" {
    return 6;
  } else if x == "C" && y == "X" {
    return 6;
  } else if x == "C" && y == "Y" {
    return 0;
  } else {
    return 3;
  }
}

func getOutcomeScore(x string) int {
  if x == "X" {
    return 0;
  } else if x == "Y" {
    return 3;
  } else {
    return 6;
  }
}

func getBaseScoreFromOutcome(x, outcome string) int {
  rock := 1
  paper := 2
  scisors := 3
  if outcome == "X" {
    if x == "A" {
      return scisors;
    } else if x == "B" {
      return rock;
    } else {
      return paper;
    }
  } else if outcome == "Y" {
    if x == "A" {
      return rock;
    } else if x == "B" {
      return paper;
    } else {
      return scisors;
    }
  } else {
    if x == "A" {
      return paper;
    } else if x == "B" {
      return scisors;
    } else {
      return rock;
    }
  }
}
