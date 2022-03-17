package main

import (
	"fmt"
	"time"
)

func backtracking(tasks []int, sessionTimes []int, result *int, index int, sessionTime int) {
	if len(sessionTimes) >= *result {
		return
	} else if index >= len(tasks) && *result > len(sessionTimes) {
		*result = len(sessionTimes)
	} else {

		for i := 0; i < len(sessionTimes); i++ {
			if sessionTimes[i]+tasks[index] <= sessionTime {
				sessionTimes[i] = sessionTimes[i] + tasks[index]
				backtracking(tasks, sessionTimes, result, index+1, sessionTime)
				sessionTimes[i] = sessionTimes[i] - tasks[index]

			}
		}
		sessionTimes = append(sessionTimes, tasks[index])
		backtracking(tasks, sessionTimes, result, index+1, sessionTime)
		sessionTimes = sessionTimes[:len(sessionTimes)-1]
	}

}

func minSessions(tasks []int, sessionTime int) int {
	sessionTimes := make([]int, 0)
	result := len(tasks)
	backtracking(tasks, sessionTimes, &result, 0, sessionTime)
	return result
}

//
//func RemoveIndex(s []int, index int) []int {
//	ret := make([]int, 0)
//	ret = append(ret, s[:index]...)
//	return append(ret, s[index+1:]...)
//}
//
//func findMinimumSession(tasks []int, minimumSession *int, workPerSession int, currentSession int, sessionTime int, count int, length int) {
//	if count == length {
//		currentSession = currentSession + 1
//		if currentSession < *minimumSession {
//			*minimumSession = currentSession
//		}
//		return
//	}
//
//	notFound := true
//	for i := 0; i < len(tasks); i++ {
//		if workPerSession+tasks[i] <= sessionTime {
//			notFound = false
//			findMinimumSession(RemoveIndex(tasks, i), minimumSession, workPerSession+tasks[i], currentSession, sessionTime, count+1, length)
//		}
//	}
//	if currentSession+1 >= *minimumSession {
//		return
//	}
//	if notFound {
//		findMinimumSession(tasks, minimumSession, 0, currentSession+1, sessionTime, count, length)
//	}
//
//}

func main() {
	start := time.Now()
	fmt.Println(minSessions([]int{1, 2, 3}, 3))
	fmt.Println((time.Now().Sub(start)).Seconds())
}
