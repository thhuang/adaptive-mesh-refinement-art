package entity

import (
	"container/list"
	"math"
)

type Quadtree struct {
	maxLv  int
	mat    [][]bool
	layers [][][]bool
	root   *node
}

func (t *Quadtree) Init(mat [][]bool, maxLv int) *Quadtree {
	size := len(mat) - 1
	root := &node{
		lv:   0,
		size: size,
		nw:   coord{0, 0},
		sw:   coord{size, 0},
		se:   coord{size, size},
		ne:   coord{0, size},
	}

	lvSize := make([]int, maxLv)
	layers := make([][][]bool, maxLv)
	for lv := 0; lv < maxLv; lv++ {
		count := int(math.Pow(2, float64(lv)))
		lvSize[lv] = (size) / count
		layers[lv] = make([][]bool, count)
		for c := 0; c < count; c++ {
			layers[lv][c] = make([]bool, count)
		}

		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat); j++ {
				if !mat[i][j] {
					continue
				}

				// (2, 5)
				// 0 -> 1 -> 8: (0-2, 0-5)
				// 1 -> 2 -> 4: (0-2, 1-1)
				// 2 -> 4 -> 2: (1-0, 2-1)
				// 3 -> 8 -> 1: (2-0, 5-0)

				ri, rj := i%lvSize[lv], j%lvSize[lv]
				if ri*rj == 0 {
					continue
				}

				qi, qj := i/lvSize[lv], j/lvSize[lv]
				layers[lv][qi][qj] = true
			}
		}
	}

	t = &Quadtree{maxLv, mat, layers, root}
	return t
}

func (t *Quadtree) Refine() {
	q := list.New()
	q.PushBack(t.root)
	var next *list.Element
	for e := q.Front(); e != nil; e = next {
		n := e.Value.(*node)

		children := n.subdivide()
		for _, c := range children.getNodes() {
			if t.shouldSubdivide(c) {
				q.PushBack(c)
			}
		}

		next = e.Next()
		q.Remove(e)
	}
}

func (t *Quadtree) GetAMRMat() [][]bool {
	mat := make([][]bool, len(t.mat))
	for i := range mat {
		mat[i] = make([]bool, len(t.mat))
	}

	t.root.plot(mat)

	return mat
}

func (t *Quadtree) GetLayerMat(lv int) [][]bool {
	mat := make([][]bool, len(t.layers[lv]))
	for i := range mat {
		mat[i] = make([]bool, len(t.layers[lv]))
	}

	for i, r := range t.layers[lv] {
		for j, v := range r {
			mat[i][j] = v
		}
	}

	return mat
}

func (t *Quadtree) shouldSubdivide(node *node) bool {
	if node.lv == t.maxLv {
		return false
	}
	if node.children != nil {
		return false
	}

	i, j := node.nw.x/node.size, node.nw.y/node.size
	return t.layers[node.lv][i][j]
}

type coord struct {
	x int
	y int
}

type node struct {
	lv   int
	size int

	nw coord
	sw coord
	se coord
	ne coord

	parent    *node
	children  *children
	neighbors *neighbors
}

func (n *node) subdivide() *children {
	///////////////////
	// (0)--(7)--(3) //
	//  |    |    |  //
	// (4)--(8)--(6) //
	//  |    |    |  //
	// (1)--(5)--(2) //
	///////////////////

	coords := []coord{
		n.nw,
		n.sw,
		n.se,
		n.ne,
		{(n.nw.x + n.sw.x) / 2, (n.nw.y)},
		{n.sw.x, (n.sw.y + n.se.y) / 2},
		{(n.ne.x + n.se.x) / 2, (n.ne.y)},
		{n.nw.x, (n.nw.y + n.ne.y) / 2},
		{(n.nw.x + n.se.x) / 2, (n.nw.y + n.se.y) / 2},
	}

	n.children = &children{
		nw: &node{nw: coords[0], sw: coords[4], se: coords[8], ne: coords[7], lv: n.lv + 1, parent: n, size: n.size / 2},
		sw: &node{nw: coords[4], sw: coords[1], se: coords[5], ne: coords[8], lv: n.lv + 1, parent: n, size: n.size / 2},
		se: &node{nw: coords[8], sw: coords[5], se: coords[2], ne: coords[6], lv: n.lv + 1, parent: n, size: n.size / 2},
		ne: &node{nw: coords[7], sw: coords[8], se: coords[6], ne: coords[3], lv: n.lv + 1, parent: n, size: n.size / 2},
	}

	return n.children
}

func (n *node) plot(mat [][]bool) {
	xArr := make([]int, n.size+1)
	yArr := make([]int, n.size+1)
	for i := 0; i <= n.size; i++ {
		xArr[i] = n.nw.x + i
		yArr[i] = n.nw.y + i
	}

	for _, x := range xArr {
		mat[x][yArr[0]] = true
		mat[x][yArr[n.size]] = true
	}
	for _, y := range yArr {
		mat[xArr[0]][y] = true
		mat[xArr[n.size]][y] = true
	}

	if n.children == nil {
		return
	}

	for _, c := range n.children.getNodes() {
		c.plot(mat)
	}
}

type neighbors struct {
	nw *node
	ne *node
	sw *node
	se *node
	en *node
	es *node
	wn *node
	ws *node
}

type children struct {
	nw *node
	sw *node
	se *node
	ne *node
}

func (c *children) getNodes() []*node {
	return []*node{c.nw, c.sw, c.se, c.ne}
}
