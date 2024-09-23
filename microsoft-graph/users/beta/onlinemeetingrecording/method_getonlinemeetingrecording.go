package onlinemeetingrecording

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOnlineMeetingRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CallRecording
}

type GetOnlineMeetingRecordingOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetOnlineMeetingRecordingOperationOptions() GetOnlineMeetingRecordingOperationOptions {
	return GetOnlineMeetingRecordingOperationOptions{}
}

func (o GetOnlineMeetingRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingRecordingOperationOptions) ToOData() *odata.Query {
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

func (o GetOnlineMeetingRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingRecording - Get callRecording. Get a callRecording object associated with a scheduled onlineMeeting.
// This API doesn't support getting call recordings from channel meetings. For a recording, this API returns the
// metadata of the single recording associated with the online meeting. For the content of a recording, this API returns
// the stream of bytes associated with the recording.
func (c OnlineMeetingRecordingClient) GetOnlineMeetingRecording(ctx context.Context, id beta.UserIdOnlineMeetingIdRecordingId, options GetOnlineMeetingRecordingOperationOptions) (result GetOnlineMeetingRecordingOperationResponse, err error) {
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

	var model beta.CallRecording
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
