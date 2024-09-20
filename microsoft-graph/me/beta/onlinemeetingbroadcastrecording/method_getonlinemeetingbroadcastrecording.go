package onlinemeetingbroadcastrecording

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

type GetOnlineMeetingBroadcastRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOnlineMeetingBroadcastRecordingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetOnlineMeetingBroadcastRecordingOperationOptions() GetOnlineMeetingBroadcastRecordingOperationOptions {
	return GetOnlineMeetingBroadcastRecordingOperationOptions{}
}

func (o GetOnlineMeetingBroadcastRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingBroadcastRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnlineMeetingBroadcastRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingBroadcastRecording - Get broadcastRecording for the navigation property onlineMeetings from me
func (c OnlineMeetingBroadcastRecordingClient) GetOnlineMeetingBroadcastRecording(ctx context.Context, id beta.MeOnlineMeetingId, options GetOnlineMeetingBroadcastRecordingOperationOptions) (result GetOnlineMeetingBroadcastRecordingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/broadcastRecording", id.ID()),
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
