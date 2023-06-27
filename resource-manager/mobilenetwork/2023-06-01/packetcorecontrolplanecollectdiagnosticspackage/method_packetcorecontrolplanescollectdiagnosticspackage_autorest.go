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

type PacketCoreControlPlanesCollectDiagnosticsPackageOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PacketCoreControlPlanesCollectDiagnosticsPackage ...
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) PacketCoreControlPlanesCollectDiagnosticsPackage(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) (result PacketCoreControlPlanesCollectDiagnosticsPackageOperationResponse, err error) {
	req, err := c.preparerForPacketCoreControlPlanesCollectDiagnosticsPackage(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient", "PacketCoreControlPlanesCollectDiagnosticsPackage", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPacketCoreControlPlanesCollectDiagnosticsPackage(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanecollectdiagnosticspackage.PacketCoreControlPlaneCollectDiagnosticsPackageClient", "PacketCoreControlPlanesCollectDiagnosticsPackage", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PacketCoreControlPlanesCollectDiagnosticsPackageThenPoll performs PacketCoreControlPlanesCollectDiagnosticsPackage then polls until it's completed
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) PacketCoreControlPlanesCollectDiagnosticsPackageThenPoll(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) error {
	result, err := c.PacketCoreControlPlanesCollectDiagnosticsPackage(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing PacketCoreControlPlanesCollectDiagnosticsPackage: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PacketCoreControlPlanesCollectDiagnosticsPackage: %+v", err)
	}

	return nil
}

// preparerForPacketCoreControlPlanesCollectDiagnosticsPackage prepares the PacketCoreControlPlanesCollectDiagnosticsPackage request.
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) preparerForPacketCoreControlPlanesCollectDiagnosticsPackage(ctx context.Context, id PacketCoreControlPlaneId, input PacketCoreControlPlaneCollectDiagnosticsPackage) (*http.Request, error) {
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

// senderForPacketCoreControlPlanesCollectDiagnosticsPackage sends the PacketCoreControlPlanesCollectDiagnosticsPackage request. The method will close the
// http.Response Body if it receives an error.
func (c PacketCoreControlPlaneCollectDiagnosticsPackageClient) senderForPacketCoreControlPlanesCollectDiagnosticsPackage(ctx context.Context, req *http.Request) (future PacketCoreControlPlanesCollectDiagnosticsPackageOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
