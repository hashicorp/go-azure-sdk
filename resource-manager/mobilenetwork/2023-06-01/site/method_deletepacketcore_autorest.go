package site

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

type DeletePacketCoreOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeletePacketCore ...
func (c SiteClient) DeletePacketCore(ctx context.Context, id SiteId, input SiteDeletePacketCore) (result DeletePacketCoreOperationResponse, err error) {
	req, err := c.preparerForDeletePacketCore(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "site.SiteClient", "DeletePacketCore", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeletePacketCore(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "site.SiteClient", "DeletePacketCore", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeletePacketCoreThenPoll performs DeletePacketCore then polls until it's completed
func (c SiteClient) DeletePacketCoreThenPoll(ctx context.Context, id SiteId, input SiteDeletePacketCore) error {
	result, err := c.DeletePacketCore(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DeletePacketCore: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeletePacketCore: %+v", err)
	}

	return nil
}

// preparerForDeletePacketCore prepares the DeletePacketCore request.
func (c SiteClient) preparerForDeletePacketCore(ctx context.Context, id SiteId, input SiteDeletePacketCore) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deletePacketCore", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeletePacketCore sends the DeletePacketCore request. The method will close the
// http.Response Body if it receives an error.
func (c SiteClient) senderForDeletePacketCore(ctx context.Context, req *http.Request) (future DeletePacketCoreOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
