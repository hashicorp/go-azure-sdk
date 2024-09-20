package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

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

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentation
}

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions() GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions {
	return GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions{}
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToOData() *odata.Query {
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

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation - Get presentations from deviceManagement. The
// group policy presentations associated with the definition.
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId, options GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) (result GetGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse, err error) {
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
