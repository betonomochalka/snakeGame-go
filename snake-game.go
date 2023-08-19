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
  x = 0
  y = 0
  direction = "d"
)

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

func move() {
  duration := 1 * time.Second

	for {
    lock.Lock() 
      switch direction {
    case "w":
      if y != 0 {
        y--
      }
    case "a":
      if x != 0 {
        x-- 
      }
    case "s":
      if y != 49 {
        y++
      }
    case "d":
      if x != 99 {
        x++
      }
    }
    lock.Unlock()
    time.Sleep(duration)
  }
}

func main() {
  duratin := 100 * time.Millisecond

  snake := "*"
  
  term.Init()

  clearCMD()

  go move()
  go takeInput()

  var coordinates[50][100] string



  for {
   clearCMD()  

    for i := 0; i < 50; i++ {
      for j := 0; j < 100; j++ {
        coordinates[i][j] = "-"
      }
    }

    lock.Lock()
    y, x := y, x
    lock.Unlock()
    
    coordinates[y][x] = snake

    for i := 0; i < len(coordinates); i++ {
      for j := 0; j < len(coordinates[i]); j++ {
        fmt.Printf(coordinates[i][j])
      }
      fmt.Println()
    }
    time.Sleep(duratin)
  }
}
