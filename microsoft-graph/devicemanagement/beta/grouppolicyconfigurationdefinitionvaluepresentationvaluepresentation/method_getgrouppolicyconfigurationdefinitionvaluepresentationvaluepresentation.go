package grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation

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

type GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GroupPolicyPresentation
}

type GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions() GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions {
	return GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions{}
}

func (o GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions) ToOData() *odata.Query {
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

func (o GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentation - Get presentation from deviceManagement. The
// group policy presentation associated with the presentation value.
func (c GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient) GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationIdDefinitionValueIdPresentationValueId, options GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationOptions) (result GetGroupPolicyConfigurationDefinitionValuePresentationValuePresentationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/presentation", id.ID()),
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
	model, err := beta.UnmarshalGroupPolicyPresentationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
