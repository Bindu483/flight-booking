package flight_booking

import "time"

type Booking struct {
	DepartureTime  time.Time
	ArrivalTime    time.Time
	DeparturePlace string
	ArrivalPlace   string
	CustomerInfo   *Customer
	BookingId      string
	FlightInfo     *Flight
	Status bool
}
type Customer struct {
	CustomerName   string
	PhoneNumber    string
	PassportNumber string
	DateOfBirth    time.Time
}
type BookingStatus struct {
	SeatNumber string
	BookingId  string
	Status bool
}
type Flight struct {
	FlightNumber string
	CarrierName  string
}
