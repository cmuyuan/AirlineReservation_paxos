package bookrpc

type Status int

const (
	OK Status = iota + 1
	NoSuchUser
	NoSuchFlight
)

type Flight struct {
	From    string
	To      string
	Airline string
}

type BookArgs struct {
	Flight Flight
	UserID string
}

type BookReply struct {
	Status Status
}

type CancelArgs struct {
	Flight Flight
	UserID string
}

type CancelReply struct {
	Status Status
}

type SearchArgs struct {
	From string
	To   string
}

type SearchReply struct {
	Status Status
}

type CreateTicketArgs struct {
	Flight Flight
}

type CreateTicketReply struct {
	Status Status
}

type CreateAirlineArgs struct {
	Flight Flight
}

type CreateAirlineReply struct {
	Status Status
}
