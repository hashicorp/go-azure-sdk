package subscriptionundermanagementgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupSubscriptionsGetSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SubscriptionUnderManagementGroup
}

type ManagementGroupSubscriptionsGetSubscriptionOperationOptions struct {
	CacheControl *string
}

func DefaultManagementGroupSubscriptionsGetSubscriptionOperationOptions() ManagementGroupSubscriptionsGetSubscriptionOperationOptions {
	return ManagementGroupSubscriptionsGetSubscriptionOperationOptions{}
}

func (o ManagementGroupSubscriptionsGetSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o ManagementGroupSubscriptionsGetSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ManagementGroupSubscriptionsGetSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ManagementGroupSubscriptionsGetSubscription ...
func (c SubscriptionUnderManagementGroupsClient) ManagementGroupSubscriptionsGetSubscription(ctx context.Context, id SubscriptionId, options ManagementGroupSubscriptionsGetSubscriptionOperationOptions) (result ManagementGroupSubscriptionsGetSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model SubscriptionUnderManagementGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
