package volumequotarules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVolumeOperationResponse struct {
	HttpResponse *http.Response
	Model        *VolumeQuotaRulesList
}

// ListByVolume ...
func (c VolumeQuotaRulesClient) ListByVolume(ctx context.Context, id VolumeId) (result ListByVolumeOperationResponse, err error) {
	req, err := c.preparerForListByVolume(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumequotarules.VolumeQuotaRulesClient", "ListByVolume", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumequotarules.VolumeQuotaRulesClient", "ListByVolume", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByVolume(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumequotarules.VolumeQuotaRulesClient", "ListByVolume", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByVolume prepares the ListByVolume request.
func (c VolumeQuotaRulesClient) preparerForListByVolume(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/volumeQuotaRules", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByVolume handles the response to the ListByVolume request. The method always
// closes the http.Response Body.
func (c VolumeQuotaRulesClient) responderForListByVolume(resp *http.Response) (result ListByVolumeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
