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

type UbscriptionsCreateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SubscriptionUnderManagementGroup
}

type UbscriptionsCreateOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsCreateOperationOptions() UbscriptionsCreateOperationOptions {
	return UbscriptionsCreateOperationOptions{}
}

func (o UbscriptionsCreateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o UbscriptionsCreateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o UbscriptionsCreateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UbscriptionsCreate ...
func (c ManagementGroupsClient) UbscriptionsCreate(ctx context.Context, id SubscriptionId, options UbscriptionsCreateOperationOptions) (result UbscriptionsCreateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		Path:          id.ID(),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
