package repository

type Log interface {
	Write(o ...interface{})
}
