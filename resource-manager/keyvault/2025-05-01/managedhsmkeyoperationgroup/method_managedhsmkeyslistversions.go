package managedhsmkeyoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedHsmKeysListVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedHsmKey
}

type ManagedHsmKeysListVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedHsmKey
}

type ManagedHsmKeysListVersionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ManagedHsmKeysListVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ManagedHsmKeysListVersions ...
func (c ManagedHsmKeyOperationGroupClient) ManagedHsmKeysListVersions(ctx context.Context, id KeyId) (result ManagedHsmKeysListVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ManagedHsmKeysListVersionsCustomPager{},
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
		Values *[]ManagedHsmKey `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ManagedHsmKeysListVersionsComplete retrieves all the results into a single object
func (c ManagedHsmKeyOperationGroupClient) ManagedHsmKeysListVersionsComplete(ctx context.Context, id KeyId) (ManagedHsmKeysListVersionsCompleteResult, error) {
	return c.ManagedHsmKeysListVersionsCompleteMatchingPredicate(ctx, id, ManagedHsmKeyOperationPredicate{})
}

// ManagedHsmKeysListVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedHsmKeyOperationGroupClient) ManagedHsmKeysListVersionsCompleteMatchingPredicate(ctx context.Context, id KeyId, predicate ManagedHsmKeyOperationPredicate) (result ManagedHsmKeysListVersionsCompleteResult, err error) {
	items := make([]ManagedHsmKey, 0)

	resp, err := c.ManagedHsmKeysListVersions(ctx, id)
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

	result = ManagedHsmKeysListVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
