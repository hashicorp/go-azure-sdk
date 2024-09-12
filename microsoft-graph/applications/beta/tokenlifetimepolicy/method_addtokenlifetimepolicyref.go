package tokenlifetimepolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddTokenLifetimePolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddTokenLifetimePolicyRef - Assign tokenLifetimePolicy. Assign a tokenLifetimePolicy to an application or
// servicePrincipal. You can have multiple tokenLifetimePolicy policies in a tenant but can assign only one
// tokenLifetimePolicy per application.
func (c TokenLifetimePolicyClient) AddTokenLifetimePolicyRef(ctx context.Context, id beta.ApplicationId, input beta.ReferenceCreate) (result AddTokenLifetimePolicyRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/tokenLifetimePolicies/$ref", id.ID()),
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

	return
}
