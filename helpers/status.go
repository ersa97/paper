package helpers

func SetStatus(data string) int {
	switch data {
	case "posted":
		return 1
	case "reviewed":
		return 2
	case "approved":
		return 3
	case "executed":
		return 4
	case "cancelled":
		return 5
	case "aborted":
		return 6
	default:
		return 0
	}
}

// GetStatus is
func GetStatus(data int) string {
	switch data {
	case 1:
		return "posted"
	case 2:
		return "reviewed"
	case 3:
		return "approved"
	case 4:
		return "executed"
	case 5:
		return "cancelled"
	case 6:
		return "aborted"
	default:
		return "undefined"
	}
}
