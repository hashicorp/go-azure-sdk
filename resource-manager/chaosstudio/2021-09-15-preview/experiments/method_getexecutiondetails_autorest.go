package experiments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetExecutionDetailsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExperimentExecutionDetails
}

// GetExecutionDetails ...
func (c ExperimentsClient) GetExecutionDetails(ctx context.Context, id ExecutionDetailId) (result GetExecutionDetailsOperationResponse, err error) {
	req, err := c.preparerForGetExecutionDetails(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "GetExecutionDetails", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "GetExecutionDetails", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetExecutionDetails(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "GetExecutionDetails", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetExecutionDetails prepares the GetExecutionDetails request.
func (c ExperimentsClient) preparerForGetExecutionDetails(ctx context.Context, id ExecutionDetailId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetExecutionDetails handles the response to the GetExecutionDetails request. The method always
// closes the http.Response Body.
func (c ExperimentsClient) responderForGetExecutionDetails(resp *http.Response) (result GetExecutionDetailsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
