package batchdeployment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type CreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdate ...
func (c BatchDeploymentClient) CreateOrUpdate(ctx context.Context, id DeploymentId, input BatchDeploymentTrackedResource) (result CreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchdeployment.BatchDeploymentClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchdeployment.BatchDeploymentClient", "CreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateThenPoll performs CreateOrUpdate then polls until it's completed
func (c BatchDeploymentClient) CreateOrUpdateThenPoll(ctx context.Context, id DeploymentId, input BatchDeploymentTrackedResource) error {
	result, err := c.CreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdate prepares the CreateOrUpdate request.
func (c BatchDeploymentClient) preparerForCreateOrUpdate(ctx context.Context, id DeploymentId, input BatchDeploymentTrackedResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdate sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c BatchDeploymentClient) senderForCreateOrUpdate(ctx context.Context, req *http.Request) (future CreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
