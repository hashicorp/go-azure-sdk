package sshpublickeys

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type UpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SshPublicKeyResource
	OData        *odata.OData
}

// Update ...
func (c SshPublicKeysClient) Update(ctx context.Context, id SshPublicKeyId, input SshPublicKeyUpdateResource) (result UpdateOperationResponse, err error) {
	req, err := c.Client2.NewPatchRequest(ctx, id.ID(), defaultApiVersion, odata.Query{}, input)
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
