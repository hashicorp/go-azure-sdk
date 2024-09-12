package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition

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

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyDefinition
}

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions() GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions {
	return GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions{}
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinition - Get definition from deviceManagement.
// The group policy definition associated with the presentation.
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient) GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId, options GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationOptions) (result GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationDefinitionOperationResponse, err error) {
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
