package grouppolicyconfigurationdefinitionvalue

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

type CreateGroupPolicyConfigurationDefinitionValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicyDefinitionValue
}

type CreateGroupPolicyConfigurationDefinitionValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyConfigurationDefinitionValueOperationOptions() CreateGroupPolicyConfigurationDefinitionValueOperationOptions {
	return CreateGroupPolicyConfigurationDefinitionValueOperationOptions{}
}

func (o CreateGroupPolicyConfigurationDefinitionValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyConfigurationDefinitionValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyConfigurationDefinitionValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyConfigurationDefinitionValue - Create new navigation property to definitionValues for
// deviceManagement
func (c GroupPolicyConfigurationDefinitionValueClient) CreateGroupPolicyConfigurationDefinitionValue(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, input beta.GroupPolicyDefinitionValue, options CreateGroupPolicyConfigurationDefinitionValueOperationOptions) (result CreateGroupPolicyConfigurationDefinitionValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/definitionValues", id.ID()),
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

	var model beta.GroupPolicyDefinitionValue
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
