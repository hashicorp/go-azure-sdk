package heatmaps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HeatMapGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *HeatMapModel
}

type HeatMapGetOperationOptions struct {
	BotRight *string
	TopLeft  *string
}

func DefaultHeatMapGetOperationOptions() HeatMapGetOperationOptions {
	return HeatMapGetOperationOptions{}
}

func (o HeatMapGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o HeatMapGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.BotRight != nil {
		out["botRight"] = *o.BotRight
	}

	if o.TopLeft != nil {
		out["topLeft"] = *o.TopLeft
	}

	return out
}

// HeatMapGet ...
func (c HeatMapsClient) HeatMapGet(ctx context.Context, id TrafficManagerProfileId, options HeatMapGetOperationOptions) (result HeatMapGetOperationResponse, err error) {
	req, err := c.preparerForHeatMapGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "heatmaps.HeatMapsClient", "HeatMapGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "heatmaps.HeatMapsClient", "HeatMapGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForHeatMapGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "heatmaps.HeatMapsClient", "HeatMapGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForHeatMapGet prepares the HeatMapGet request.
func (c HeatMapsClient) preparerForHeatMapGet(ctx context.Context, id TrafficManagerProfileId, options HeatMapGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/heatMaps/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForHeatMapGet handles the response to the HeatMapGet request. The method always
// closes the http.Response Body.
func (c HeatMapsClient) responderForHeatMapGet(resp *http.Response) (result HeatMapGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
