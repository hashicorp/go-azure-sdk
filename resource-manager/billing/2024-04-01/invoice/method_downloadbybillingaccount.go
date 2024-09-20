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

type DownloadByBillingAccountOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DocumentDownloadResult
}

type DownloadByBillingAccountOperationOptions struct {
	DocumentName *string
}

func DefaultDownloadByBillingAccountOperationOptions() DownloadByBillingAccountOperationOptions {
	return DownloadByBillingAccountOperationOptions{}
}

func (o DownloadByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DownloadByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DownloadByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DocumentName != nil {
		out.Append("documentName", fmt.Sprintf("%v", *o.DocumentName))
	}
	return &out
}

// DownloadByBillingAccount ...
func (c InvoiceClient) DownloadByBillingAccount(ctx context.Context, id BillingAccountInvoiceId, options DownloadByBillingAccountOperationOptions) (result DownloadByBillingAccountOperationResponse, err error) {
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

// DownloadByBillingAccountThenPoll performs DownloadByBillingAccount then polls until it's completed
func (c InvoiceClient) DownloadByBillingAccountThenPoll(ctx context.Context, id BillingAccountInvoiceId, options DownloadByBillingAccountOperationOptions) error {
	result, err := c.DownloadByBillingAccount(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DownloadByBillingAccount: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadByBillingAccount: %+v", err)
	}

	return nil
}
