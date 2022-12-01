package pollers

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/helpers"
	"time"
)

type Poller struct {
	// initialDelayDuration specifies the duration of the initial delay when polling
	// this is also used for retries should a `latestResponse` not be available, for
	// example when a connection is dropped.
	initialDelayDuration time.Duration

	// latestResponse contains the polling status from the latest response.
	latestResponse *PollResult

	// poller is a reference to the PollerType, for example a LongRunningOperationPoller
	// which should be polled to determine the latest state.
	poller PollerType
}

// NewPoller returns a Poller based on the given PollerType.
func NewPoller(pollerType PollerType, initialDelayDuration time.Duration) Poller {
	return Poller{
		initialDelayDuration: initialDelayDuration,
		poller:               pollerType,
	}
}

// LatestResponse returns the latest HTTP Response returned when polling
func (p *Poller) LatestResponse() *client.Response {
	if p.latestResponse == nil {
		return nil
	}

	return p.latestResponse.HttpResponse
}

// LatestStatus returns the latest status returned when polling
func (p *Poller) LatestStatus() PollingStatus {
	if p.latestResponse == nil {
		return PollingStatusUnknown
	}

	return p.latestResponse.Status
}

// PollUntilDone polls until the poller determines that the operation has been completed
func (p *Poller) PollUntilDone(ctx context.Context) error {
	if p.poller == nil {
		return fmt.Errorf("internal-error: `poller` was nil`")
	}

	endTime := time.Now().Add(p.initialDelayDuration)
	helpers.SleepUntil(time.Until(endTime), ctx.Done())

	for true {
		// determine the next retry duration / how long to poll for
		retryDuration := p.initialDelayDuration
		if p.latestResponse != nil {
			retryDuration = p.latestResponse.PollInterval
		}
		endTime := time.Now().Add(retryDuration)
		helpers.SleepUntil(time.Until(endTime), ctx.Done())

		var err error
		p.latestResponse, err = p.poller.Poll(ctx)
		if err != nil {
			return fmt.Errorf("polling: %+v", err)
		}

		if resp := p.latestResponse; resp != nil {
			retryDuration = p.latestResponse.PollInterval
			switch p.latestResponse.Status {
			// both Cancelled and Failed should be raised as errors
			// containing additional information as available

			case PollingStatusInProgress:
				break

			case PollingStatusSucceeded:
				return nil

			default:
				return fmt.Errorf("internal-error: unimplemented polling status %q", string(p.latestResponse.Status))
			}
		}
	}

	return nil
}
