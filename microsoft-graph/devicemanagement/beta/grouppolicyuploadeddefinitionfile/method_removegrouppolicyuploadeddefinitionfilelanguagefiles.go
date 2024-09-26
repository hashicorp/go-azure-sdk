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

type RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions() RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions {
	return RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions{}
}

func (o RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveGroupPolicyUploadedDefinitionFileLanguageFiles - Invoke action removeLanguageFiles
func (c GroupPolicyUploadedDefinitionFileClient) RemoveGroupPolicyUploadedDefinitionFileLanguageFiles(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, input RemoveGroupPolicyUploadedDefinitionFileLanguageFilesRequest, options RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions) (result RemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/removeLanguageFiles", id.ID()),
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
