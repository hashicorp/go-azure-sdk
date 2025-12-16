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

type ManagementGroupSubscriptionsDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ManagementGroupSubscriptionsDeleteOperationOptions struct {
	CacheControl *string
}

func DefaultManagementGroupSubscriptionsDeleteOperationOptions() ManagementGroupSubscriptionsDeleteOperationOptions {
	return ManagementGroupSubscriptionsDeleteOperationOptions{}
}

func (o ManagementGroupSubscriptionsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o ManagementGroupSubscriptionsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ManagementGroupSubscriptionsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ManagementGroupSubscriptionsDelete ...
func (c SubscriptionUnderManagementGroupsClient) ManagementGroupSubscriptionsDelete(ctx context.Context, id SubscriptionId, options ManagementGroupSubscriptionsDeleteOperationOptions) (result ManagementGroupSubscriptionsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
