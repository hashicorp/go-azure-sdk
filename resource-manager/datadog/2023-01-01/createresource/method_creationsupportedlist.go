package createresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreationSupportedListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateResourceSupportedResponseList
}

type CreationSupportedListOperationOptions struct {
	DatadogOrganizationId *string
}

func DefaultCreationSupportedListOperationOptions() CreationSupportedListOperationOptions {
	return CreationSupportedListOperationOptions{}
}

func (o CreationSupportedListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreationSupportedListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o CreationSupportedListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatadogOrganizationId != nil {
		out.Append("datadogOrganizationId", fmt.Sprintf("%v", *o.DatadogOrganizationId))
	}
	return &out
}

// CreationSupportedList ...
func (c CreateResourceClient) CreationSupportedList(ctx context.Context, id commonids.SubscriptionId, options CreationSupportedListOperationOptions) (result CreationSupportedListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          fmt.Sprintf("%s/providers/Microsoft.Datadog/subscriptionStatuses", id.ID()),
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

	var model CreateResourceSupportedResponseList
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
