package bookrpc

type RemoteBookServer interface {
	Book(args *BookArgs, reply *BookReply) error

	Cancel(args *CancelArgs, reply *CancelReply) error

	Search(args *SearchArgs, reply *SearchReply) error

	CreateTicket(args *CreateTicketArgs, reply *CreateTicketReply) error

	CreateAirline(args *CreateAirlineArgs, reply *CreateAirlineReply) error
}

type BookServer struct {
	RemoteBookServer
}

func Wrap(b RemoteBookServer) RemoteBookServer {
	return &BookServer{b}
}
