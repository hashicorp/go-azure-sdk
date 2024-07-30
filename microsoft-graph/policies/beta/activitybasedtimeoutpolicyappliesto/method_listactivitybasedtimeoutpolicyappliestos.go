package activitybasedtimeoutpolicyappliesto

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

type ListActivityBasedTimeoutPolicyAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListActivityBasedTimeoutPolicyAppliesTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListActivityBasedTimeoutPolicyAppliesTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListActivityBasedTimeoutPolicyAppliesTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListActivityBasedTimeoutPolicyAppliesTos ...
func (c ActivityBasedTimeoutPolicyAppliesToClient) ListActivityBasedTimeoutPolicyAppliesTos(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId) (result ListActivityBasedTimeoutPolicyAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListActivityBasedTimeoutPolicyAppliesTosCustomPager{},
		Path:       fmt.Sprintf("%s/appliesTo", id.ID()),
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListActivityBasedTimeoutPolicyAppliesTosComplete retrieves all the results into a single object
func (c ActivityBasedTimeoutPolicyAppliesToClient) ListActivityBasedTimeoutPolicyAppliesTosComplete(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId) (ListActivityBasedTimeoutPolicyAppliesTosCompleteResult, error) {
	return c.ListActivityBasedTimeoutPolicyAppliesTosCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListActivityBasedTimeoutPolicyAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActivityBasedTimeoutPolicyAppliesToClient) ListActivityBasedTimeoutPolicyAppliesTosCompleteMatchingPredicate(ctx context.Context, id PolicyActivityBasedTimeoutPolicyId, predicate DirectoryObjectOperationPredicate) (result ListActivityBasedTimeoutPolicyAppliesTosCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListActivityBasedTimeoutPolicyAppliesTos(ctx, id)
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

	result = ListActivityBasedTimeoutPolicyAppliesTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
