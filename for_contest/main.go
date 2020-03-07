package main

func sum(a []int) int {
	var res int
	for _, v := range a {
		res += v
	}
	return res
}

func unique(a []string) []string {
	res := make([]string, 0)
	dic := map[string]bool{}
	for _, v := range a {
		if _, ok := dic[v]; !ok {
			res = append(res, v)
			dic[v] = true
		}
	}
	return res
}

func contain(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
