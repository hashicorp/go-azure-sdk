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

type CreateOnlineMeetingRecordingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CallRecording
}

type CreateOnlineMeetingRecordingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnlineMeetingRecordingOperationOptions() CreateOnlineMeetingRecordingOperationOptions {
	return CreateOnlineMeetingRecordingOperationOptions{}
}

func (o CreateOnlineMeetingRecordingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingRecordingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingRecordingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingRecording - Create new navigation property to recordings for me
func (c OnlineMeetingRecordingClient) CreateOnlineMeetingRecording(ctx context.Context, id beta.MeOnlineMeetingId, input beta.CallRecording, options CreateOnlineMeetingRecordingOperationOptions) (result CreateOnlineMeetingRecordingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recordings", id.ID()),
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

	var model beta.CallRecording
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
