package packetcorecontrolplanereinstall

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

type PacketCoreControlPlaneReinstallOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PacketCoreControlPlaneReinstall ...
func (c PacketCoreControlPlaneReinstallClient) PacketCoreControlPlaneReinstall(ctx context.Context, id PacketCoreControlPlaneId) (result PacketCoreControlPlaneReinstallOperationResponse, err error) {
	req, err := c.preparerForPacketCoreControlPlaneReinstall(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanereinstall.PacketCoreControlPlaneReinstallClient", "PacketCoreControlPlaneReinstall", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPacketCoreControlPlaneReinstall(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcorecontrolplanereinstall.PacketCoreControlPlaneReinstallClient", "PacketCoreControlPlaneReinstall", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PacketCoreControlPlaneReinstallThenPoll performs PacketCoreControlPlaneReinstall then polls until it's completed
func (c PacketCoreControlPlaneReinstallClient) PacketCoreControlPlaneReinstallThenPoll(ctx context.Context, id PacketCoreControlPlaneId) error {
	result, err := c.PacketCoreControlPlaneReinstall(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PacketCoreControlPlaneReinstall: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PacketCoreControlPlaneReinstall: %+v", err)
	}

	return nil
}

// preparerForPacketCoreControlPlaneReinstall prepares the PacketCoreControlPlaneReinstall request.
func (c PacketCoreControlPlaneReinstallClient) preparerForPacketCoreControlPlaneReinstall(ctx context.Context, id PacketCoreControlPlaneId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reinstall", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForPacketCoreControlPlaneReinstall sends the PacketCoreControlPlaneReinstall request. The method will close the
// http.Response Body if it receives an error.
func (c PacketCoreControlPlaneReinstallClient) senderForPacketCoreControlPlaneReinstall(ctx context.Context, req *http.Request) (future PacketCoreControlPlaneReinstallOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
