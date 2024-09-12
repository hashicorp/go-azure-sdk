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

type DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions() DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions {
	return DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions{}
}

func (o DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition - Delete navigation property
// previousVersionDefinition for deviceManagement
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient) DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationOptions) (result DeleteGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/nextVersionDefinition/previousVersionDefinition", id.ID()),
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

	return
}
