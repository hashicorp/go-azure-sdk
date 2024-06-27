package billingroleassignment

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

type ResolveByBillingAccountOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ResolveByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ResolveByBillingAccountOperationOptions struct {
	Filter                   *string
	ResolveScopeDisplayNames *bool
}

func DefaultResolveByBillingAccountOperationOptions() ResolveByBillingAccountOperationOptions {
	return ResolveByBillingAccountOperationOptions{}
}

func (o ResolveByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResolveByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ResolveByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.ResolveScopeDisplayNames != nil {
		out.Append("resolveScopeDisplayNames", fmt.Sprintf("%v", *o.ResolveScopeDisplayNames))
	}
	return &out
}

// ResolveByBillingAccount ...
func (c BillingRoleAssignmentClient) ResolveByBillingAccount(ctx context.Context, id BillingAccountId, options ResolveByBillingAccountOperationOptions) (result ResolveByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/resolveBillingRoleAssignments", id.ID()),
		OptionsObject: options,
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

// ResolveByBillingAccountThenPoll performs ResolveByBillingAccount then polls until it's completed
func (c BillingRoleAssignmentClient) ResolveByBillingAccountThenPoll(ctx context.Context, id BillingAccountId, options ResolveByBillingAccountOperationOptions) error {
	result, err := c.ResolveByBillingAccount(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ResolveByBillingAccount: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResolveByBillingAccount: %+v", err)
	}

	return nil
}
