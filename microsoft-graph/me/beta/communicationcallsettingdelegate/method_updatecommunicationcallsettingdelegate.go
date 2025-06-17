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

type UpdateCommunicationCallSettingDelegateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCommunicationCallSettingDelegateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCommunicationCallSettingDelegateOperationOptions() UpdateCommunicationCallSettingDelegateOperationOptions {
	return UpdateCommunicationCallSettingDelegateOperationOptions{}
}

func (o UpdateCommunicationCallSettingDelegateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCommunicationCallSettingDelegateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCommunicationCallSettingDelegateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCommunicationCallSettingDelegate - Update the navigation property delegates in me
func (c CommunicationCallSettingDelegateClient) UpdateCommunicationCallSettingDelegate(ctx context.Context, id beta.MeCommunicationCallSettingDelegateId, input beta.DelegationSettings, options UpdateCommunicationCallSettingDelegateOperationOptions) (result UpdateCommunicationCallSettingDelegateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
