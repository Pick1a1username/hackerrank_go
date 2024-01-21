package main

const (
	MAX int = 1000003
)

var (
	v     [][]int32 = make([][]int32, MAX) // relationships
	visit []bool    = make([]bool, MAX)
	a     []int64   = make([]int64, MAX) // the color(id) of each node
	c     int                            // the return of bfs()
	id    int64                          // target id(color)
)

type Pair struct {
	Id       int64
	Distance int
}

func bfs(i int64) {
	q := []Pair{}
	q = append(q, Pair{Id: i, Distance: 0})
	visit[i] = true

	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		for _, x := range v[p.Id] {
			if !visit[x] {
				if a[x] == id {
					c = p.Distance + 1
					return
				}
				visit[x] = true
				q = append(q, Pair{Id: int64(x), Distance: p.Distance + 1})
			}
		}
	}
}

func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, ids []int64, val int32) int32 {
	// solve here
	// https://www.hackerrank.com/challenges/find-the-nearest-clone/editorial?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=graphs
	// the initialization here is necessary if you want to run unit tests.
	initGlobalVars()
	if len(graphFrom) != len(graphTo) {
		panic("findShortest() the length of graphFrom and the length of graphTo are not same.")
	}
	for i := 0; i < len(graphFrom); i++ {
		x := graphFrom[i]
		y := graphTo[i]
		x -= 1
		y -= 1
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}
	a = ids
	id = int64(val)

	ansMax := 1000000000
	ans := ansMax
	for i := 0; i < int(graphNodes); i++ {
		if a[i] == id {
			c = ansMax
			bfs(int64(i))
			ans = min(ans, c)
		}
	}
	if ans == ansMax {
		ans = -1
	}
	return int32(ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initGlobalVars() {
	v = make([][]int32, MAX)
	visit = make([]bool, MAX)
	a = make([]int64, MAX)
	c = 0
	id = 0
}
