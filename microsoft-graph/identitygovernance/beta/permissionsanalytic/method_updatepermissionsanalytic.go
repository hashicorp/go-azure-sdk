package permissionsanalytic

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePermissionsAnalyticOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePermissionsAnalyticOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePermissionsAnalyticOperationOptions() UpdatePermissionsAnalyticOperationOptions {
	return UpdatePermissionsAnalyticOperationOptions{}
}

func (o UpdatePermissionsAnalyticOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePermissionsAnalyticOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePermissionsAnalyticOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePermissionsAnalytic - Update the navigation property permissionsAnalytics in identityGovernance
func (c PermissionsAnalyticClient) UpdatePermissionsAnalytic(ctx context.Context, input beta.PermissionsAnalyticsAggregation, options UpdatePermissionsAnalyticOperationOptions) (result UpdatePermissionsAnalyticOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics",
		RetryFunc:     options.RetryFunc,
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

	return
}
