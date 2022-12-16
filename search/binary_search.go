/*
1. Escribe una funcion que realice una buqueda binaria de un ID `n` sobre el siguiente slice:

	 `productIDs := []int {2, 5, 7, 9, 23, 28, 36, 39, 56, 68, 75, 76, 77, 88, 89, 99}`
	y retorne la posición donde se encontro y el numero de iteraciones para encontrarlo.
*/
package main

func searchProductsBinary(productIDs []int, encontrar int) (int, int) {
	var i int
	var la = len(productIDs) - 1
	var iter = 0

	for i <= la {
		middle := (la-1)/2 + 1
		if productIDs[middle] == encontrar {
			iter += 1
			return middle, iter
		}

		if productIDs[middle] > encontrar {
			la = middle - 1
		} else {
			i = middle + 1
		}
	}

	return -1, -1
}
