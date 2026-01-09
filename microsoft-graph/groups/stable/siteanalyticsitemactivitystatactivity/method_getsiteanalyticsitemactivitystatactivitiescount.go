package siteanalyticsitemactivitystatactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteAnalyticsItemActivityStatActivitiesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions() GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions {
	return GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions{}
}

func (o GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteAnalyticsItemActivityStatActivitiesCount - Get the number of the resource
func (c SiteAnalyticsItemActivityStatActivityClient) GetSiteAnalyticsItemActivityStatActivitiesCount(ctx context.Context, id stable.GroupIdSiteIdAnalyticsItemActivityStatId, options GetSiteAnalyticsItemActivityStatActivitiesCountOperationOptions) (result GetSiteAnalyticsItemActivityStatActivitiesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activities/$count", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
