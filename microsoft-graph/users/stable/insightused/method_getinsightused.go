package insightused

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightUsedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UsedInsight
}

type GetInsightUsedOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetInsightUsedOperationOptions() GetInsightUsedOperationOptions {
	return GetInsightUsedOperationOptions{}
}

func (o GetInsightUsedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetInsightUsedOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetInsightUsedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetInsightUsed - Get used from users. Calculated relationship that identifies the latest documents viewed or modified
// by a user, including OneDrive for work or school and SharePoint documents, ranked by recency of use.
func (c InsightUsedClient) GetInsightUsed(ctx context.Context, id stable.UserIdInsightUsedId, options GetInsightUsedOperationOptions) (result GetInsightUsedOperationResponse, err error) {
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

	var model stable.UsedInsight
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
