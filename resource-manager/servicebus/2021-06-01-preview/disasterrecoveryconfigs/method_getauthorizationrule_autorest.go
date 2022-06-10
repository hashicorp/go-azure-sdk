package disasterrecoveryconfigs

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetAuthorizationRuleOperationResponse struct {
	HttpResponse *http.Response
	Model        *SBAuthorizationRule
}

// GetAuthorizationRule ...
func (c DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, id DisasterRecoveryConfigAuthorizationRuleId) (result GetAuthorizationRuleOperationResponse, err error) {
	req, err := c.preparerForGetAuthorizationRule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disasterrecoveryconfigs.DisasterRecoveryConfigsClient", "GetAuthorizationRule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "disasterrecoveryconfigs.DisasterRecoveryConfigsClient", "GetAuthorizationRule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthorizationRule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "disasterrecoveryconfigs.DisasterRecoveryConfigsClient", "GetAuthorizationRule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthorizationRule prepares the GetAuthorizationRule request.
func (c DisasterRecoveryConfigsClient) preparerForGetAuthorizationRule(ctx context.Context, id DisasterRecoveryConfigAuthorizationRuleId) (*http.Request, error) {
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

// responderForGetAuthorizationRule handles the response to the GetAuthorizationRule request. The method always
// closes the http.Response Body.
func (c DisasterRecoveryConfigsClient) responderForGetAuthorizationRule(resp *http.Response) (result GetAuthorizationRuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
