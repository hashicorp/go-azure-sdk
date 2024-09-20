package siteanalyticsitemactivitystatactivity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteAnalyticsItemActivityStatActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ItemActivity
}

type GetSiteAnalyticsItemActivityStatActivityOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSiteAnalyticsItemActivityStatActivityOperationOptions() GetSiteAnalyticsItemActivityStatActivityOperationOptions {
	return GetSiteAnalyticsItemActivityStatActivityOperationOptions{}
}

func (o GetSiteAnalyticsItemActivityStatActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteAnalyticsItemActivityStatActivityOperationOptions) ToOData() *odata.Query {
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

func (o GetSiteAnalyticsItemActivityStatActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteAnalyticsItemActivityStatActivity - Get activities from groups. Exposes the itemActivities represented in this
// itemActivityStat resource.
func (c SiteAnalyticsItemActivityStatActivityClient) GetSiteAnalyticsItemActivityStatActivity(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, options GetSiteAnalyticsItemActivityStatActivityOperationOptions) (result GetSiteAnalyticsItemActivityStatActivityOperationResponse, err error) {
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

	var model stable.ItemActivity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
