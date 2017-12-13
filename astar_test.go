package astar

import (
  "testing"
  "container/heap"
)

func TestAstar(t *testing.T) {
  m := astar.NewMapList(20, 20)
  m.Write(astar.Point{4, 3}, 3)
  m.Write(astar.Point{4, 4}, 3)
  m.Write(astar.Point{4, 5}, 3)
  m.Write(astar.Point{3, 6}, 3)
  m.Write(astar.Point{2, 6}, 3)
  m.Write(astar.Point{1, 6}, 3)
  m.Write(astar.Point{0, 6}, 3)
  m.Write(astar.Point{4, 8}, 3)
  m.Write(astar.Point{3, 8}, 3)
  m.Write(astar.Point{2, 8}, 3)
  m.Write(astar.Point{5, 8}, 3)
  m.Write(astar.Point{6, 8}, 3)
  m.Write(astar.Point{6, 7}, 3)
  m.Write(astar.Point{6, 6}, 3)
  m.Write(astar.Point{6, 5}, 3)
  m.Write(astar.Point{6, 4}, 3)
  m.Write(astar.Point{6, 3}, 3)
  m.Write(astar.Point{6, 2}, 3)
  m.Write(astar.Point{6, 1}, 3)
  m.Write(astar.Point{10, 5}, 3)
  m.Write(astar.Point{10, 4}, 3)
  m.Write(astar.Point{10, 3}, 3)
  m.Write(astar.Point{10, 2}, 3)
  m.Write(astar.Point{10, 1}, 3)
  m.Write(astar.Point{10, 0}, 3)

  start := astar.Point{3, 3}
  end := astar.Point{17, 18}
  astar.Astar(start, end, &m)
}

func TestPointH(t *testing.T) {
  start := Point{1, 3}
  end := Point{4, 8}
  h := start.H(end)
  t.Log(h)
}

func TestNewStarNode(t *testing.T) {
  start := Point{1, 3}
  end := Point{4, 8}
  n1 := NewStarNode(start, end, nil)
  n2 := NewStarNode(start, end, &n1)
  t.Log(n1)
  t.Log(n2)
}

func TestStarDqeue(t *testing.T) {
  start := Point{1, 3}
  current := Point{2, 3}
  end := Point{4, 8}
  n1 := NewStarNode(start, end, nil)
  n2 := NewStarNode(current, end, &n1)
  dq := StarDqeue{}
  heap.Init(&dq)
  heap.Push(&dq, n1)
  heap.Push(&dq, n2)
  t.Log(heap.Pop(&dq))
}

func TestNewMapList(t *testing.T) {
  m := NewMapList(3, 9)
  t.Log(m)
}

func TestMapList(t *testing.T) {
  m := NewMapList(8, 8)
  t.Log(m.In(Point{3, 4}))
  t.Log(m.In(Point{9, 10}))
  m.Write(Point{3, 4}, 2)
  t.Log(m.Read(Point{3, 4}))
  t.Log(m.Neighber(Point{3, 4}))
  m.Print()
}
