package usercredentialusagedetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserCredentialUsageDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserCredentialUsageDetails
}

type GetUserCredentialUsageDetailOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserCredentialUsageDetailOperationOptions() GetUserCredentialUsageDetailOperationOptions {
	return GetUserCredentialUsageDetailOperationOptions{}
}

func (o GetUserCredentialUsageDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserCredentialUsageDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserCredentialUsageDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserCredentialUsageDetail - Get userCredentialUsageDetails from reports. Represents the self-service password
// reset (SSPR) usage for a given tenant.
func (c UserCredentialUsageDetailClient) GetUserCredentialUsageDetail(ctx context.Context, id beta.ReportUserCredentialUsageDetailId, options GetUserCredentialUsageDetailOperationOptions) (result GetUserCredentialUsageDetailOperationResponse, err error) {
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

	var model beta.UserCredentialUsageDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
