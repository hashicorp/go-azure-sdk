package setgroupsitetermstoresettermchildren

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSetGroupByIdSiteByIdTermStoreByIdSetByIdTermByIdChildrenByIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *models.TermStoreSet
}

// GetSetGroupByIdSiteByIdTermStoreByIdSetByIdTermByIdChildrenById ...
func (c SetGroupSiteTermStoreSetTermChildrenClient) GetSetGroupByIdSiteByIdTermStoreByIdSetByIdTermByIdChildrenById(ctx context.Context, id GroupSiteTermStoreSetTermChildrenId) (result GetSetGroupByIdSiteByIdTermStoreByIdSetByIdTermByIdChildrenByIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/set", id.ID()),
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
