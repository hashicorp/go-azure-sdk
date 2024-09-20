package onlinemeetingtranscript

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOnlineMeetingTranscriptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CallTranscript
}

type CreateOnlineMeetingTranscriptOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnlineMeetingTranscriptOperationOptions() CreateOnlineMeetingTranscriptOperationOptions {
	return CreateOnlineMeetingTranscriptOperationOptions{}
}

func (o CreateOnlineMeetingTranscriptOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingTranscriptOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingTranscriptOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingTranscript - Create new navigation property to transcripts for me
func (c OnlineMeetingTranscriptClient) CreateOnlineMeetingTranscript(ctx context.Context, id stable.MeOnlineMeetingId, input stable.CallTranscript, options CreateOnlineMeetingTranscriptOperationOptions) (result CreateOnlineMeetingTranscriptOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/transcripts", id.ID()),
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

	var model stable.CallTranscript
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
