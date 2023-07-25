package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPublishingProfileXmlWithSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// ListPublishingProfileXmlWithSecrets ...
func (c WebAppsClient) ListPublishingProfileXmlWithSecrets(ctx context.Context, id SiteId, input CsmPublishingProfileOptions) (result ListPublishingProfileXmlWithSecretsOperationResponse, err error) {
	req, err := c.preparerForListPublishingProfileXmlWithSecrets(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingProfileXmlWithSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingProfileXmlWithSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListPublishingProfileXmlWithSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingProfileXmlWithSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListPublishingProfileXmlWithSecrets prepares the ListPublishingProfileXmlWithSecrets request.
func (c WebAppsClient) preparerForListPublishingProfileXmlWithSecrets(ctx context.Context, id SiteId, input CsmPublishingProfileOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/publishxml", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListPublishingProfileXmlWithSecrets handles the response to the ListPublishingProfileXmlWithSecrets request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPublishingProfileXmlWithSecrets(resp *http.Response) (result ListPublishingProfileXmlWithSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
