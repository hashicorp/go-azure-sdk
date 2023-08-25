package keys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Key
}

type ListVersionsCompleteResult struct {
	Items []Key
}

// ListVersions ...
func (c KeysClient) ListVersions(ctx context.Context, id commonids.KeyVaultKeyId) (result ListVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/versions", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Key `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVersionsComplete retrieves all the results into a single object
func (c KeysClient) ListVersionsComplete(ctx context.Context, id commonids.KeyVaultKeyId) (ListVersionsCompleteResult, error) {
	return c.ListVersionsCompleteMatchingPredicate(ctx, id, KeyOperationPredicate{})
}

// ListVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c KeysClient) ListVersionsCompleteMatchingPredicate(ctx context.Context, id commonids.KeyVaultKeyId, predicate KeyOperationPredicate) (result ListVersionsCompleteResult, err error) {
	items := make([]Key, 0)

	resp, err := c.ListVersions(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListVersionsCompleteResult{
		Items: items,
	}
	return
}
