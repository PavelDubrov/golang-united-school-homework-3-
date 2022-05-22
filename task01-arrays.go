package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var elements float32
	sum = 0
	elements = float32(len(input))

	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	result = sum / elements
	return result
}
