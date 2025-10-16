package keyoperationgroup

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

type KeysListVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Key
}

type KeysListVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Key
}

type KeysListVersionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *KeysListVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// KeysListVersions ...
func (c KeyOperationGroupClient) KeysListVersions(ctx context.Context, id commonids.KeyVaultKeyId) (result KeysListVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &KeysListVersionsCustomPager{},
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

// KeysListVersionsComplete retrieves all the results into a single object
func (c KeyOperationGroupClient) KeysListVersionsComplete(ctx context.Context, id commonids.KeyVaultKeyId) (KeysListVersionsCompleteResult, error) {
	return c.KeysListVersionsCompleteMatchingPredicate(ctx, id, KeyOperationPredicate{})
}

// KeysListVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c KeyOperationGroupClient) KeysListVersionsCompleteMatchingPredicate(ctx context.Context, id commonids.KeyVaultKeyId, predicate KeyOperationPredicate) (result KeysListVersionsCompleteResult, err error) {
	items := make([]Key, 0)

	resp, err := c.KeysListVersions(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = KeysListVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
