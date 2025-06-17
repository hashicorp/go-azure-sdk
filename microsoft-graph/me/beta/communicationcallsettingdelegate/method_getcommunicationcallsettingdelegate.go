package communicationcallsettingdelegate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCommunicationCallSettingDelegateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DelegationSettings
}

type GetCommunicationCallSettingDelegateOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetCommunicationCallSettingDelegateOperationOptions() GetCommunicationCallSettingDelegateOperationOptions {
	return GetCommunicationCallSettingDelegateOperationOptions{}
}

func (o GetCommunicationCallSettingDelegateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCommunicationCallSettingDelegateOperationOptions) ToOData() *odata.Query {
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

func (o GetCommunicationCallSettingDelegateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCommunicationCallSettingDelegate - Get delegationSettings. Read the properties and relationships of a
// delegationSettings object.
func (c CommunicationCallSettingDelegateClient) GetCommunicationCallSettingDelegate(ctx context.Context, id beta.MeCommunicationCallSettingDelegateId, options GetCommunicationCallSettingDelegateOperationOptions) (result GetCommunicationCallSettingDelegateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.DelegationSettings
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
