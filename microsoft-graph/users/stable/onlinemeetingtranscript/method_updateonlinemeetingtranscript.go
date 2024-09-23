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

type UpdateOnlineMeetingTranscriptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnlineMeetingTranscriptOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnlineMeetingTranscriptOperationOptions() UpdateOnlineMeetingTranscriptOperationOptions {
	return UpdateOnlineMeetingTranscriptOperationOptions{}
}

func (o UpdateOnlineMeetingTranscriptOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnlineMeetingTranscriptOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnlineMeetingTranscriptOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnlineMeetingTranscript - Update the navigation property transcripts in users
func (c OnlineMeetingTranscriptClient) UpdateOnlineMeetingTranscript(ctx context.Context, id stable.UserIdOnlineMeetingIdTranscriptId, input stable.CallTranscript, options UpdateOnlineMeetingTranscriptOperationOptions) (result UpdateOnlineMeetingTranscriptOperationResponse, err error) {
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
