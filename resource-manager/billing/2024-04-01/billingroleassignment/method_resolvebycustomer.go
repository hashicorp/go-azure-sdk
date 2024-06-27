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

type ResolveByCustomerOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ResolveByCustomerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ResolveByCustomerOperationOptions struct {
	Filter                   *string
	ResolveScopeDisplayNames *bool
}

func DefaultResolveByCustomerOperationOptions() ResolveByCustomerOperationOptions {
	return ResolveByCustomerOperationOptions{}
}

func (o ResolveByCustomerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResolveByCustomerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ResolveByCustomerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.ResolveScopeDisplayNames != nil {
		out.Append("resolveScopeDisplayNames", fmt.Sprintf("%v", *o.ResolveScopeDisplayNames))
	}
	return &out
}

// ResolveByCustomer ...
func (c BillingRoleAssignmentClient) ResolveByCustomer(ctx context.Context, id BillingProfileCustomerId, options ResolveByCustomerOperationOptions) (result ResolveByCustomerOperationResponse, err error) {
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

// ResolveByCustomerThenPoll performs ResolveByCustomer then polls until it's completed
func (c BillingRoleAssignmentClient) ResolveByCustomerThenPoll(ctx context.Context, id BillingProfileCustomerId, options ResolveByCustomerOperationOptions) error {
	result, err := c.ResolveByCustomer(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ResolveByCustomer: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResolveByCustomer: %+v", err)
	}

	return nil
}
