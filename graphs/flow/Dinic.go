type Edge struct {
	v1, v2, cap, flow int
	rev               *Edge
}

type Dinic struct {
	adj     [][]*Edge
	n, s, t int
	dist, q []int
	blocked []bool
}

func NewDinic(n int) *Dinic {
	res := &Dinic{make([][]*Edge, n+2), n, n, n + 1,
		make([]int, n+2), []int{}, make([]bool, n+2)}
	for i := range res.adj {
		res.adj[i] = []*Edge{}
	}
	return res
}

func (d *Dinic) Add(u, v, c int) {
	a, b := &Edge{u, v, c, 0, nil}, &Edge{v, u, 0, 0, nil}
	a.rev, b.rev = b, a
	d.adj[u] = append(d.adj[u], a)
	d.adj[v] = append(d.adj[v], b)
}

func (d *Dinic) Bfs() bool {
	d.q = d.q[:0]
	Fill(d.dist, -1)
	d.dist[d.t] = 0
	d.q = append(d.q, d.t)

	for len(d.q) > 0 {
		curr := d.q[0]
		d.q = d.q[1:]
		if curr == d.s {
			return true
		}
		for _, e := range d.adj[curr] {
			if e.rev.cap > e.rev.flow && d.dist[e.v2] == -1 {
				d.dist[e.v2] = d.dist[curr] + 1
				d.q = append(d.q, e.v2)
			}
		}
	}
	return d.dist[d.s] != -1
}

func (d *Dinic) Dfs(ind, min int) int {
	if ind == d.t {
		return min
	}
	flow := 0
	for _, e := range d.adj[ind] {
		if !d.blocked[e.v2] && d.dist[e.v2] == d.dist[ind]-1 && e.cap > e.flow {
			cur := d.Dfs(e.v2, Min(min-flow, e.cap-e.flow))
			e.flow += cur
			e.rev.flow = -e.flow
			flow += cur
		}
		if flow == min {
			return flow
		}
	}
	d.blocked[ind] = flow != min
	return flow
}

func (d *Dinic) Flow() int {
	res := 0
	for d.Bfs() {
		FillB(d.blocked, false)
		res += d.Dfs(d.s, oo)
	}
	return res
}
