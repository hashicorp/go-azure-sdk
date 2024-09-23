package onlinemeetingrecording

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnlineMeetingRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnlineMeetingRecordingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnlineMeetingRecordingOperationOptions() UpdateOnlineMeetingRecordingOperationOptions {
	return UpdateOnlineMeetingRecordingOperationOptions{}
}

func (o UpdateOnlineMeetingRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnlineMeetingRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnlineMeetingRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnlineMeetingRecording - Update the navigation property recordings in users
func (c OnlineMeetingRecordingClient) UpdateOnlineMeetingRecording(ctx context.Context, id stable.UserIdOnlineMeetingIdRecordingId, input stable.CallRecording, options UpdateOnlineMeetingRecordingOperationOptions) (result UpdateOnlineMeetingRecordingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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
