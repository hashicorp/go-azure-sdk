package siteanalyticsitemactivitystatactivity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSiteAnalyticsItemActivityStatActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteAnalyticsItemActivityStatActivityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteAnalyticsItemActivityStatActivityOperationOptions() UpdateSiteAnalyticsItemActivityStatActivityOperationOptions {
	return UpdateSiteAnalyticsItemActivityStatActivityOperationOptions{}
}

func (o UpdateSiteAnalyticsItemActivityStatActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteAnalyticsItemActivityStatActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteAnalyticsItemActivityStatActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteAnalyticsItemActivityStatActivity - Update the navigation property activities in groups
func (c SiteAnalyticsItemActivityStatActivityClient) UpdateSiteAnalyticsItemActivityStatActivity(ctx context.Context, id beta.GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, input beta.ItemActivity, options UpdateSiteAnalyticsItemActivityStatActivityOperationOptions) (result UpdateSiteAnalyticsItemActivityStatActivityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
