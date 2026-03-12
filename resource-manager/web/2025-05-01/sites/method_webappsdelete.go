package sites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WebAppsDeleteOperationOptions struct {
	DeleteEmptyServerFarm *bool
	DeleteMetrics         *bool
}

func DefaultWebAppsDeleteOperationOptions() WebAppsDeleteOperationOptions {
	return WebAppsDeleteOperationOptions{}
}

func (o WebAppsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WebAppsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WebAppsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DeleteEmptyServerFarm != nil {
		out.Append("deleteEmptyServerFarm", fmt.Sprintf("%v", *o.DeleteEmptyServerFarm))
	}
	if o.DeleteMetrics != nil {
		out.Append("deleteMetrics", fmt.Sprintf("%v", *o.DeleteMetrics))
	}
	return &out
}

// WebAppsDelete ...
func (c SitesClient) WebAppsDelete(ctx context.Context, id commonids.AppServiceId, options WebAppsDeleteOperationOptions) (result WebAppsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
