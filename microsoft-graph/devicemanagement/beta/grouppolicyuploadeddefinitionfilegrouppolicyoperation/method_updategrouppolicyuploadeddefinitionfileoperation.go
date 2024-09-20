package grouppolicyuploadeddefinitionfilegrouppolicyoperation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyUploadedDefinitionFileOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions() UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions {
	return UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions{}
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyUploadedDefinitionFileOperation - Update the navigation property groupPolicyOperations in
// deviceManagement
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) UpdateGroupPolicyUploadedDefinitionFileOperation(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId, input beta.GroupPolicyOperation, options UpdateGroupPolicyUploadedDefinitionFileOperationOperationOptions) (result UpdateGroupPolicyUploadedDefinitionFileOperationOperationResponse, err error) {
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
