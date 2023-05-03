package state

import (
	"fmt"

	"github.com/livepeer/catalyst-api/events"
	v0 "github.com/livepeer/catalyst-api/schema/v0"
)

type SignedEvent struct {
	Action any
}

func HandleEvent(e *events.SignedEvent) error {
	switch act := e.Action.(type) {

	case *v0.ChannelDefinition:
		return nil

	default:
		return fmt.Errorf("unknown action type: %s", act)
	}
}
