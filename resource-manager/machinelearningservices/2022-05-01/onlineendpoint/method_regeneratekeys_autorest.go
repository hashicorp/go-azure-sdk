package onlineendpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type RegenerateKeysOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegenerateKeys ...
func (c OnlineEndpointClient) RegenerateKeys(ctx context.Context, id OnlineEndpointId, input RegenerateEndpointKeysRequest) (result RegenerateKeysOperationResponse, err error) {
	req, err := c.preparerForRegenerateKeys(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onlineendpoint.OnlineEndpointClient", "RegenerateKeys", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegenerateKeys(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onlineendpoint.OnlineEndpointClient", "RegenerateKeys", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegenerateKeysThenPoll performs RegenerateKeys then polls until it's completed
func (c OnlineEndpointClient) RegenerateKeysThenPoll(ctx context.Context, id OnlineEndpointId, input RegenerateEndpointKeysRequest) error {
	result, err := c.RegenerateKeys(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegenerateKeys: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegenerateKeys: %+v", err)
	}

	return nil
}

// preparerForRegenerateKeys prepares the RegenerateKeys request.
func (c OnlineEndpointClient) preparerForRegenerateKeys(ctx context.Context, id OnlineEndpointId, input RegenerateEndpointKeysRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateKeys", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRegenerateKeys sends the RegenerateKeys request. The method will close the
// http.Response Body if it receives an error.
func (c OnlineEndpointClient) senderForRegenerateKeys(ctx context.Context, req *http.Request) (future RegenerateKeysOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
