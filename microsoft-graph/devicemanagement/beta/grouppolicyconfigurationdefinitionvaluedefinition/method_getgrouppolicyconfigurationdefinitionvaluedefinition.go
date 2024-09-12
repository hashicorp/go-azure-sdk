package grouppolicyconfigurationdefinitionvaluedefinition

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

type GetGroupPolicyConfigurationDefinitionValueDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyDefinition
}

type GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions() GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions {
	return GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions{}
}

func (o GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyConfigurationDefinitionValueDefinition - Get definition from deviceManagement. The associated group
// policy definition with the value.
func (c GroupPolicyConfigurationDefinitionValueDefinitionClient) GetGroupPolicyConfigurationDefinitionValueDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationIdDefinitionValueId, options GetGroupPolicyConfigurationDefinitionValueDefinitionOperationOptions) (result GetGroupPolicyConfigurationDefinitionValueDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/definition", id.ID()),
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
