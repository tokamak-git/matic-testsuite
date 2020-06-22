package runner

type Action interface {
	Exec() error
}

type Waitable interface {
	Wait()
}

type Restful interface {
	Rest()
}
