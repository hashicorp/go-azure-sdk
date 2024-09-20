package grouppolicydefinitionpreviousversiondefinition

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

type DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions() DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions {
	return DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions{}
}

func (o DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteGroupPolicyDefinitionPreviousVersionDefinition - Delete navigation property previousVersionDefinition for
// deviceManagement
func (c GroupPolicyDefinitionPreviousVersionDefinitionClient) DeleteGroupPolicyDefinitionPreviousVersionDefinition(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationOptions) (result DeleteGroupPolicyDefinitionPreviousVersionDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/previousVersionDefinition", id.ID()),
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
