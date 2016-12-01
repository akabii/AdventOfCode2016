package main

import (
    "fmt"
    "io/ioutil"
    "math"
    "strings"
    "strconv"
)

func main() {
    dat, err := ioutil.ReadFile("input")
    check(err)
    directions := strings.Split(string(dat), ", ")
    problem1(directions)
    problem2(directions)
}

func problem1(directions []string) {
    var x, y, dir int = 0, 0, 0

    for i := 0; i < len(directions); i++ {
        direction := directions[i][0];
        steps, _ := strconv.Atoi(directions[i][1:len(directions[i])])
        if direction == 'R' {
            dir = (dir + 1) & 3;
        } else {
            dir = (dir - 1) & 3;
        }

        switch dir {
            case 0:
                y += steps
            case 1:
                x += steps
            case 2:
                y -= steps
            case 3:
                x -= steps
        }
    }

    fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

type Point struct {
    x int
    y int
}

type Locations map[Point] bool

func problem2(directions []string) {

    visited := make(Locations)

    var dir int = 0
    var location = Point{x: 0, y: 0}
    visited[location] = true

    for i := 0; i < len(directions); i++ {
        direction := directions[i][0];
        steps, _ := strconv.Atoi(directions[i][1:len(directions[i])])
        if direction == 'R' {
            dir = (dir + 1) & 3;
        } else {
            dir = (dir - 1) & 3;
        }

        for j := 0; j < steps; j++ {
            switch dir {
                case 0:
                    location.y += 1
                case 1:
                    location.x += 1
                case 2:
                    location.y -= 1
                case 3:
                    location.x -= 1
            }

        // fmt.Println("%v", location)
            if _, ok := visited[location]; ok {
                fmt.Println(math.Abs(float64(location.x)) + math.Abs(float64(location.y)))
                return;
            } else {
                visited[location] = true
            }
        }
    }
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}
