package insightshared

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightSharedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SharedInsight
}

type GetInsightSharedOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetInsightSharedOperationOptions() GetInsightSharedOperationOptions {
	return GetInsightSharedOperationOptions{}
}

func (o GetInsightSharedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetInsightSharedOperationOptions) ToOData() *odata.Query {
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

func (o GetInsightSharedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetInsightShared - Get shared from me. Calculated relationship that identifies documents shared with or by the user.
// This includes URLs, file attachments, and reference attachments to OneDrive for work or school and SharePoint files
// found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations.
// Ordered by recency of share.
func (c InsightSharedClient) GetInsightShared(ctx context.Context, id stable.MeInsightSharedId, options GetInsightSharedOperationOptions) (result GetInsightSharedOperationResponse, err error) {
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

	var model stable.SharedInsight
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
