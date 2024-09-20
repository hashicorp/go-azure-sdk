package siteanalyticsitemactivitystat

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSiteAnalyticsItemActivityStatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteAnalyticsItemActivityStatOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteAnalyticsItemActivityStatOperationOptions() UpdateSiteAnalyticsItemActivityStatOperationOptions {
	return UpdateSiteAnalyticsItemActivityStatOperationOptions{}
}

func (o UpdateSiteAnalyticsItemActivityStatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteAnalyticsItemActivityStatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteAnalyticsItemActivityStatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteAnalyticsItemActivityStat - Update the navigation property itemActivityStats in groups
func (c SiteAnalyticsItemActivityStatClient) UpdateSiteAnalyticsItemActivityStat(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatId, input stable.ItemActivityStat, options UpdateSiteAnalyticsItemActivityStatOperationOptions) (result UpdateSiteAnalyticsItemActivityStatOperationResponse, err error) {
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
