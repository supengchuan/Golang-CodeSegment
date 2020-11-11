package container

type (
	Set struct {
		hash map[interface{}]nothing
	}

	nothing struct{}
)

func NewSet(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]nothing)}

	for _, v := range initial {
		s.Insert(v)
	}
	return s
}

//  find difference between two sets
func (this *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range this.hash {
		if _, ok := set.hash[k]; !ok {
			n[k] = nothing{}
		}
	}
	return &Set{n}
}

func (this *Set) Do(f func(interface{})) {
	for k := range this.hash {
		f(k)
	}
}

func (this *Set) Has(element interface{}) bool {
	_, ok := this.hash[element]
	return ok
}

func (this *Set) Insert(v interface{}) {
	this.hash[v] = nothing{}
}

func (this *Set) Remove(element interface{}) {
	delete(this.hash, element)
}
