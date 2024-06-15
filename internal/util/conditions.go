package util

// yield condition will run if all the condition in array are true
func ConditionAnd(conditions []bool, yield func()) {
	for _, v := range conditions {
		if !v {
			return
		}
	}
	yield()
}
