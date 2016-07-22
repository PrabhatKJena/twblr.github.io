package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	for i:= range vals {
		vals[i] = op(vals[i])
	}
	return vals
}

func filterInts(op filterOperation, vals []int32) []int32 {
	a := make([]int32,0)
	for i:= range vals {
		if r:= op(vals[i]); r==true{
			a = append(a, vals[i])
		}
	}
	return a[:]
}

func concatenate(dest []string, newValues ...string) []string {
	d := make([]string, len(dest))
	d = dest[:]
	for _,s:= range newValues{
		d = append(d, s)
	} 
	return d
}

func equals(list1 []string, list2 []string) bool {
	l1 := len(list1)
	l2 := len(list2)
	if l1 == 0 && l2 == 0{
		return true
	}
	if l1 != l2{
		return false
	}
	for i,s:=range list1{
		if(s != list2[i]){
			return false
		}
	}
	return true
}

func partialReverse(src []int, from, to int) []int {
	rev := make([]int, 0)
	for i:=to;i>=from;i--{
		rev = append(rev, src[i])
	}
	return rev
}
