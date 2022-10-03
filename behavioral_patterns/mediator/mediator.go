package main

import "fmt"

/*
Mediator is a behavioral design pattern that reduces
coupling between components of a program by making
them communicate indirectly, through a special mediator
object.

An excellent example of the Mediator pattern is a
railway station traffic system. Two trains never
communicate between themselves for the availability
of the platform. The stationManager acts as a mediator
and makes the platform available to only one of the
arriving trains while keeping the rest in a queue.
A departing train notifies the stations, which lets the
 next train in the queue to arrive.
*/

type Train interface {
	arrive()
	depart()
}

type Manager interface {
	canArrive(t Train) bool
	notifyAboutDeparture()
}

type PassengerTrain struct {
	m Manager
}

func (p *PassengerTrain) arrive() {
	if p.m.canArrive(p) {
		fmt.Println("Passenger Train arrived")
		return
	}

	fmt.Println("Passenger Train arrive blocked!")
}

func (p *PassengerTrain) depart() {
	fmt.Println("Passenger Train depart.")
	p.m.notifyAboutDeparture()
}

type FreightTrain struct {
	m Manager
}

func (p *FreightTrain) arrive() {
	if p.m.canArrive(p) {
		fmt.Println("Freight Train arrived")
		return
	}

	fmt.Println("Freight Train arrive blocked!")
}

func (p *FreightTrain) depart() {
	fmt.Println("Freight Train depart.")
	p.m.notifyAboutDeparture()
}

type StationManager struct {
	stationStatus bool
	queue         []Train
}

func (s *StationManager) canArrive(t Train) bool {
	// true == station is empty
	if s.stationStatus {
		s.stationStatus = false
		return true
	}

	s.queue = append(s.queue, t)

	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.stationStatus {
		s.stationStatus = true
	}

	if len(s.queue) > 0 {
		firstTrain := s.queue[0]
		s.queue = s.queue[1:]
		firstTrain.arrive()
	}
}

func main() {
	sm := StationManager{
		stationStatus: true,
	}

	t1 := PassengerTrain{
		m: &sm,
	}
	t2 := FreightTrain{
		m: &sm,
	}

	t1.arrive()
	t2.arrive()
	t1.depart()
}
