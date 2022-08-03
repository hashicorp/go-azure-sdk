package sshpublickeys

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type CreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SshPublicKeyResource
	OData        *odata.OData
}

// Create ...
func (c SshPublicKeysClient) Create(ctx context.Context, id SshPublicKeyId, input SshPublicKeyResource) (result CreateOperationResponse, err error) {
	req, err := c.Client2.NewPutRequest(ctx, id.ID(), defaultApiVersion, odata.Query{}, input)
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
