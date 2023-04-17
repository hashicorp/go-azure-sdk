package checkpolicyrestrictions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRestrictionsCheckAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckRestrictionsResult
}

// PolicyRestrictionsCheckAtManagementGroupScope ...
func (c CheckPolicyRestrictionsClient) PolicyRestrictionsCheckAtManagementGroupScope(ctx context.Context, id ManagementGroupId, input CheckManagementGroupRestrictionsRequest) (result PolicyRestrictionsCheckAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForPolicyRestrictionsCheckAtManagementGroupScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPolicyRestrictionsCheckAtManagementGroupScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkpolicyrestrictions.CheckPolicyRestrictionsClient", "PolicyRestrictionsCheckAtManagementGroupScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPolicyRestrictionsCheckAtManagementGroupScope prepares the PolicyRestrictionsCheckAtManagementGroupScope request.
func (c CheckPolicyRestrictionsClient) preparerForPolicyRestrictionsCheckAtManagementGroupScope(ctx context.Context, id ManagementGroupId, input CheckManagementGroupRestrictionsRequest) (*http.Request, error) {
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

// responderForPolicyRestrictionsCheckAtManagementGroupScope handles the response to the PolicyRestrictionsCheckAtManagementGroupScope request. The method always
// closes the http.Response Body.
func (c CheckPolicyRestrictionsClient) responderForPolicyRestrictionsCheckAtManagementGroupScope(resp *http.Response) (result PolicyRestrictionsCheckAtManagementGroupScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
