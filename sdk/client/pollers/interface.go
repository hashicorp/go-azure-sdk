package pollers

import (
	"context"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"time"
)

// TODO: return, populate these and default them
type PollingCancelledError error
type PollingFailedError error

// PollerType allows custom pollers to be created to determine when a particular Operation has
// been Completed, Cancelled or Failed.
type PollerType interface {
	// Poll performs a poll to determine whether the Operation has been Completed/Cancelled or Failed.
	Poll(ctx context.Context) (*PollResult, error)
}

type PollResult struct {
	// HttpResponse is a copy of the HttpResponse returned from the API in the last request.
	HttpResponse *client.Response

	// PollInterval specifies the interval until this poller should be called again.
	PollInterval time.Duration

	// Status specifies the polling status of this resource at the time of the last request.
	Status PollingStatus
}

type PollingStatus string

const (
	// PollingStatusCancelled states that the resource change was cancelled (for example, due to a timeout).
	PollingStatusCancelled PollingStatus = "Cancelled"

	// PollingStatusFailed states that the resource change failed (for example due to a lack of capacity).
	PollingStatusFailed PollingStatus = "Failed"

	// PollingStatusInProgress states that the resource change is still occurring/in-progress.
	PollingStatusInProgress PollingStatus = "InProgress"

	// PollingStatusSucceeded states that the resource change was successful.
	PollingStatusSucceeded PollingStatus = "Succeeded"

	// PollingStatusUnknown states that the resource change state is unknown/unexpected.
	PollingStatusUnknown PollingStatus = "Unknown"
)
