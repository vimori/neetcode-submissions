func canFinish(numCourses int, prerequisites [][]int) bool {
    inStack := make(map[int][]int)

	for _, val := range prerequisites{
		inStack[val[0]] = append(inStack[val[0]], val[1])
	}

	visited := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(cur int) bool{
		if visited[cur] { //обнаружен цикл
			return false
		}
		if len(inStack[cur]) == 0 { //курс можно пройти сразу 
			return true
		}

		visited[cur] = true //помечаем, что курс обрабатывается 
		for _, val := range inStack[cur]{
			if !dfs(val) {
				return false
			}
		}
		visited[cur] = false //обработка курса закончена
		inStack[cur] = []int{} //отмечаем, что курс больше не обрабатывается
		//(избавляемся от лишнего запуска цикла) 
		return true //курс можно пройти 
	}
	for course := 0; course < numCourses; course++{
		if !dfs(course){
			return false
		}
	}
	return true
}
