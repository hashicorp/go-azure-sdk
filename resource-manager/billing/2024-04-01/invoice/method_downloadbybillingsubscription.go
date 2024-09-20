package invoice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DownloadByBillingSubscriptionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DocumentDownloadResult
}

type DownloadByBillingSubscriptionOperationOptions struct {
	DocumentName *string
}

func DefaultDownloadByBillingSubscriptionOperationOptions() DownloadByBillingSubscriptionOperationOptions {
	return DownloadByBillingSubscriptionOperationOptions{}
}

func (o DownloadByBillingSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DownloadByBillingSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DownloadByBillingSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DocumentName != nil {
		out.Append("documentName", fmt.Sprintf("%v", *o.DocumentName))
	}
	return &out
}

// DownloadByBillingSubscription ...
func (c InvoiceClient) DownloadByBillingSubscription(ctx context.Context, id BillingSubscriptionInvoiceId, options DownloadByBillingSubscriptionOperationOptions) (result DownloadByBillingSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/download", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DownloadByBillingSubscriptionThenPoll performs DownloadByBillingSubscription then polls until it's completed
func (c InvoiceClient) DownloadByBillingSubscriptionThenPoll(ctx context.Context, id BillingSubscriptionInvoiceId, options DownloadByBillingSubscriptionOperationOptions) error {
	result, err := c.DownloadByBillingSubscription(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DownloadByBillingSubscription: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadByBillingSubscription: %+v", err)
	}

	return nil
}
