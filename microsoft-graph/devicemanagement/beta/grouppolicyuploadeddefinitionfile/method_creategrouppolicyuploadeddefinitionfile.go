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

type CreateGroupPolicyUploadedDefinitionFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyUploadedDefinitionFile
}

type CreateGroupPolicyUploadedDefinitionFileOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateGroupPolicyUploadedDefinitionFileOperationOptions() CreateGroupPolicyUploadedDefinitionFileOperationOptions {
	return CreateGroupPolicyUploadedDefinitionFileOperationOptions{}
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyUploadedDefinitionFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyUploadedDefinitionFile - Create new navigation property to groupPolicyUploadedDefinitionFiles for
// deviceManagement
func (c GroupPolicyUploadedDefinitionFileClient) CreateGroupPolicyUploadedDefinitionFile(ctx context.Context, input beta.GroupPolicyUploadedDefinitionFile, options CreateGroupPolicyUploadedDefinitionFileOperationOptions) (result CreateGroupPolicyUploadedDefinitionFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/groupPolicyUploadedDefinitionFiles",
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

	var model beta.GroupPolicyUploadedDefinitionFile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
