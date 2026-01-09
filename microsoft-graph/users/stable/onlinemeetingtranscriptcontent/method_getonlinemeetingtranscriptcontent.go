package onlinemeetingtranscriptcontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOnlineMeetingTranscriptContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOnlineMeetingTranscriptContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetOnlineMeetingTranscriptContentOperationOptions() GetOnlineMeetingTranscriptContentOperationOptions {
	return GetOnlineMeetingTranscriptContentOperationOptions{}
}

func (o GetOnlineMeetingTranscriptContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingTranscriptContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnlineMeetingTranscriptContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingTranscriptContent - Get callTranscript. Retrieve a callTranscript object associated with a scheduled
// onlineMeeting. This API supports the retrieval of call transcripts from private chat meetings and channel meetings.
// However, private channel meetings are not supported at this time. Retrieving the transcript returns the metadata of
// the single transcript associated with the online meeting. Retrieving the content of the transcript returns the stream
// of text associated with the transcript.
func (c OnlineMeetingTranscriptContentClient) GetOnlineMeetingTranscriptContent(ctx context.Context, id stable.UserIdOnlineMeetingIdTranscriptId, options GetOnlineMeetingTranscriptContentOperationOptions) (result GetOnlineMeetingTranscriptContentOperationResponse, err error) {
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
