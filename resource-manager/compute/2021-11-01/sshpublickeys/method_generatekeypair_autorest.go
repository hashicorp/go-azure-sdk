package sshpublickeys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type GenerateKeyPairOperationResponse struct {
	HttpResponse *http.Response
	Model        *SshPublicKeyGenerateKeyPairResult
	OData        *odata.OData
}

// GenerateKeyPair ...
func (c SshPublicKeysClient) GenerateKeyPair(ctx context.Context, id SshPublicKeyId) (result GenerateKeyPairOperationResponse, err error) {
	req, err := c.Client2.NewPostRequest(ctx, fmt.Sprintf("%s/generateKeyPair", id.ID()), defaultApiVersion, odata.Query{}, nil)
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
