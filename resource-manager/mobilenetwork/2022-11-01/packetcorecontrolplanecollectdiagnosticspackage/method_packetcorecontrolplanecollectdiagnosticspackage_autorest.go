package packetcorecontrolplanecollectdiagnosticspackage

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

type PacketCoreControlPlaneCollectDiagnosticsPackageOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PacketCoreControlPlaneCollectDiagnosticsPackage ...
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) PacketCoreControlPlaneCollectDiagnosticsPackage(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) (result PacketCoreControlPlaneCollectDiagnosticsPackageOperationResponse, err error) {
	req, err := c.preparerForPacketCoreControlPlaneCollectDiagnosticsPackage(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient", "PacketCoreControlPlaneCollectDiagnosticsPackage", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPacketCoreControlPlaneCollectDiagnosticsPackage(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient", "PacketCoreControlPlaneCollectDiagnosticsPackage", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PacketCoreControlPlaneCollectDiagnosticsPackageThenPoll performs PacketCoreControlPlaneCollectDiagnosticsPackage then polls until it's completed
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) PacketCoreControlPlaneCollectDiagnosticsPackageThenPoll(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) error {
	result, err := c.PacketCoreControlPlaneCollectDiagnosticsPackage(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PacketCoreControlPlaneCollectDiagnosticsPackage: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PacketCoreControlPlaneCollectDiagnosticsPackage: %+v", err)
	}

	return nil
}

// preparerForPacketCoreControlPlaneCollectDiagnosticsPackage prepares the PacketCoreControlPlaneCollectDiagnosticsPackage request.
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) preparerForPacketCoreControlPlaneCollectDiagnosticsPackage(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/collectDiagnosticsPackage", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForPacketCoreControlPlaneCollectDiagnosticsPackage sends the PacketCoreControlPlaneCollectDiagnosticsPackage request. The method will close the
// http.Response Body if it receives an error.
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) senderForPacketCoreControlPlaneCollectDiagnosticsPackage(ctx context.Context, req *http.Request) (future PacketCoreControlPlaneCollectDiagnosticsPackageOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
