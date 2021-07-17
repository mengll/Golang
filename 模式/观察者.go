package main

import (
	"fmt"
	"time"
)

type Event struct {
	Data string
}

type Observer interface {
	Update(*Event)
}

// b eo 被观察对象
type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify(*Event)
}

type ConcreteObserver struct {
	Id int
}

func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg :%s\n", co.Id, e.Data)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Register(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

func (cs *ConcreteSubject) Deregister(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}

	observer1 := &ConcreteObserver{1}
	observer2 := &ConcreteObserver{2}
	cs.Register(observer1)
	cs.Register(observer2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
