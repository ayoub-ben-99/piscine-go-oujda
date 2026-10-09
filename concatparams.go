package piscine

func ConcatParams(args []string) string {
	newArr := ""
	for i, v := range args {
		for _, j := range v {
			newArr += string(j)
		}
		if len(args)-1 != i {
			newArr += "\n"
		}
	}
	return newArr
}
