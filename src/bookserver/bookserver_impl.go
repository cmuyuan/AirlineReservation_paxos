package bookserver

import (
	"errors"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"rpc/bookrpc"
	"sort"
	"strings"
	"sync"
	"time"
)

type bookserver struct {
	myHostPort
	serverList []Node
	servermap  map[uint32]string
	srvConns   map[uint32]*rpc.Client
	mutex      *sync.RWMutex

	epochChan chan struct{}
	closeChan chan struct{}
}

func NewBookServer(masterServerHostPort, myHostPort string) (BookServer, error) {

}

func (bs *bookserver) Book(args *BookArgs, reply *BookReply) error {

}

func (bs *bookserver) Cancel(args *CancelArgs, reply *CancelReply) error {

}

func (bs *bookserver) Search(args *SearchArgs, reply *SearchReply) error {

}

func (bs *bookserver) CreateTicket(args *CreateTicketArgs, reply *CreateTicketReply) error {

}

func (bs *bookserver) CreateAirline(args *CreateAirlineArgs, reply *CreateAirlineReply) error {

}
