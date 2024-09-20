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

type GetGroupPolicyUploadedDefinitionFileOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyOperation
}

type GetGroupPolicyUploadedDefinitionFileOperationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetGroupPolicyUploadedDefinitionFileOperationOperationOptions() GetGroupPolicyUploadedDefinitionFileOperationOperationOptions {
	return GetGroupPolicyUploadedDefinitionFileOperationOperationOptions{}
}

func (o GetGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToOData() *odata.Query {
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

func (o GetGroupPolicyUploadedDefinitionFileOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyUploadedDefinitionFileOperation - Get groupPolicyOperations from deviceManagement. The list of
// operations on the uploaded ADMX file.
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) GetGroupPolicyUploadedDefinitionFileOperation(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileIdGroupPolicyOperationId, options GetGroupPolicyUploadedDefinitionFileOperationOperationOptions) (result GetGroupPolicyUploadedDefinitionFileOperationOperationResponse, err error) {
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

	var model beta.GroupPolicyOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
