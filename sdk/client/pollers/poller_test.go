package pollers_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

// polling then cancelled, polling then failed, polling then succeeded
// timeouts for each

func TestPoller_ImmediatelySucceeded(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusSucceeded,
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

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

func TestPoller_ImmediatelyCancelled_NoDetails(t *testing.T) {
	pollerType := fakePollerWithResults([]pollResult{
		errorResult{
			Error: pollers.PollingCancelledError{
				// no http body or message, presuming the context was cancelled or similar
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(500*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	expectedErrorMessage := "polling was cancelled"
	if err.Error() != expectedErrorMessage {
		t.Fatalf("expected the error to be %q but got %q", expectedErrorMessage, err.Error())
	}
	if poller.LatestStatus() != pollers.PollingStatusCancelled {
		t.Fatalf("expected LatestStatus to be Cancelled but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != nil {
		t.Fatalf("expected LatestResponse to be nil but got: %+v", poller.LatestResponse())
	}
}

func TestPoller_ImmediatelyCancelled_WithDetails(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		errorResult{
			Error: pollers.PollingCancelledError{
				HttpResponse: expectedResponse,
				Message:      "not today!",
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	expectedErrorMessage := "polling was cancelled: not today!"
	if err.Error() != expectedErrorMessage {
		t.Fatalf("expected the error to be %q but got %q", expectedErrorMessage, err.Error())
	}
	if poller.LatestStatus() != pollers.PollingStatusCancelled {
		t.Fatalf("expected LatestStatus to be Cancelled but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != expectedResponse {
		t.Fatalf("expected LatestResponse to be match but it didn't")
	}
}

func TestPoller_ImmediatelyFailed_NoDetails(t *testing.T) {
	pollerType := fakePollerWithResults([]pollResult{
		errorResult{
			Error: pollers.PollingFailedError{
				// no http body or message, presuming API issues or similar
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	expectedErrorMessage := "polling failed"
	if err.Error() != expectedErrorMessage {
		t.Fatalf("expected the error to be %q but got %q", expectedErrorMessage, err.Error())
	}
	if poller.LatestStatus() != pollers.PollingStatusFailed {
		t.Fatalf("expected LatestStatus to be Failed but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != nil {
		t.Fatalf("expected LatestResponse to be nil but got: %+v", poller.LatestResponse())
	}
}

func TestPoller_ImmediatelyFailed_WithDetails(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		errorResult{
			Error: pollers.PollingFailedError{
				HttpResponse: expectedResponse,
				Message:      "not today!",
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	expectedErrorMessage := "polling failed: not today!"
	if err.Error() != expectedErrorMessage {
		t.Fatalf("expected the error to be %q but got %q", expectedErrorMessage, err.Error())
	}
	if poller.LatestStatus() != pollers.PollingStatusFailed {
		t.Fatalf("expected LatestStatus to be Failed but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != expectedResponse {
		t.Fatalf("expected LatestResponse to be match but it didn't")
	}
}

func TestPoller_ImmediatelyShouldRaiseAnErrorWhenTimeoutExceeded(t *testing.T) {
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1500 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusUnknown {
		t.Fatalf("expected LatestStatus to be PollingStatusUnknown but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != nil {
		t.Fatalf("expected LatestResponse to be nil but got: %+v", poller.LatestResponse())
	}
	if pollerType.count != 1 {
		t.Fatalf("expected the fakePoller to be called 1 time but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldPollUntilSucceeded(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		pollers.PollResult{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusSucceeded,
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	if err := poller.PollUntilDone(ctx); err != nil {
		t.Fatalf("when polling: %+v", err)
	}
	if poller.LatestStatus() != pollers.PollingStatusSucceeded {
		t.Fatalf("expected LatestStatus to be PollingStatusSucceeded but got %q", string(poller.LatestStatus()))
	}
	if expectedResponse != poller.LatestResponse() {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
	if pollerType.count != 3 {
		t.Fatalf("expected the fakePoller to be called 3 times but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldContinuePollingWhenConnectionIsDropped_UntilSuccessful(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		pollers.PollResult{
			HttpResponse: expectedResponse,
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusSucceeded,
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	if err := poller.PollUntilDone(ctx); err != nil {
		t.Fatalf("when polling: %+v", err)
	}
	if poller.LatestStatus() != pollers.PollingStatusSucceeded {
		t.Fatalf("expected LatestStatus to be PollingStatusSucceeded but got %q", string(poller.LatestStatus()))
	}
	if expectedResponse != poller.LatestResponse() {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
	if pollerType.count != 7 {
		t.Fatalf("expected the fakePoller to be called 7 times but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldContinuePollingWhenConnectionIsDropped_UntilCancelled(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		errorResult{
			Error: pollers.PollingCancelledError{
				HttpResponse: expectedResponse,
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusCancelled {
		t.Fatalf("expected LatestStatus to be PollingStatusCancelled but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != expectedResponse {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
	if pollerType.count != 7 {
		t.Fatalf("expected the fakePoller to be called 7 times but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldContinuePollingWhenConnectionIsDropped_UntilFailed(t *testing.T) {
	expectedResponse := &client.Response{}
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		errorResult{
			Error: pollers.PollingFailedError{
				HttpResponse: expectedResponse,
			},
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusFailed {
		t.Fatalf("expected LatestStatus to be PollingStatusFailed but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != expectedResponse {
		t.Fatalf("expected LatestResponse to match but it didn't")
	}
	if pollerType.count != 7 {
		t.Fatalf("expected the fakePoller to be called 7 times but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldRaiseAnErrorWhenNumberOfAllowedConnectionDropsIsExceeded(t *testing.T) {
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 1 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		connectionDroppedResult{},
		connectionDroppedResult{},
		connectionDroppedResult{},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusUnknown {
		t.Fatalf("expected LatestStatus to be PollingStatusUnknown but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != nil {
		t.Fatalf("expected LatestResponse to be nil but got: %+v", poller.LatestResponse())
	}
	if pollerType.count != 4 {
		t.Fatalf("expected the fakePoller to be called 4 times but got %d", pollerType.count)
	}
}

func TestPoller_InProgress_ShouldRaiseAnErrorWhenTimeoutExceeded(t *testing.T) {
	pollerType := fakePollerWithResults([]pollResult{
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 750 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
		pollers.PollResult{
			HttpResponse: &client.Response{},
			PollInterval: 750 * time.Millisecond,
			Status:       pollers.PollingStatusInProgress,
		},
	})
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if poller.LatestStatus() != pollers.PollingStatusUnknown {
		t.Fatalf("expected LatestStatus to be PollingStatusUnknown but got %q", string(poller.LatestStatus()))
	}
	if poller.LatestResponse() != nil {
		t.Fatalf("expected LatestResponse to be nil but got: %+v", poller.LatestResponse())
	}
	if pollerType.count != 2 {
		t.Fatalf("expected the fakePoller to be called 2 times but got %d", pollerType.count)
	}
}

type pollResult interface {
}

var _ pollResult = pollers.PollResult{}

var _ pollResult = connectionDroppedResult{}

type connectionDroppedResult struct {
	// we don't include an error here since PollingDroppedConnectionError is also available
	// however we don't generally treat dropped connections as errors (unless there's a sequence of them)
}

var _ pollResult = errorResult{}
var _ pollers.PollerType = &fakePoller{}

type errorResult struct {
	Error error
}

type fakePoller struct {
	count     int
	responses []pollResult
}

func fakePollerWithResults(results []pollResult) *fakePoller {
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

	// a dropped connection doesn't necessarily have to have an error, so this type forces this to fail with no resp/error
	if _, ok := result.(connectionDroppedResult); ok {
		return nil, nil
	}

	if errorResult, ok := result.(errorResult); ok {
		return nil, errorResult.Error
	}

	pollResult := result.(pollers.PollResult)
	return &pollResult, nil
}
