package main

import (
  "fmt"
  "math/rand"
  "os"
  "os/exec"
  "sync"
  "time"

  term "github.com/nsf/termbox-go"
)

var (
  lock        = sync.Mutex{}
  xApple      int
  yApple      int
  direction   = "d"
  coordinates [25][25]string
  ln          = 1
)

var (
  x = 12
  y = 12

  oldX int
  oldY int

  x2 = x - 1
  y2 = y

  oldX2 int
  oldY2 int

  x3 = x - 2
  y3 = y

  oldX3 int
  oldY3 int

  x4 = x - 3
  y4 = y

  oldX4 int
  oldY4 int

  x5 = x - 4
  y5 = y

  oldX5 int
  oldY5 int
)

const (
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

func makeSnake() {
  coordinates[y][x] = snake
  coordinates[y2][x2] = snake
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
  if ln == 2 {
    lock.Lock()
    if direction == "w" {
      oldX, oldY = x, y
      y--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      coordinates[oldY2][oldX2] = empty
    } else if direction == "d" {
      oldX, oldY = x, y
      x++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      coordinates[oldY2][oldX2] = empty
    } else if direction == "a" {
      oldX, oldY = x, y
      x--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      coordinates[oldY2][oldX2] = empty
    } else if direction == "s" {
      oldX, oldY = x, y
      y++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      coordinates[oldY2][oldX2] = empty
    }

    lock.Unlock()
  } else if ln == 3 {
    lock.Lock()
    if direction == "w" {
      oldX, oldY = x, y
      y--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      coordinates[oldY3][oldX3] = empty
    } else if direction == "d" {
      oldX, oldY = x, y
      x++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      coordinates[oldY3][oldX3] = empty
    } else if direction == "a" {
      oldX, oldY = x, y
      x--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      coordinates[oldY3][oldX3] = empty
    } else if direction == "s" {
      oldX, oldY = x, y
      y++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      coordinates[oldY3][oldX3] = empty
    }
    lock.Unlock()
  } else if ln == 4 {
    lock.Lock()
    if direction == "w" {
      oldX, oldY = x, y
      y--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      coordinates[oldY4][oldX4] = empty
    } else if direction == "d" {
      oldX, oldY = x, y
      x++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      coordinates[oldY4][oldX4] = empty
    } else if direction == "a" {
      oldX, oldY = x, y
      x--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      coordinates[oldY4][oldX4] = empty
    } else if direction == "s" {
      oldX, oldY = x, y
      y++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      coordinates[oldY4][oldX4] = empty
    }
    lock.Unlock()
  } else if ln == 5 {
    lock.Lock()
    if direction == "w" {
      oldX, oldY = x, y
      y--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      oldX5, oldY5 = x5, y5
      x5, y5 = oldX4, oldY4
      coordinates[y5][x5] = snake
      coordinates[oldY5][oldX5] = empty
    } else if direction == "d" {
      oldX, oldY = x, y
      x++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      oldX5, oldY5 = x5, y5
      x5, y5 = oldX4, oldY4
      coordinates[y5][x5] = snake
      coordinates[oldY5][oldX5] = empty
    } else if direction == "a" {
      oldX, oldY = x, y
      x--
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      oldX5, oldY5 = x5, y5
      x5, y5 = oldX4, oldY4
      coordinates[y5][x5] = snake
      coordinates[oldY5][oldX5] = empty
    } else if direction == "s" {
      oldX, oldY = x, y
      y++
      coordinates[y][x] = snake
      oldX2, oldY2 = x2, y2
      x2, y2 = oldX, oldY
      coordinates[y2][x2] = snake
      oldX3, oldY3 = x3, y3
      x3, y3 = oldX2, oldY2
      coordinates[y3][x3] = snake
      oldX4, oldY4 = x4, y4
      x4, y4 = oldX3, oldY3
      coordinates[y4][x4] = snake
      oldX5, oldY5 = x5, y5
      x5, y5 = oldX4, oldY4
      coordinates[y5][x5] = snake
      coordinates[oldY5][oldX5] = empty
    }
    lock.Unlock()
  }
}

func gameOver() {
  if y < 1 || y > 23 || x < 1 || x > 23 {
    os.Exit(1)
  }
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

    xApple = rand.Intn(23-2) + 1
    yApple = rand.Intn(23-2) + 1

    coordinates[yApple][xApple] = apple

    ln++
  }
}

func main() {
  duration := 17 * time.Millisecond

  term.Init()

  clearCMD()
  makeField()
  makeSnake()
  go takeInput()

  for i := 10; ; i++ {
    clearCMD()
    if i%10 == 0 {
      moveSnake()
    }
    drawField()
    makeApple()
    gameOver()

    time.Sleep(duration)
  }
}
