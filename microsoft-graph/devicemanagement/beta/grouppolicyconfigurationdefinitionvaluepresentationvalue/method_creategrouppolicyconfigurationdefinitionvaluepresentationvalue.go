package grouppolicyconfigurationdefinitionvaluepresentationvalue

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

type CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentationValue
}

type CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions() CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions {
	return CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions{}
}

func (o CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyConfigurationDefinitionValuePresentationValue - Create new navigation property to presentationValues
// for deviceManagement
func (c GroupPolicyConfigurationDefinitionValuePresentationValueClient) CreateGroupPolicyConfigurationDefinitionValuePresentationValue(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationIdDefinitionValueId, input beta.GroupPolicyPresentationValue, options CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationOptions) (result CreateGroupPolicyConfigurationDefinitionValuePresentationValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/presentationValues", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalGroupPolicyPresentationValueImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
