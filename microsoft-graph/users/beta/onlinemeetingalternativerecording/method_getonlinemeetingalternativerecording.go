package onlinemeetingalternativerecording

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

type GetOnlineMeetingAlternativeRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOnlineMeetingAlternativeRecordingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetOnlineMeetingAlternativeRecordingOperationOptions() GetOnlineMeetingAlternativeRecordingOperationOptions {
	return GetOnlineMeetingAlternativeRecordingOperationOptions{}
}

func (o GetOnlineMeetingAlternativeRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnlineMeetingAlternativeRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetOnlineMeetingAlternativeRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnlineMeetingAlternativeRecording - Get alternativeRecording for the navigation property onlineMeetings from
// users. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (c OnlineMeetingAlternativeRecordingClient) GetOnlineMeetingAlternativeRecording(ctx context.Context, id beta.UserIdOnlineMeetingId, options GetOnlineMeetingAlternativeRecordingOperationOptions) (result GetOnlineMeetingAlternativeRecordingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/alternativeRecording", id.ID()),
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
