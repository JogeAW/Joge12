package numbers

type Counter struct {
	value   int
	limit   int
	BufChan chan bool
}

func NewCounter(limit int) *Counter {
	c := Counter{
		limit:   limit,
		value:   0,
		BufChan: make(chan bool, 1),
	}
	c.BufChan <- true
	return &c
}

func (c *Counter) Add(amount int) bool {
	<-c.BufChan

	if c.value >= c.limit {
		c.BufChan <- true
		return false
	}

	c.value += amount
	c.BufChan <- true
	return true
}

func (c *Counter) Value() int {
	<-c.BufChan
	v := c.value
	c.BufChan <- true
	return v
}
