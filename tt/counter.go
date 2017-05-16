package main

// Counter class
type Counter struct {
	n int
}

// Method reset() sets the count to 0
func (this *Counter) reset() {
	this.n = 0
}

// Method next() returns and increments the count
func (this *Counter) next() int {
	n := this.n
	this.n++
	return n
}

