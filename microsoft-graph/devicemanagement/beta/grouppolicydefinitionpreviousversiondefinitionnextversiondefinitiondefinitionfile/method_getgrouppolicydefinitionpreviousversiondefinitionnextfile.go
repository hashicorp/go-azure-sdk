package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyDefinitionFile
}

type GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions() GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions {
	return GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions{}
}

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions) ToOData() *odata.Query {
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

func (o GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionPreviousVersionDefinitionNextFile - Get definitionFile from deviceManagement. The group
// policy file associated with the definition.
func (c GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient) GetGroupPolicyDefinitionPreviousVersionDefinitionNextFile(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationOptions) (result GetGroupPolicyDefinitionPreviousVersionDefinitionNextFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/previousVersionDefinition/nextVersionDefinition/definitionFile", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalGroupPolicyDefinitionFileImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
