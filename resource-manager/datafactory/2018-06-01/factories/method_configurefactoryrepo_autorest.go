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

type ConfigureFactoryRepoOperationResponse struct {
	HttpResponse *http.Response
	Model        *Factory
}

// ConfigureFactoryRepo ...
func (c FactoriesClient) ConfigureFactoryRepo(ctx context.Context, id LocationId, input FactoryRepoUpdate) (result ConfigureFactoryRepoOperationResponse, err error) {
	req, err := c.preparerForConfigureFactoryRepo(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "ConfigureFactoryRepo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "ConfigureFactoryRepo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForConfigureFactoryRepo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "ConfigureFactoryRepo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForConfigureFactoryRepo prepares the ConfigureFactoryRepo request.
func (c FactoriesClient) preparerForConfigureFactoryRepo(ctx context.Context, id LocationId, input FactoryRepoUpdate) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configureFactoryRepo", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForConfigureFactoryRepo handles the response to the ConfigureFactoryRepo request. The method always
// closes the http.Response Body.
func (c FactoriesClient) responderForConfigureFactoryRepo(resp *http.Response) (result ConfigureFactoryRepoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
