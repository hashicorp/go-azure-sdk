package identityb2xuserflowlanguageoverridespage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIdentityB2xUserFlowByIdLanguageByIdOverridesPageByIdValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// GetIdentityB2xUserFlowByIdLanguageByIdOverridesPageByIdValue ...
func (c IdentityB2xUserFlowLanguageOverridesPageClient) GetIdentityB2xUserFlowByIdLanguageByIdOverridesPageByIdValue(ctx context.Context, id IdentityB2xUserFlowLanguageOverridesPageId) (result GetIdentityB2xUserFlowByIdLanguageByIdOverridesPageByIdValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/$value", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
