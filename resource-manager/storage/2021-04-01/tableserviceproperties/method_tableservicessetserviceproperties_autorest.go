package tableserviceproperties

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type TableServicesSetServicePropertiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *TableServiceProperties
}

// TableServicesSetServiceProperties ...
func (c TableServicePropertiesClient) TableServicesSetServiceProperties(ctx context.Context, id TableServiceId, input TableServiceProperties) (result TableServicesSetServicePropertiesOperationResponse, err error) {
	req, err := c.preparerForTableServicesSetServiceProperties(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tableserviceproperties.TableServicePropertiesClient", "TableServicesSetServiceProperties", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tableserviceproperties.TableServicePropertiesClient", "TableServicesSetServiceProperties", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTableServicesSetServiceProperties(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tableserviceproperties.TableServicePropertiesClient", "TableServicesSetServiceProperties", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTableServicesSetServiceProperties prepares the TableServicesSetServiceProperties request.
func (c TableServicePropertiesClient) preparerForTableServicesSetServiceProperties(ctx context.Context, id TableServiceId, input TableServiceProperties) (*http.Request, error) {
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

// responderForTableServicesSetServiceProperties handles the response to the TableServicesSetServiceProperties request. The method always
// closes the http.Response Body.
func (c TableServicePropertiesClient) responderForTableServicesSetServiceProperties(resp *http.Response) (result TableServicesSetServicePropertiesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
