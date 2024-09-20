package onlinemeetingtranscript

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOnlineMeetingTranscriptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CallTranscript
}

type GetOnlineMeetingTranscriptOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetOnlineMeetingTranscriptOperationOptions() GetOnlineMeetingTranscriptOperationOptions {
	return GetOnlineMeetingTranscriptOperationOptions{}
}

func (o GetOnlineMeetingTranscriptOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingTranscriptOperationOptions) ToOData() *odata.Query {
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

func (o GetOnlineMeetingTranscriptOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingTranscript - Get callTranscript. Retrieve a callTranscript object associated with a scheduled
// onlineMeeting. This API doesn't support getting call transcripts from channel meetings. Retrieving the transcript
// returns the metadata of the single transcript associated with the online meeting. Retrieving the content of the
// transcript returns the stream of text associated with the transcript.
func (c OnlineMeetingTranscriptClient) GetOnlineMeetingTranscript(ctx context.Context, id stable.UserIdOnlineMeetingIdTranscriptId, options GetOnlineMeetingTranscriptOperationOptions) (result GetOnlineMeetingTranscriptOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.CallTranscript
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
