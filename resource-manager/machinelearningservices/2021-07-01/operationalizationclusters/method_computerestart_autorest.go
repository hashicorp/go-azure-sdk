package operationalizationclusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type ComputeRestartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ComputeRestart ...
func (c OperationalizationClustersClient) ComputeRestart(ctx context.Context, id ComputeId) (result ComputeRestartOperationResponse, err error) {
	req, err := c.preparerForComputeRestart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalizationclusters.OperationalizationClustersClient", "ComputeRestart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForComputeRestart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalizationclusters.OperationalizationClustersClient", "ComputeRestart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ComputeRestartThenPoll performs ComputeRestart then polls until it's completed
func (c OperationalizationClustersClient) ComputeRestartThenPoll(ctx context.Context, id ComputeId) error {
	result, err := c.ComputeRestart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ComputeRestart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ComputeRestart: %+v", err)
	}

	return nil
}

// preparerForComputeRestart prepares the ComputeRestart request.
func (c OperationalizationClustersClient) preparerForComputeRestart(ctx context.Context, id ComputeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForComputeRestart sends the ComputeRestart request. The method will close the
// http.Response Body if it receives an error.
func (c OperationalizationClustersClient) senderForComputeRestart(ctx context.Context, req *http.Request) (future ComputeRestartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
