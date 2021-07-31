package pkg_sort

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (us UserSlice) Len() int {
	return len(us)
}

func (us UserSlice) Less(i, j int) bool {
	return us[i].Age < us[j].Age
}

func (us UserSlice) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}
