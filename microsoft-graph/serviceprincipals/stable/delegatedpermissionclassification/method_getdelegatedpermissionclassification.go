package delegatedpermissionclassification

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDelegatedPermissionClassificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DelegatedPermissionClassification
}

type GetDelegatedPermissionClassificationOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDelegatedPermissionClassificationOperationOptions() GetDelegatedPermissionClassificationOperationOptions {
	return GetDelegatedPermissionClassificationOperationOptions{}
}

func (o GetDelegatedPermissionClassificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDelegatedPermissionClassificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDelegatedPermissionClassificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDelegatedPermissionClassification - Get delegatedPermissionClassifications from servicePrincipals
func (c DelegatedPermissionClassificationClient) GetDelegatedPermissionClassification(ctx context.Context, id stable.ServicePrincipalIdDelegatedPermissionClassificationId, options GetDelegatedPermissionClassificationOperationOptions) (result GetDelegatedPermissionClassificationOperationResponse, err error) {
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

	var model stable.DelegatedPermissionClassification
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
