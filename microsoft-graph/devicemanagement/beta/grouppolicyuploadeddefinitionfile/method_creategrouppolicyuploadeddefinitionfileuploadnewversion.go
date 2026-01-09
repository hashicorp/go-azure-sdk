package grouppolicyuploadeddefinitionfile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions() CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions {
	return CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions{}
}

func (o CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyUploadedDefinitionFileUploadNewVersion - Invoke action uploadNewVersion
func (c GroupPolicyUploadedDefinitionFileClient) CreateGroupPolicyUploadedDefinitionFileUploadNewVersion(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, input CreateGroupPolicyUploadedDefinitionFileUploadNewVersionRequest, options CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions) (result CreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/uploadNewVersion", id.ID()),
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
