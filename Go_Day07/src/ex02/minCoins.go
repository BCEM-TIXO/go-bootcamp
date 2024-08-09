package ex02

import (
	"math"
)

func dedupeInts(arr []int) []int {
	m, uniq := make(map[int]struct{}), make([]int, 0, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return uniq
}

/*
MinCoins2 решает задачу нахождения минимального количества монет для суммы val
Оптимизация: решение теперь всегда правильное из-за измененного алгоритма
*/

func MinCoins2(val int, coins []int) []int {
	coins = dedupeInts(coins) // Удаление повторяющихся монет
	dp := make([]int, val+1)  // Массив для хранения минимального количества монет для каждой суммы до val
	for i := 1; i <= val; i++ {
		dp[i] = math.MaxUint32 // Инициализация элементов массива максимальным значением для поиска минимума
	}

	for i := 1; i <= val; i++ {
		for _, coin := range coins {
			if coin >= 0 && i >= coin && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	// Если значение для суммы val осталось максимальным, значит сумму нельзя набрать
	if dp[val] == math.MaxUint32 {
		return []int{}
	}
	result := make([]int, 0)
	for val > 0 {
		for _, coin := range coins {
			// Ищем монету, использование которой приводит к минимальному количеству монет для текущей суммы val
			if coin >= 0 && val >= coin && dp[val-coin] == dp[val]-1 {
				result = append(result, coin)
				val -= coin
				break
			}
		}
	}

	return result
}
