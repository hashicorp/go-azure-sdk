package featurerolloutpolicyappliesto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListFeatureRolloutPolicyAppliesToRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListFeatureRolloutPolicyAppliesToRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListFeatureRolloutPolicyAppliesToRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFeatureRolloutPolicyAppliesToRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFeatureRolloutPolicyAppliesToRefs ...
func (c FeatureRolloutPolicyAppliesToClient) ListFeatureRolloutPolicyAppliesToRefs(ctx context.Context, id PolicyFeatureRolloutPolicyId) (result ListFeatureRolloutPolicyAppliesToRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListFeatureRolloutPolicyAppliesToRefsCustomPager{},
		Path:       fmt.Sprintf("%s/appliesTo/$ref", id.ID()),
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListFeatureRolloutPolicyAppliesToRefsComplete retrieves all the results into a single object
func (c FeatureRolloutPolicyAppliesToClient) ListFeatureRolloutPolicyAppliesToRefsComplete(ctx context.Context, id PolicyFeatureRolloutPolicyId) (result ListFeatureRolloutPolicyAppliesToRefsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListFeatureRolloutPolicyAppliesToRefs(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListFeatureRolloutPolicyAppliesToRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
