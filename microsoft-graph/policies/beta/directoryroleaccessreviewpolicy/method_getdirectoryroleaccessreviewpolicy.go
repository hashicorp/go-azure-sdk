package directoryroleaccessreviewpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDirectoryRoleAccessReviewPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DirectoryRoleAccessReviewPolicy
}

type GetDirectoryRoleAccessReviewPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDirectoryRoleAccessReviewPolicyOperationOptions() GetDirectoryRoleAccessReviewPolicyOperationOptions {
	return GetDirectoryRoleAccessReviewPolicyOperationOptions{}
}

func (o GetDirectoryRoleAccessReviewPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleAccessReviewPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDirectoryRoleAccessReviewPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleAccessReviewPolicy - Get directoryRoleAccessReviewPolicy from policies
func (c DirectoryRoleAccessReviewPolicyClient) GetDirectoryRoleAccessReviewPolicy(ctx context.Context, options GetDirectoryRoleAccessReviewPolicyOperationOptions) (result GetDirectoryRoleAccessReviewPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/directoryRoleAccessReviewPolicy",
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

	var model beta.DirectoryRoleAccessReviewPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
