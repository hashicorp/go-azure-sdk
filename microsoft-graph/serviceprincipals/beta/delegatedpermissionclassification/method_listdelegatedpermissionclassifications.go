package delegatedpermissionclassification

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

type ListDelegatedPermissionClassificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DelegatedPermissionClassification
}

type ListDelegatedPermissionClassificationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DelegatedPermissionClassification
}

type ListDelegatedPermissionClassificationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDelegatedPermissionClassificationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDelegatedPermissionClassifications ...
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassifications(ctx context.Context, id ServicePrincipalId) (result ListDelegatedPermissionClassificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDelegatedPermissionClassificationsCustomPager{},
		Path:       fmt.Sprintf("%s/delegatedPermissionClassifications", id.ID()),
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
		Values *[]beta.DelegatedPermissionClassification `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDelegatedPermissionClassificationsComplete retrieves all the results into a single object
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassificationsComplete(ctx context.Context, id ServicePrincipalId) (ListDelegatedPermissionClassificationsCompleteResult, error) {
	return c.ListDelegatedPermissionClassificationsCompleteMatchingPredicate(ctx, id, DelegatedPermissionClassificationOperationPredicate{})
}

// ListDelegatedPermissionClassificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassificationsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate DelegatedPermissionClassificationOperationPredicate) (result ListDelegatedPermissionClassificationsCompleteResult, err error) {
	items := make([]beta.DelegatedPermissionClassification, 0)

	resp, err := c.ListDelegatedPermissionClassifications(ctx, id)
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

	result = ListDelegatedPermissionClassificationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
