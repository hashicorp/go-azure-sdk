package b2xuserflowlanguagedefaultpage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowLanguageDefaultPageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

// GetB2xUserFlowLanguageDefaultPageValue - Get userFlowLanguagePage. Read the values in a userFlowLanguagePage object
// for a language in a user flow. These values are shown to a user during a user journey defined by a user flow.
func (c B2xUserFlowLanguageDefaultPageClient) GetB2xUserFlowLanguageDefaultPageValue(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdDefaultPageId) (result GetB2xUserFlowLanguageDefaultPageValueOperationResponse, err error) {
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
