package urx

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UrDashboardClient  The Dashboard client
type UrDashboardClient struct {
	Addr string
	Port string
	//conn net.Conn

	incoming chan string
	outgoing chan string

	reader *bufio.Reader
	writer *bufio.Writer
}

func (d *UrDashboardClient) sendCmd(cmd string) {
	d.outgoing <- cmd
}

func (d *UrDashboardClient) Quit() {
	d.sendCmd("quit")
}

func NewUrDashboardClient(connection net.Conn) *UrDashboardClient {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &UrDashboardClient{
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
	}

	client.Listen()
	return client
}

func (client *UrDashboardClient) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
	}
}

func (client *UrDashboardClient) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func (d *UrDashboardClient) Listen() {
	go d.Read()
	go d.Write()
}

func (d *UrDashboardClient) Output() {
	for {
		for line := range d.incoming {
			fmt.Println(line)
		}
	}
}

func (d *UrDashboardClient) Input() {
	in := bufio.NewReader(os.Stdin)
	for {
		text, _ := in.ReadString('\n')
		d.sendCmd(text)
	}
}
