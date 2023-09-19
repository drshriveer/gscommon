// Code generated by gerrors DO NOT EDIT.
package internal

// GSortSortable implements the sort.Sort interface for Sortable.
type GSortSortable []Sortable

func (s GSortSortable) Len() int {
	return len(s)
}
func (s GSortSortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s GSortSortable) Less(i, j int) bool {
	if s[i].Category.String() == s[j].Category.String() {
		if s[i].Property1 == s[j].Property1 {
			return s[i].Property2 < s[j].Property2
		}
		return s[i].Property1 < s[j].Property1
	}
	return s[i].Category.String() < s[j].Category.String()
}
