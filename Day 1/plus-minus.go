package main

import "fmt"

func main() {
	arr := []int32{1, 1, 0, -1, -1}

	np := 0
	nn := 0
	nz := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			nn++
		} else if arr[i] > 0 {
			np++
		} else {
			nz++
		}
	}

	total := np + nn + nz

	pnp := float64(np) / float64(total)
	pnn := float64(nn) / float64(total)
	pnz := float64(nz) / float64(total)

	fmt.Printf("%.6f\n", pnp)
	fmt.Printf("%.6f\n", pnn)
	fmt.Printf("%.6f\n", pnz)

}
