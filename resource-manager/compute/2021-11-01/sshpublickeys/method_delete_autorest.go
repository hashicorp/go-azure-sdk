package sshpublickeys

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type DeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// Delete ...
func (c SshPublicKeysClient) Delete(ctx context.Context, id SshPublicKeyId) (result DeleteOperationResponse, err error) {
	req, err := c.Client2.NewDeleteRequest(ctx, id.ID(), defaultApiVersion, odata.Query{})
	if err != nil {
		return
	}

	var resp *client.Response
	resp, result.OData, _, err = req.Execute()
	result.HttpResponse = resp.Response
	if err != nil {
		return
	}

	return
}
