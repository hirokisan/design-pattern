package observer

import (
	"errors"
)

var (
	ErrNofityObserver = errors.New("failed to notify observer")
)

type Observer interface {
	update(Subject) error
}

type observer struct {
}

func (o *observer) update(subject Subject) error {
	if subject.Name() != "subject" {
		return ErrNofityObserver
	}

	return nil
}

func NewObserver() *observer {
	return &observer{}
}

type Subject interface {
	Name() string
	SetObserver(Observer)
	notifyObserver() error
}

type subject struct {
	name     string
	observer Observer
}

func NewSubject(name string) *subject {
	return &subject{
		name: name,
	}
}

func (s *subject) Name() string {
	return s.name
}

func (s *subject) SetObserver(observer Observer) {
	s.observer = observer
}

func (s *subject) ChangeName(name string) error {
	s.name = name

	var err error
	{
		err = s.notifyObserver()
	}

	return err
}

func (s *subject) notifyObserver() error {
	return s.observer.update(s)
}
