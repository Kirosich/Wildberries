package main

func UnpackString(str string) string {
	var resString string
	for i := 0; i < len(str); i++ {
		IsNextNum := false
		PrevIsNum := false
		NextNum := 0

		_, IsNum := DetectRuneNumber(str, i)
		if len(str) != i+1 {
			NextNum, IsNextNum = DetectRuneNumber(str, i+1)
		}
		if i != 0 {
			_, PrevIsNum = DetectRuneNumber(str, i-1)
		}

		if !IsNum && IsNextNum {
			for j := 0; j < NextNum; j++ {
				resString += string(str[i])
			}
		} else if IsNum && PrevIsNum {
			resString = "[ERROR] - Некорректная строка"
			return resString
		} else if !IsNum {
			resString += string(str[i])
		}

	}

	return resString

}

func DetectRuneNumber(str string, i int) (int, bool) {
	var Num int
	if str[i] == 49 || str[i] == 50 || str[i] == 51 || str[i] == 52 || str[i] == 53 || str[i] == 54 || str[i] == 55 || str[i] == 56 || str[i] == 57 {
		Num = int(str[i] - '0')
	} else {
		return 0, false
	}
	return Num, true
}
