package appservicecertificateorders

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

type CreateOrUpdateCertificateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateCertificate ...
func (c AppServiceCertificateOrdersClient) CreateOrUpdateCertificate(ctx context.Context, id CertificateId, input AppServiceCertificateResource) (result CreateOrUpdateCertificateOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateCertificate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "CreateOrUpdateCertificate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateCertificate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "CreateOrUpdateCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateCertificateThenPoll performs CreateOrUpdateCertificate then polls until it's completed
func (c AppServiceCertificateOrdersClient) CreateOrUpdateCertificateThenPoll(ctx context.Context, id CertificateId, input AppServiceCertificateResource) error {
	result, err := c.CreateOrUpdateCertificate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateCertificate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateCertificate: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateCertificate prepares the CreateOrUpdateCertificate request.
func (c AppServiceCertificateOrdersClient) preparerForCreateOrUpdateCertificate(ctx context.Context, id CertificateId, input AppServiceCertificateResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdateCertificate sends the CreateOrUpdateCertificate request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceCertificateOrdersClient) senderForCreateOrUpdateCertificate(ctx context.Context, req *http.Request) (future CreateOrUpdateCertificateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
