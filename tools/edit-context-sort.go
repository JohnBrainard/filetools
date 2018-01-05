package tools

type FileContextsByName []FileContext

func (a FileContextsByName) Len() int {
	return len(a)
}

func (a FileContextsByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a FileContextsByName) Less(i, j int) bool {
	return a[i].Path < a[j].Path
}

type FileContextsByDate []FileContext

func (a FileContextsByDate) Len() int {
	return len(a)
}

func (a FileContextsByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a FileContextsByDate) Less(i, j int) bool {
	return a[j].FileInfo.ModTime().After(a[i].FileInfo.ModTime())
}
