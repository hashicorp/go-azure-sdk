package managementgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UbscriptionsGetSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SubscriptionUnderManagementGroup
}

type UbscriptionsGetSubscriptionOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsGetSubscriptionOperationOptions() UbscriptionsGetSubscriptionOperationOptions {
	return UbscriptionsGetSubscriptionOperationOptions{}
}

func (o UbscriptionsGetSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o UbscriptionsGetSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o UbscriptionsGetSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UbscriptionsGetSubscription ...
func (c ManagementGroupsClient) UbscriptionsGetSubscription(ctx context.Context, id SubscriptionId, options UbscriptionsGetSubscriptionOperationOptions) (result UbscriptionsGetSubscriptionOperationResponse, err error) {
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
