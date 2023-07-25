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

type IsCloneableOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteCloneability
}

// IsCloneable ...
func (c WebAppsClient) IsCloneable(ctx context.Context, id SiteId) (result IsCloneableOperationResponse, err error) {
	req, err := c.preparerForIsCloneable(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneable", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneable", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIsCloneable(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "IsCloneable", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIsCloneable prepares the IsCloneable request.
func (c WebAppsClient) preparerForIsCloneable(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/iscloneable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIsCloneable handles the response to the IsCloneable request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForIsCloneable(resp *http.Response) (result IsCloneableOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
