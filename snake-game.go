package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
  "sync"

  term "github.com/nsf/termbox-go"
)

var (
  lock = sync.Mutex{}
  x = 1
  y = 1
  direction = "d"
  coordinates[25][25] string
)

const snake = "*"

func clearCMD() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func takeInput() string {
  for {
    switch ev := term.PollEvent(); ev.Type {
      case term.EventKey:
        switch ev.Key {
          case term.KeyArrowRight:
            direction = "d"
          case term.KeyArrowLeft:
            direction = "a"
          case term.KeyArrowDown:
            direction = "s"
          case term.KeyArrowUp:
            direction = "w"
        }
    }
  }
  
  if direction == "w" || direction == "s" || direction == "a" || direction == "d" {
    return direction 
  }
  return direction
}

func main() {
  duration := 250 * time.Millisecond
  
  term.Init()

  clearCMD()

  go takeInput()

  for {
    clearCMD()

    for i := 0; i < 25; i++ {
      for j := 0; j < 25; j++ {
        coordinates[i][j] = " "
      }
    }

    lock.Lock()
    switch direction {
    case "w":
      y--
    case "a":
      x--
    case "s":
      y++
    case "d":
      x++
    }

    if y < 0 || y > 24 || x < 0 || x > 24 {
      os.Exit(1)
    }

    coordinates[y][x] = snake
    lock.Unlock()

    for i := 0; i < len(coordinates); i++ {
      for j := 0; j < len(coordinates[i]); j++ {
        fmt.Printf(coordinates[i][j])
      }
      fmt.Println()
    }
    time.Sleep(duration)
  }
}
