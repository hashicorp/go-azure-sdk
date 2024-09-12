package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentation
}

type GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions() GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions {
	return GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions{}
}

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentation - Get presentations from deviceManagement. The
// group policy presentations associated with the definition.
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient) GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionNextVersionDefinitionPresentationId, options GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationOptions) (result GetGroupPolicyDefinitionPreviousVersionDefinitionNextPresentationOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalGroupPolicyPresentationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
