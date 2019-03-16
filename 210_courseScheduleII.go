func findOrder(numCourses int, prerequisites [][]int) []int {
	ans := []int{}

	if numCourses <= 0 {
		return ans
	}
	if prerequisites == nil {
		for i := 0; i < numCourses; i++ {
			ans = append(ans, i)
		}
		return ans
	}
	inDegree := make([]int, numCourses)
	for _, v := range prerequisites {
		inDegree[v[0]]++
	}

	queue := make([]int, 0)

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		ans = append(ans, course)
		queue = queue[1:]
		for _, v := range prerequisites {
			if v[1] == course {
				inDegree[v[0]]--
				if inDegree[v[0]] == 0 {
					queue = append(queue, v[0])
				}
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}
	return nil
}