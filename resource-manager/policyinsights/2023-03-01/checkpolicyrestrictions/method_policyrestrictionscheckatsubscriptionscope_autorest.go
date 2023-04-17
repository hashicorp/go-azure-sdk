package checkpolicyrestrictions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRestrictionsCheckAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckRestrictionsResult
}

// PolicyRestrictionsCheckAtSubscriptionScope ...
func (c CheckPolicyRestrictionsClient) PolicyRestrictionsCheckAtSubscriptionScope(ctx context.Context, id commonids.SubscriptionId, input CheckRestrictionsRequest) (result PolicyRestrictionsCheckAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForPolicyRestrictionsCheckAtSubscriptionScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPolicyRestrictionsCheckAtSubscriptionScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPolicyRestrictionsCheckAtSubscriptionScope prepares the PolicyRestrictionsCheckAtSubscriptionScope request.
func (c CheckPolicyRestrictionsClient) preparerForPolicyRestrictionsCheckAtSubscriptionScope(ctx context.Context, id commonids.SubscriptionId, input CheckRestrictionsRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/checkPolicyRestrictions", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPolicyRestrictionsCheckAtSubscriptionScope handles the response to the PolicyRestrictionsCheckAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c CheckPolicyRestrictionsClient) responderForPolicyRestrictionsCheckAtSubscriptionScope(resp *http.Response) (result PolicyRestrictionsCheckAtSubscriptionScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
