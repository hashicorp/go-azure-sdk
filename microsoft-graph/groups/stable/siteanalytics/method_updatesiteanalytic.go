package siteanalytics

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

type UpdateSiteAnalyticOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteAnalyticOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteAnalyticOperationOptions() UpdateSiteAnalyticOperationOptions {
	return UpdateSiteAnalyticOperationOptions{}
}

func (o UpdateSiteAnalyticOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteAnalyticOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteAnalyticOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteAnalytic - Update the navigation property analytics in groups
func (c SiteAnalyticsClient) UpdateSiteAnalytic(ctx context.Context, id stable.GroupIdSiteId, input stable.ItemAnalytics, options UpdateSiteAnalyticOperationOptions) (result UpdateSiteAnalyticOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/analytics", id.ID()),
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
