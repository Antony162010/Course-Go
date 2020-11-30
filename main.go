package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) addTask(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) deleteTask(i int) {
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}

func (tl *taskList) showTasks() {
	for _, task := range tl.tasks {
		fmt.Printf("%+v\n", *task)
	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateCompleted() {
	t.completed = !t.completed
}

func main() {
	t1 := &task{
		name:        "Antony",
		description: "Best developer",
		completed:   true,
	}
	t2 := &task{
		name:        "Christian",
		description: "Best engineer",
	}

	tl := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	t3 := &task{
		name:        "Diego",
		description: "Best frontend",
		completed:   false,
	}
	tl.addTask(t3)
	fmt.Println(tl.tasks)

	tl.deleteTask(1)
	fmt.Println(tl.tasks)

	tl.showTasks()

	mapTasks := make(map[string]*taskList)

	mapTasks["Friends"] = tl
	mapTasks["Friends"].showTasks()
}
