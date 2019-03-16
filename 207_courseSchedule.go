func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return false
	}

	if prerequisites == nil || len(prerequisites) <= 1 {
		return true
	}

	inDegree := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		tmp := prerequisites[i]
		inDegree[tmp[0]]++
	}

	ans := []int{}
	queue := make([]int, 0)

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		ans = append(ans, course)

		for i := 0; i < len(prerequisites); i++ {
			tmp := prerequisites[i]
			if tmp[1] == course {
				inDegree[tmp[0]]--
				if inDegree[tmp[0]] == 0 {
					queue = append(queue, tmp[0])
				}
			}
		}
	}

	if len(ans) == numCourses {
		return true
	}
	return false
}