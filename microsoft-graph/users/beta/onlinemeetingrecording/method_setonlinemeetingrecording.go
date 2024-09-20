package onlinemeetingrecording

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

type SetOnlineMeetingRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnlineMeetingRecordingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetOnlineMeetingRecordingOperationOptions() SetOnlineMeetingRecordingOperationOptions {
	return SetOnlineMeetingRecordingOperationOptions{}
}

func (o SetOnlineMeetingRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnlineMeetingRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnlineMeetingRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnlineMeetingRecording - Update recording for the navigation property onlineMeetings in users. The content stream
// of the recording of a Teams live event. Read-only.
func (c OnlineMeetingRecordingClient) SetOnlineMeetingRecording(ctx context.Context, id beta.UserIdOnlineMeetingId, input []byte, options SetOnlineMeetingRecordingOperationOptions) (result SetOnlineMeetingRecordingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recording", id.ID()),
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
