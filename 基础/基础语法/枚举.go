package main

func main() {
	const (
		_  = iota //0
		KB = 1 << (10 * iota)
		MB
		GB
	)
	//中断iota自增，则必须显示恢复。且后续自增值按行序递增
	const (
		a = iota
		b
		c = 100
		d
		e = iota
		f
	)
	println(c)
}
