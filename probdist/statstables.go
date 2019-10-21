package stats

type statsTable struct {
	table [][]float64
	row   []float64
	col   []float64
}

// interpolateCols returns the value which is between the two known columns supposing a lineal distribution between them.
//    l0   ----  c0 --- r0.
//    l    ----  ?  --- r.
func (tab statsTable) interpolateCols(c0 float64, rowidx int) float64 {
	l0, r0, lid, rid := tab.getColValues(c0)
	l := tab.table[rowidx][lid]
	r := tab.table[rowidx][rid]
	if l0 == r0 {
		return tab.row[rowidx]
	}
	return ((r-l)/(r0-l0))*(c0-l0) + l
}

func (tab statsTable) getColValues(c0 float64) (l0 float64, r0 float64, lid int, rid int) {
	if c0 > tab.col[0] || c0 < tab.col[len(tab.col)-1] {
		panic("stats: c0 value not valid in the table")
	}

	for i, a := range tab.col[:len(tab.col)-1] {
		if c0 <= a {
			l0, r0 = a, float64(tab.col[i+1])
			lid, rid = i, i+1
			break
		}
	}
	return
}
