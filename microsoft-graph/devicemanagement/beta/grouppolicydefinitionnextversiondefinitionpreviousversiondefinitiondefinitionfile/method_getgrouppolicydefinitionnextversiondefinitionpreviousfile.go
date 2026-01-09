package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyDefinitionFile
}

type GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions() GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions {
	return GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions{}
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions) ToOData() *odata.Query {
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

func (o GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyDefinitionNextVersionDefinitionPreviousFile - Get definitionFile from deviceManagement. The group
// policy file associated with the definition.
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient) GetGroupPolicyDefinitionNextVersionDefinitionPreviousFile(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, options GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationOptions) (result GetGroupPolicyDefinitionNextVersionDefinitionPreviousFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/nextVersionDefinition/previousVersionDefinition/definitionFile", id.ID()),
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
