package locationcapabilities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilitiesListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *LocationCapabilities
}

type CapabilitiesListByLocationOperationOptions struct {
	Include *CapabilityGroup
}

func DefaultCapabilitiesListByLocationOperationOptions() CapabilitiesListByLocationOperationOptions {
	return CapabilitiesListByLocationOperationOptions{}
}

func (o CapabilitiesListByLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CapabilitiesListByLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CapabilitiesListByLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Include != nil {
		out.Append("include", fmt.Sprintf("%v", *o.Include))
	}
	return &out
}

// CapabilitiesListByLocation ...
func (c LocationCapabilitiesClient) CapabilitiesListByLocation(ctx context.Context, id LocationId, options CapabilitiesListByLocationOperationOptions) (result CapabilitiesListByLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/capabilities", id.ID()),
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

	var model LocationCapabilities
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
