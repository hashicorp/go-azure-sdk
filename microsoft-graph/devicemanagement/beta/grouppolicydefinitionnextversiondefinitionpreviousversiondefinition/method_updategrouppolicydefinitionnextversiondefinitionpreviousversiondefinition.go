package grouppolicydefinitionnextversiondefinitionpreviousversiondefinition

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

type UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions() UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions {
	return UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions{}
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition - Update the navigation property
// previousVersionDefinition in deviceManagement
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient) UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, input beta.GroupPolicyDefinition, options UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) (result UpdateGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/nextVersionDefinition/previousVersionDefinition", id.ID()),
		RetryFunc:     options.RetryFunc,
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
