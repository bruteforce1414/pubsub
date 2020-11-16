package observer

type subject interface {
	Register(Observer observer)
	Deregister(Observer observer)
	NotifyAll()
}