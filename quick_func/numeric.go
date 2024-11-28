package quick_func

import "strconv"

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func Int32ToString(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func Float32ToString(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', -1, 32)
}
