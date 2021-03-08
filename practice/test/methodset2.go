package main

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInfo(appender Appender, start, end int) {
	for i := start; i <= end; i++ {
		appender.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {

	//var lst List
	//fmt.Println("lst: ", lst)
	////CountInfo(lst, 1, 10)
	//if LongEnough(lst) {
	//	fmt.Printf("- lst is long enough\n")
	//}
	//
	//plst := new(List)
	//CountInfo(plst, 1, 10)
	//
	//if LongEnough(plst) {
	//	fmt.Printf("- plst is long enough\n")
	//}

}
