package cpu

type ChannelBus struct {
	Q   chan int
	Log []int
}

func (b *ChannelBus) In() int {
	return <-b.Q // this will block
}

func (b *ChannelBus) Out(output int) {
	b.Log = append(b.Log, output)
	b.Q <- output
}
