package factories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGitHubAccessTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *GitHubAccessTokenResponse
}

// GetGitHubAccessToken ...
func (c FactoriesClient) GetGitHubAccessToken(ctx context.Context, id FactoryId, input GitHubAccessTokenRequest) (result GetGitHubAccessTokenOperationResponse, err error) {
	req, err := c.preparerForGetGitHubAccessToken(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetGitHubAccessToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetGitHubAccessToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetGitHubAccessToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetGitHubAccessToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetGitHubAccessToken prepares the GetGitHubAccessToken request.
func (c FactoriesClient) preparerForGetGitHubAccessToken(ctx context.Context, id FactoryId, input GitHubAccessTokenRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getGitHubAccessToken", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetGitHubAccessToken handles the response to the GetGitHubAccessToken request. The method always
// closes the http.Response Body.
func (c FactoriesClient) responderForGetGitHubAccessToken(resp *http.Response) (result GetGitHubAccessTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
