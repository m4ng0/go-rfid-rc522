package rfid

import "time"

type ReaderChan struct {
	reader RfidReader
	ch     chan string
}

func NewReaderChan(reader RfidReader, sleepDuration time.Duration) (*ReaderChan, error) {
	reducReader, err := NewReducer(reader)
	if err != nil {
		return nil, err
	}
	c := &ReaderChan{
		reader: reducReader,
		ch:     make(chan string, 1),
	}
	go c.loop(sleepDuration)
	return c, nil
}

func (c *ReaderChan) loop(sleepDuration time.Duration) {
	for {
		time.Sleep(sleepDuration)
		id, err := c.reader.ReadId()
		if err != nil {
			continue
		}
		c.ch <- id
	}
}

func (c *ReaderChan) ReadId() (string, error) {
	return <-c.ch, nil
}

func (c *ReaderChan) GetChan() chan string {
	return c.ch
}
