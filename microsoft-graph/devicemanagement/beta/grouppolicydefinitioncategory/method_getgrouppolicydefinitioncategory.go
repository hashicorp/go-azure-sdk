package grouppolicydefinitioncategory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyDefinitionCategoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyCategory
}

type GetGroupPolicyDefinitionCategoryOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetGroupPolicyDefinitionCategoryOperationOptions() GetGroupPolicyDefinitionCategoryOperationOptions {
	return GetGroupPolicyDefinitionCategoryOperationOptions{}
}

func (o GetGroupPolicyDefinitionCategoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionCategoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyDefinitionCategoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionCategory - Get category from deviceManagement. The group policy category associated with the
// definition.
func (c GroupPolicyDefinitionCategoryClient) GetGroupPolicyDefinitionCategory(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options GetGroupPolicyDefinitionCategoryOperationOptions) (result GetGroupPolicyDefinitionCategoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/category", id.ID()),
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

	var model beta.GroupPolicyCategory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
