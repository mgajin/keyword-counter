package job

import (
	"context"

	"github.com/pkg/errors"
)

// Channel
type Channel interface {
	// Send
	Send(ctx context.Context, j *Job) error

	// Stream
	Stream(ctx context.Context) <-chan *Job
}

// channel implements Channel interface.
type channel struct {
	scanType ScanType
	jobs     chan *Job
}

// NewChannel
func NewChannel(scanType ScanType, buffer int) Channel {
	return &channel{
		scanType: scanType,
		jobs:     make(chan *Job, buffer),
	}
}

func (ch *channel) Send(_ context.Context, j *Job) error {
	if j.ScanType != ch.scanType {
		return errors.Errorf(
			"error: can't send %s job to %s channel",
			j.ScanType, ch.scanType,
		)
	}
	ch.jobs <- j
	return nil
}

func (ch *channel) Stream(_ context.Context) <-chan *Job { return ch.jobs }