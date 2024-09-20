package grouppolicyuploadeddefinitionfile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyUploadedDefinitionFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyUploadedDefinitionFileOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateGroupPolicyUploadedDefinitionFileOperationOptions() UpdateGroupPolicyUploadedDefinitionFileOperationOptions {
	return UpdateGroupPolicyUploadedDefinitionFileOperationOptions{}
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyUploadedDefinitionFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyUploadedDefinitionFile - Update the navigation property groupPolicyUploadedDefinitionFiles in
// deviceManagement
func (c GroupPolicyUploadedDefinitionFileClient) UpdateGroupPolicyUploadedDefinitionFile(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, input beta.GroupPolicyUploadedDefinitionFile, options UpdateGroupPolicyUploadedDefinitionFileOperationOptions) (result UpdateGroupPolicyUploadedDefinitionFileOperationResponse, err error) {
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
