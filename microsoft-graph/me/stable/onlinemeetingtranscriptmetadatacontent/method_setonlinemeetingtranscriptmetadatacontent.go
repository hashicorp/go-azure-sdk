package onlinemeetingtranscriptmetadatacontent

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

type SetOnlineMeetingTranscriptMetadataContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnlineMeetingTranscriptMetadataContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetOnlineMeetingTranscriptMetadataContentOperationOptions() SetOnlineMeetingTranscriptMetadataContentOperationOptions {
	return SetOnlineMeetingTranscriptMetadataContentOperationOptions{}
}

func (o SetOnlineMeetingTranscriptMetadataContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnlineMeetingTranscriptMetadataContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnlineMeetingTranscriptMetadataContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnlineMeetingTranscriptMetadataContent - Update metadataContent for the navigation property transcripts in me. The
// time-aligned metadata of the utterances in the transcript. Read-only.
func (c OnlineMeetingTranscriptMetadataContentClient) SetOnlineMeetingTranscriptMetadataContent(ctx context.Context, id stable.MeOnlineMeetingIdTranscriptId, input []byte, options SetOnlineMeetingTranscriptMetadataContentOperationOptions) (result SetOnlineMeetingTranscriptMetadataContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/metadataContent", id.ID()),
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
