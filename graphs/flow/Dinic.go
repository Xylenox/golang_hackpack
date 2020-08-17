
type Dinic struct {
	lvl, ptr, q []int
	adj         [][]*Edge
}

type Edge struct {
	to, rev int
	c, oc   int64
}

func (d *Dinic) addEdge(a, b int, c int64) {
	d.adj[a] = append(d.adj[a], &Edge{b, len(d.adj[b]), c, c})
	d.adj[b] = append(d.adj[b], &Edge{a, len(d.adj[a]) - 1, 0, 0})
}

func (d *Dinic) dfs(v, t int, f int64) int64 {
	if v == t || f == 0 {
		return f
	}

	for i := d.ptr[v]; i < len(d.adj[v]); i++ {
		e := d.adj[v][i]
		if d.lvl[e.to] == d.lvl[v]+1 {
			p := d.dfs(e.to, t, Min(f, e.c))
			if p != 0 {
				e.c -= p
				d.adj[e.to][e.rev].c += p
				return p
			}
		}
	}
	return 0
}

func (d *Dinic) calc(s, t int) int64 {
	var flow int64 = 0
	d.q[0] = s
	for L := 0; L < 31; L++ {
		for {
			d.ptr = make([]int, len(d.q))
			d.lvl = make([]int, len(d.q))
			qi, qe := 0, 1
			d.lvl[s] = 1
			for qi < qe && d.lvl[t] == 0 {
				v := d.q[qi]
				qi++
				for _, e := range d.adj[v] {
					if d.lvl[e.to] == 0 && e.c>>(30-L) != 0 {
						d.q[qe] = e.to
						d.lvl[e.to] = d.lvl[v] + 1
						qe++
					}
				}
			}
			for {
				p := d.dfs(s, t, 1<<63-1)
				if p == 0 {
					break
				}
				flow += p
			}
			if d.lvl[t] == 0 {
				break
			}
		}
	}
	return flow
}

func solve() {
	n, m := nextInt(), nextInt()
	tot := 0
	flow := &Dinic{make([]int, 2*n+2), make([]int, 2*n+2), make([]int, 2*n+2), make([][]*Edge, 2*n+2)}
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		a[i], b[i] = nextInt(), nextInt()
		s := (nextInt() + m - 1) / m
		tot += s
		flow.addEdge(2*n, i, int64(s))
		flow.addEdge(i+n, 2*n+1, int64(s))
	}
	for i := range a {
		for j := range a {
			c := nextInt()
			if b[i]+c < a[j] {
				flow.addEdge(i, j+n, 1<<30)
			}
		}
	}
	printf("%d\n", int64(tot)-flow.calc(2*n, 2*n+1))
}
