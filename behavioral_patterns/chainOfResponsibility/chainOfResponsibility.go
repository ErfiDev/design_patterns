package main

import "fmt"

/*
Chain of Responsibility is behavioral design pattern
that allows passing request along the chain of potential
 handlers until one of them handles request.

Let’s look at the Chain of Responsibility pattern with the case of a
hospital app. A hospital could have multiple departments such as:

    Reception
    Doctor
    Medical examination room
    Cashier

Whenever any patient arrives, they first get to Reception, then to Doctor,
then to Medicine Room, and then to Cashier (and so on). The patient is
being sent through a chain of departments, where each department sends the
patient further down the chain once their function is completed.
*/

type Department interface {
	setNext(Department)
	execute(*Patient)
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("patient registration done")
		r.next.execute(p)
		return
	}

	fmt.Println("reception registering patient!")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(d Department) {
	r.next = d
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("doctor checkup completed!")
		d.next.execute(p)
		return
	}

	fmt.Println("doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(dep Department) {
	d.next = dep
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicalDone {
		fmt.Println("medicine already given to patient")
		m.next.execute(p)
		return
	}

	fmt.Println("medical giving medicine to the patient")
	p.doctorCheckUpDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(d Department) {
	m.next = d
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicalDone       bool
	paymentDone       bool
}

func main() {
	// initializing departments
	reception := new(Reception)
	doctor := new(Doctor)
	medical := new(Medical)
	cashier := new(Cashier)

	medical.setNext(cashier)
	doctor.setNext(medical)
	reception.setNext(doctor)

	patient := &Patient{
		name: "Erfan",
	}

	reception.execute(patient)
}
