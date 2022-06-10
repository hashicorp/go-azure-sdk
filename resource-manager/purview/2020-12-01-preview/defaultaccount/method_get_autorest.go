package defaultaccount

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DefaultAccountPayload
}

type GetOperationOptions struct {
	Scope         *string
	ScopeTenantId *string
	ScopeType     *ScopeType
}

func DefaultGetOperationOptions() GetOperationOptions {
	return GetOperationOptions{}
}

func (o GetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Scope != nil {
		out["scope"] = *o.Scope
	}

	if o.ScopeTenantId != nil {
		out["scopeTenantId"] = *o.ScopeTenantId
	}

	if o.ScopeType != nil {
		out["scopeType"] = *o.ScopeType
	}

	return out
}

// Get ...
func (c DefaultAccountClient) Get(ctx context.Context, options GetOperationOptions) (result GetOperationResponse, err error) {
	req, err := c.preparerForGet(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Get", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Get", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Get", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGet prepares the Get request.
func (c DefaultAccountClient) preparerForGet(ctx context.Context, options GetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath("/providers/Microsoft.Purview/getDefaultAccount"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGet handles the response to the Get request. The method always
// closes the http.Response Body.
func (c DefaultAccountClient) responderForGet(resp *http.Response) (result GetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
