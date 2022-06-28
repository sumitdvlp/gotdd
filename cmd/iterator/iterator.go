package iterator

func Repeat(st string, n int) (op string) {
	for i := 0; i < n; i++ {
		op = op + st
	}
	return
}
