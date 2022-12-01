package pollers_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

// dropped connection
// polling then cancelled, polling then failed, polling then succeeded
// timeouts for each

func TestPoller_ImmediatelySucceeded(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollers.PollResult{
		{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusSucceeded,
		},
	})
	poller := pollers.NewPoller(pollerType, 1*time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	if err := poller.PollUntilDone(ctx); err != nil {
		t.Fatalf("polling: %+v", err)
	}
	if poller.LatestStatus() != pollers.PollingStatusSucceeded {
		t.Fatalf("expected LatestStatus to be Succeeded but got %q", string(poller.LatestStatus()))
	}
	if expectedResponse != poller.LatestResponse() {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
}

func TestPoller_ImmediatelyCancelled_ShouldReturnAnErrorIfNoError(t *testing.T) {
	// the PollingStatus of Cancelled should always be returned as an error and include more information
	// about what happened
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollers.PollResult{
		{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusCancelled,
		},
	})
	poller := pollers.NewPoller(pollerType, 1*time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error when polling but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusCancelled {
		t.Fatalf("expected LatestStatus to be Failed but got %q", string(poller.LatestStatus()))
	}
	if expectedResponse != poller.LatestResponse() {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
}

func TestPoller_ImmediatelyFailed_ShouldReturnAnErrorIfNoError(t *testing.T) {
	// the PollingStatus of Failed should always be returned as an error and include more information
	// about what happened
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollers.PollResult{
		{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusFailed,
		},
	})
	poller := pollers.NewPoller(pollerType, 1*time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error when polling but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusFailed {
		t.Fatalf("expected LatestStatus to be Failed but got %q", string(poller.LatestStatus()))
	}
	if expectedResponse != poller.LatestResponse() {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
}

var _ pollers.PollerType = &fakePoller{}

type fakePoller struct {
	count     int
	responses []pollers.PollResult
}

func fakePollerWithResults(results []pollers.PollResult) *fakePoller {
	return &fakePoller{
		count:     0,
		responses: results,
	}
}

func (f *fakePoller) Poll(ctx context.Context) (*pollers.PollResult, error) {
	f.count++

	if _, ok := ctx.Deadline(); !ok {
		return nil, fmt.Errorf("context must have a deadline but didn't get one")
	}

	if f.count > len(f.responses) {
		return nil, fmt.Errorf("out of responses")
	}

	result := f.responses[f.count-1]
	return &result, nil
}
