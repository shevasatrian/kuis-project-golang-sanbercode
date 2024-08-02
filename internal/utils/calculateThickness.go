package utils

func CalculateThickness(totalPage int) string {
	if totalPage > 100 {
		return "tebal"
	}
	return "tipis"
}
