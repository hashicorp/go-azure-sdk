package grouppolicydefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyDefinition
}

type GetGroupPolicyDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetGroupPolicyDefinitionOperationOptions() GetGroupPolicyDefinitionOperationOptions {
	return GetGroupPolicyDefinitionOperationOptions{}
}

func (o GetGroupPolicyDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinition - Get groupPolicyDefinitions from deviceManagement. The available group policy definitions
// for this account.
func (c GroupPolicyDefinitionClient) GetGroupPolicyDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options GetGroupPolicyDefinitionOperationOptions) (result GetGroupPolicyDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model beta.GroupPolicyDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
