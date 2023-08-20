package main

import (
  "fmt"
  "os"
  "os/exec"
  "time"
  "sync"
  "math/rand"

  term "github.com/nsf/termbox-go"
)

var (
  lock = sync.Mutex{}
  x = 1
  y = 5
  xApple int
  yApple int
  direction = "d"
  coordinates[25][25] string
)

const(
  snake = "[]"
  apple = "<>"
  empty = "　"
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
        lock.Lock()
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
    lock.Unlock()
  }
}

func drawField() {
  for i := 0; i < len(coordinates); i++ {
    for j := 0; j < len(coordinates[i]); j++ {
      fmt.Printf(coordinates[i][j])
    }
    fmt.Println()
  }
}

func moveSnake() {
  oldx := x
  oldy := y
  lock.Lock()
  switch direction {
  case "w":
    y--
    coordinates[oldy][oldx] = empty
  case "a":
    x--
    coordinates[oldy][oldx] = empty
  case "s":
    y++
    coordinates[oldy][oldx] = empty
  case "d":
    x++
    coordinates[oldy][oldx] = empty
  }
  lock.Unlock()
}

func gameOver() {
  if y < 1 || y > 23 || x < 1 || x > 23 {
    os.Exit(1)
  }
}

func drawSnake() {
  coordinates[y][x] = snake
}

func makeField() {
  for i := 1; i < 24; i++ {
    for j := 1; j < 24; j++ {
      coordinates[i][j] = empty
    }
  } 

  for i := 0; i < 25; i++ {
    for j := 1; j < 24; j++ {
      coordinates[i][0] = "|"
      coordinates[i][24] = "|"
      coordinates[0][j] = "--"
      coordinates[24][j] = "--"
    }
  }
}

func makeApple() {
  if coordinates[yApple][xApple] != apple {
    rand.Seed(time.Now().UnixNano())

    xApple = rand.Intn(23-2)+1
    yApple = rand.Intn(23-2)+1


    if coordinates[yApple][xApple] == snake {
      makeApple()
      return
    }

    coordinates[yApple][xApple] = apple
  }
}

func main() {
  duration := 17 * time.Millisecond
  
  term.Init()

  clearCMD()

  go takeInput()

  makeField()

  for i := 0; ; i++{
    clearCMD()
    drawSnake()
    drawField()
    makeApple()
    if i%10 == 0 {
      moveSnake()
    }
    gameOver()

    time.Sleep(duration)
  }
}
