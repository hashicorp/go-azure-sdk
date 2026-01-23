package sshpublickeyresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SshPublicKeysGenerateKeyPairOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SshPublicKeyGenerateKeyPairResult
}

// SshPublicKeysGenerateKeyPair ...
func (c SshPublicKeyResourcesClient) SshPublicKeysGenerateKeyPair(ctx context.Context, id SshPublicKeyId, input SshGenerateKeyPairInputParameters) (result SshPublicKeysGenerateKeyPairOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/generateKeyPair", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model SshPublicKeyGenerateKeyPairResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
