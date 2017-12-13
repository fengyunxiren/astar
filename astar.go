package astar

import (
  "math"
  "fmt"
  "container/heap"
)

func Astar(start, end Point, m *MapList) {
  if !m.In(start) || !m.In(end) {
    panic("start or end not in map!")
  }
  openDq := StarDqeue{}
  closeDq := StarDqeue{}
  startNode := NewStarNode(start, end, nil)
  var ret StarNode
  heap.Init(&openDq)
  heap.Init(&closeDq)
  heap.Push(&openDq, startNode)
  out:
  for openDq.Len() > 0 {
    tmp := heap.Pop(&openDq).(StarNode)
    if tmp.Current == end {
      ret = tmp
      break out
    }
    m.Write(tmp.Current, 5)
    neighber := m.Neighber(tmp.Current)
    for _, n := range neighber {
      s := NewStarNode(n, end, &tmp)
      m.Write(n, 4)
      heap.Push(&openDq, s)
    }
    heap.Push(&closeDq, tmp)
    m.Print()
    fmt.Printf("\n")
  }
  if ret.Current != end {
    fmt.Println("Search path fail!")
    return
  }
  for ret.Parent != nil {
    m.Write(ret.Current, 6)
    ret = *(ret.Parent)
  }
  m.Write(start, 1)
  m.Write(end, 2)
  m.Print()
}

type StarNode struct{
  Parent *StarNode
  Current Point
  F float64
  G float64
}

func NewStarNode(current, end Point, parent *StarNode) StarNode {
  var g float64
  if parent != nil {
    g = parent.G + 1
  }
  h := current.H(end)
  f := g + h
  ret := StarNode{
    Parent: parent,
    Current: current,
    F: f,
    G: g,
  }
  return ret
}

type Point struct{
  X int
  Y int
}

func (p *Point) H(end Point) float64 {
  return math.Abs(float64(p.X - end.X)) + math.Abs(float64(p.Y - end.Y))
}

type StarDqeue []StarNode

func (s StarDqeue) Less(i, j int) bool {
  if s[i].F == s[j].F {
    return s[i].G > s[j].G
  } else {
  return s[i].F < s[j].F
  }
}

func (s *StarDqeue) Swap(i, j int) {
  (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s StarDqeue) Len() int {
  return len(s)
}

func (s *StarDqeue) Push(n interface{}) {
  e := n.(StarNode)
  *s = append(*s, e)
}

func (s *StarDqeue) Pop() interface{} {
  old := *s
  n := len(old)
  if n == 0 {
    panic("Error: No data pop!")
  }
  e := old[n-1]
  *s = old[:n-1]
  return e
}

type MapList [][]int

func NewMapList(x, y int) MapList {
  m := MapList{}
  for i := 0; i < x; i++ {
    v := make([]int, y)
    m = append(m, v)
  }
  return m
}

func (m MapList) Print() {
  for _, row := range m {
    for _, col := range row {
      switch col {
        case 1: fmt.Printf("s ")
        case 2: fmt.Printf("e ")
        case 3: fmt.Printf("# ")
        case 4: fmt.Printf("^ ")
        case 5: fmt.Printf("! ")
        case 6: fmt.Printf("* ")
        default: fmt.Printf(". ")
      }
    }
    fmt.Printf("\n")
  }
}

func (m MapList) In(p Point) bool {
  if p.X >= 0 && p.Y >=0 && p.X < len(m) && p.Y < len(m[0]) {
    return true
  }
  return false
}

func (m *MapList) Write(p Point, v int) {
  if m.In(p) {
    (*m)[p.X][p.Y] = v
  }
}

func (m MapList) Read(p Point) int {
  if m.In(p) {
    return m[p.X][p.Y]
  }
  return -1
}

func (m MapList) Neighber(p Point) []Point {
  ret := []Point{}
  p1 := Point{p.X-1, p.Y}
  p2 := Point{p.X+1, p.Y}
  p3 := Point{p.X, p.Y-1}
  p4 := Point{p.X, p.Y+1}
  if m.Read(p1) == 0 {
    ret = append(ret, p1)
  }
  if m.Read(p2) == 0 {
    ret = append(ret, p2)
  }
  if m.Read(p3) == 0 {
    ret = append(ret, p3)
  }
  if m.Read(p4) == 0 {
    ret = append(ret, p4)
  }
  return ret
}
