package grouppolicydefinitionfile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyDefinitionFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyDefinitionFileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateGroupPolicyDefinitionFileOperationOptions() UpdateGroupPolicyDefinitionFileOperationOptions {
	return UpdateGroupPolicyDefinitionFileOperationOptions{}
}

func (o UpdateGroupPolicyDefinitionFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyDefinitionFileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyDefinitionFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyDefinitionFile - Update the navigation property groupPolicyDefinitionFiles in deviceManagement
func (c GroupPolicyDefinitionFileClient) UpdateGroupPolicyDefinitionFile(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionFileId, input beta.GroupPolicyDefinitionFile, options UpdateGroupPolicyDefinitionFileOperationOptions) (result UpdateGroupPolicyDefinitionFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
