package experiments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type CancelOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Cancel ...
func (c ExperimentsClient) Cancel(ctx context.Context, id ExperimentId) (result CancelOperationResponse, err error) {
	req, err := c.preparerForCancel(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCancel(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "Cancel", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CancelThenPoll performs Cancel then polls until it's completed
func (c ExperimentsClient) CancelThenPoll(ctx context.Context, id ExperimentId) error {
	result, err := c.Cancel(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Cancel: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Cancel: %+v", err)
	}

	return nil
}

// preparerForCancel prepares the Cancel request.
func (c ExperimentsClient) preparerForCancel(ctx context.Context, id ExperimentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cancel", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCancel sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (c ExperimentsClient) senderForCancel(ctx context.Context, req *http.Request) (future CancelOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
