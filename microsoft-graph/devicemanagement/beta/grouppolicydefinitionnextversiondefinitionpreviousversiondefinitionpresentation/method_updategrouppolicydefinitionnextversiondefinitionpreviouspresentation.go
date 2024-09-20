package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions() UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions {
	return UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions{}
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation - Update the navigation property presentations
// in deviceManagement
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionIdNextVersionDefinitionPreviousVersionDefinitionPresentationId, input beta.GroupPolicyPresentation, options UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) (result UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
