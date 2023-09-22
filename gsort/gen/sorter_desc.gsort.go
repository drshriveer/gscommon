// Code generated by gsort DO NOT EDIT.
package gen

// SortFieldDescs implements the sort.Sort interface for SortFieldDesc.
type SortFieldDescs []*SortFieldDesc

func (s SortFieldDescs) Len() int {
	return len(s)
}
func (s SortFieldDescs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortFieldDescs) Less(i, j int) bool {
	return s[i].Priority < s[j].Priority
}

// SorterDescs implements the sort.Sort interface for SorterDesc.
type SorterDescs []*SorterDesc

func (s SorterDescs) Len() int {
	return len(s)
}
func (s SorterDescs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SorterDescs) Less(i, j int) bool {
	if s[i].TypeName == s[j].TypeName {
		return s[i].SortType < s[j].SortType
	}
	return s[i].TypeName < s[j].TypeName
}
