package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

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

type CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentation
}

type CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions() CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions {
	return CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions{}
}

func (o CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation - Create new navigation property to
// presentations for deviceManagement
func (c GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient) CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionId, input beta.GroupPolicyPresentation, options CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationOptions) (result CreateGroupPolicyDefinitionNextVersionDefinitionPreviousPresentationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/nextVersionDefinition/previousVersionDefinition/presentations", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalGroupPolicyPresentationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
