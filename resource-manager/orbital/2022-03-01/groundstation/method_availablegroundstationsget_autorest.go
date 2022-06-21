package groundstation

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type AvailableGroundStationsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *AvailableGroundStation
}

// AvailableGroundStationsGet ...
func (c GroundStationClient) AvailableGroundStationsGet(ctx context.Context, id AvailableGroundStationId) (result AvailableGroundStationsGetOperationResponse, err error) {
	req, err := c.preparerForAvailableGroundStationsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "groundstation.GroundStationClient", "AvailableGroundStationsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "groundstation.GroundStationClient", "AvailableGroundStationsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAvailableGroundStationsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "groundstation.GroundStationClient", "AvailableGroundStationsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAvailableGroundStationsGet prepares the AvailableGroundStationsGet request.
func (c GroundStationClient) preparerForAvailableGroundStationsGet(ctx context.Context, id AvailableGroundStationId) (*http.Request, error) {
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

// responderForAvailableGroundStationsGet handles the response to the AvailableGroundStationsGet request. The method always
// closes the http.Response Body.
func (c GroundStationClient) responderForAvailableGroundStationsGet(resp *http.Response) (result AvailableGroundStationsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
