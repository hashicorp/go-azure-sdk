package onlinemeetingrecordingcontent

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

type GetOnlineMeetingRecordingContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOnlineMeetingRecordingContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetOnlineMeetingRecordingContentOperationOptions() GetOnlineMeetingRecordingContentOperationOptions {
	return GetOnlineMeetingRecordingContentOperationOptions{}
}

func (o GetOnlineMeetingRecordingContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingRecordingContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnlineMeetingRecordingContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingRecordingContent - Get callRecording. Get a callRecording object associated with a scheduled
// onlineMeeting. This API supports the retrieval of call recordings from private chat meetings and channel meetings.
// However, private channel meetings are not supported at this time. For a recording, this API returns the metadata of
// the single recording associated with the online meeting. For the content of a recording, this API returns the stream
// of bytes associated with the recording.
func (c OnlineMeetingRecordingContentClient) GetOnlineMeetingRecordingContent(ctx context.Context, id beta.UserIdOnlineMeetingIdRecordingId, options GetOnlineMeetingRecordingContentOperationOptions) (result GetOnlineMeetingRecordingContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
