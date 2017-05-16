package tinker

// Counter class
type Counter struct {
	n int
}

// Method reset() sets the count to 0
func (this *Counter) Reset() {
	this.n = 0
}

// Method next() returns and increments the count
func (this *Counter) Next() int {
	n := this.n
	this.n++
	return n
}

