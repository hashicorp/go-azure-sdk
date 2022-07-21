package zones

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client"
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

	var resp *client.Response
	resp, result.OData, _, err = req.Execute()
	result.HttpResponse = resp.Response
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil { // TODO: pointer to a pointer needed?
		return
	}

	return
}
