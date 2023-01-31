package supportpackages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerSupportPackageOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// TriggerSupportPackage ...
func (c SupportPackagesClient) TriggerSupportPackage(ctx context.Context, id DataBoxEdgeDeviceId, input TriggerSupportPackageRequest) (result TriggerSupportPackageOperationResponse, err error) {
	req, err := c.preparerForTriggerSupportPackage(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "supportpackages.SupportPackagesClient", "TriggerSupportPackage", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTriggerSupportPackage(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "supportpackages.SupportPackagesClient", "TriggerSupportPackage", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TriggerSupportPackageThenPoll performs TriggerSupportPackage then polls until it's completed
func (c SupportPackagesClient) TriggerSupportPackageThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input TriggerSupportPackageRequest) error {
	result, err := c.TriggerSupportPackage(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing TriggerSupportPackage: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after TriggerSupportPackage: %+v", err)
	}

	return nil
}

// preparerForTriggerSupportPackage prepares the TriggerSupportPackage request.
func (c SupportPackagesClient) preparerForTriggerSupportPackage(ctx context.Context, id DataBoxEdgeDeviceId, input TriggerSupportPackageRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/triggerSupportPackage", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTriggerSupportPackage sends the TriggerSupportPackage request. The method will close the
// http.Response Body if it receives an error.
func (c SupportPackagesClient) senderForTriggerSupportPackage(ctx context.Context, req *http.Request) (future TriggerSupportPackageOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
