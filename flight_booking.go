package flight_booking

type FlightBooking interface {
	BookFlight(b *Booking) (*BookingStatus,error)
	CancelFlight(bookingId string) (*BookingStatus,error)
}
