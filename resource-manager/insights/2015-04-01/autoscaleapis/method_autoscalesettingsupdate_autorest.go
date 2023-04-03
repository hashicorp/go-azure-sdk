package autoscaleapis

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoscaleSettingsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AutoscaleSettingResource
}

// AutoscaleSettingsUpdate ...
func (c AutoscaleAPIsClient) AutoscaleSettingsUpdate(ctx context.Context, id AutoScaleSettingId, input AutoscaleSettingResourcePatch) (result AutoscaleSettingsUpdateOperationResponse, err error) {
	req, err := c.preparerForAutoscaleSettingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "autoscaleapis.AutoscaleAPIsClient", "AutoscaleSettingsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "autoscaleapis.AutoscaleAPIsClient", "AutoscaleSettingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAutoscaleSettingsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "autoscaleapis.AutoscaleAPIsClient", "AutoscaleSettingsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAutoscaleSettingsUpdate prepares the AutoscaleSettingsUpdate request.
func (c AutoscaleAPIsClient) preparerForAutoscaleSettingsUpdate(ctx context.Context, id AutoScaleSettingId, input AutoscaleSettingResourcePatch) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAutoscaleSettingsUpdate handles the response to the AutoscaleSettingsUpdate request. The method always
// closes the http.Response Body.
func (c AutoscaleAPIsClient) responderForAutoscaleSettingsUpdate(resp *http.Response) (result AutoscaleSettingsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
