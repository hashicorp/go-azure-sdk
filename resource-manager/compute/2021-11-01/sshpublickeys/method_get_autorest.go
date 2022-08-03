package sshpublickeys

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type GetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SshPublicKeyResource
	OData        *odata.OData
}

// Get ...
func (c SshPublicKeysClient) Get(ctx context.Context, id SshPublicKeyId) (result GetOperationResponse, err error) {
	req, err := c.Client2.NewGetRequest(ctx, id.ID(), defaultApiVersion, odata.Query{})
	if err != nil {
		return
	}

	var resp *base.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
