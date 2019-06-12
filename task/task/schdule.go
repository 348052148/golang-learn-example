package task

type Schedule struct {
	TaskList map[int]*Task
}

func NewSchedule() *Schedule {
	s := new(Schedule)
	s.TaskList = make(map[int]*Task)
	return s
}

func (schedule *Schedule) AddTask(task *Task) {
	schedule.TaskList[task.id] = task
}

func (schedule *Schedule) DelTask(task *Task) {
	delete(schedule.TaskList, task.id)
}

func (schedule *Schedule) Run() {
	for _, task := range schedule.TaskList {
		task.Run()
	}
}
