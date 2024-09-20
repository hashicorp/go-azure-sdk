package enrichment

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

type IPGeodataGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *EnrichmentIPGeodata
}

type IPGeodataGetOperationOptions struct {
	IPAddress *string
}

func DefaultIPGeodataGetOperationOptions() IPGeodataGetOperationOptions {
	return IPGeodataGetOperationOptions{}
}

func (o IPGeodataGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o IPGeodataGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o IPGeodataGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IPAddress != nil {
		out.Append("ipAddress", fmt.Sprintf("%v", *o.IPAddress))
	}
	return &out
}

// IPGeodataGet ...
func (c EnrichmentClient) IPGeodataGet(ctx context.Context, id commonids.ResourceGroupId, options IPGeodataGetOperationOptions) (result IPGeodataGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/enrichment/ip/geodata", id.ID()),
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

	var model EnrichmentIPGeodata
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
