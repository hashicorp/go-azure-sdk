package zones

import (
	"context"
	"github.com/hashicorp/go-azure-sdk/client/base"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type GetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Zone
	OData        *odata.OData
}

// Get ...
func (c ZonesClient) Get(ctx context.Context, id DnsZoneId) (result GetOperationResponse, err error) {
	req, err := c.Client2.NewGetRequest(ctx, id.ID(), defaultApiVersion, odata.Query{})
	if err != nil {
		return
	}

	var resp *base.Response
	resp, err = req.Execute()
	result.HttpResponse = resp.Response
	result.OData = resp.OData
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
