package onlinemeeting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrGetOnlineMeetingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnlineMeeting
}

type CreateOrGetOnlineMeetingsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOrGetOnlineMeetingsOperationOptions() CreateOrGetOnlineMeetingsOperationOptions {
	return CreateOrGetOnlineMeetingsOperationOptions{}
}

func (o CreateOrGetOnlineMeetingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOrGetOnlineMeetingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOrGetOnlineMeetingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOrGetOnlineMeetings - Invoke action createOrGet. Create an onlineMeeting object with a custom specified
// external ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID.
func (c OnlineMeetingClient) CreateOrGetOnlineMeetings(ctx context.Context, input CreateOrGetOnlineMeetingsRequest, options CreateOrGetOnlineMeetingsOperationOptions) (result CreateOrGetOnlineMeetingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/onlineMeetings/createOrGet",
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

	var model stable.OnlineMeeting
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
