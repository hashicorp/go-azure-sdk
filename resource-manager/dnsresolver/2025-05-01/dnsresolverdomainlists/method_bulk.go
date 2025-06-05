package dnsresolverdomainlists

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

type BulkOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DnsResolverDomainList
}

type BulkOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultBulkOperationOptions() BulkOperationOptions {
	return BulkOperationOptions{}
}

func (o BulkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o BulkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BulkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// Bulk ...
func (c DnsResolverDomainListsClient) Bulk(ctx context.Context, id DnsResolverDomainListId, input DnsResolverDomainListBulk, options BulkOperationOptions) (result BulkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bulk", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

// BulkThenPoll performs Bulk then polls until it's completed
func (c DnsResolverDomainListsClient) BulkThenPoll(ctx context.Context, id DnsResolverDomainListId, input DnsResolverDomainListBulk, options BulkOperationOptions) error {
	result, err := c.Bulk(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing Bulk: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after Bulk: %+v", err)
	}

	return nil
}
