package grouppolicydefinitionnextversiondefinition

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

type UpdateGroupPolicyDefinitionNextVersionDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions() UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions {
	return UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions{}
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyDefinitionNextVersionDefinition - Update the navigation property nextVersionDefinition in
// deviceManagement
func (c GroupPolicyDefinitionNextVersionDefinitionClient) UpdateGroupPolicyDefinitionNextVersionDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, input beta.GroupPolicyDefinition, options UpdateGroupPolicyDefinitionNextVersionDefinitionOperationOptions) (result UpdateGroupPolicyDefinitionNextVersionDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/nextVersionDefinition", id.ID()),
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
