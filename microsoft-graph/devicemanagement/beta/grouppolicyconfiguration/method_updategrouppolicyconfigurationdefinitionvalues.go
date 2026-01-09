package grouppolicyconfiguration

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

type UpdateGroupPolicyConfigurationDefinitionValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateGroupPolicyConfigurationDefinitionValuesOperationOptions() UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions {
	return UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions{}
}

func (o UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateGroupPolicyConfigurationDefinitionValues - Invoke action updateDefinitionValues
func (c GroupPolicyConfigurationClient) UpdateGroupPolicyConfigurationDefinitionValues(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, input UpdateGroupPolicyConfigurationDefinitionValuesRequest, options UpdateGroupPolicyConfigurationDefinitionValuesOperationOptions) (result UpdateGroupPolicyConfigurationDefinitionValuesOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/updateDefinitionValues", id.ID()),
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
