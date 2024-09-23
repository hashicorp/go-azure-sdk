package grouppolicyuploadeddefinitionfile

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

type RemoveGroupPolicyUploadedDefinitionFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveGroupPolicyUploadedDefinitionFileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveGroupPolicyUploadedDefinitionFileOperationOptions() RemoveGroupPolicyUploadedDefinitionFileOperationOptions {
	return RemoveGroupPolicyUploadedDefinitionFileOperationOptions{}
}

func (o RemoveGroupPolicyUploadedDefinitionFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveGroupPolicyUploadedDefinitionFileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveGroupPolicyUploadedDefinitionFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveGroupPolicyUploadedDefinitionFile - Invoke action remove
func (c GroupPolicyUploadedDefinitionFileClient) RemoveGroupPolicyUploadedDefinitionFile(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, options RemoveGroupPolicyUploadedDefinitionFileOperationOptions) (result RemoveGroupPolicyUploadedDefinitionFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/remove", id.ID()),
		RetryFunc:     options.RetryFunc,
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
