package siteanalyticsitemactivitystatactivitydriveitemcontentstream

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions() DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions {
	return DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions{}
}

func (o DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStream - Delete contentStream for the navigation property
// driveItem in groups. The content stream, if the item represents a file.
func (c SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient) DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStream(ctx context.Context, id beta.GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, options DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) (result DeleteSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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

	return
}
