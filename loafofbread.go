package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var res string
	j := 0
	for i := 0; i < len(str); i++ {
		if j < 5 && str[i] == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			res += " "
			j = 0
			continue
		}
		res += string(str[i])
		j++
	}
	res += "\n"
	return res
}
