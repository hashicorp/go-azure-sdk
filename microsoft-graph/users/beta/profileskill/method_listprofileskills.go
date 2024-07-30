package profileskill

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

type ListProfileSkillsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SkillProficiency
}

type ListProfileSkillsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SkillProficiency
}

type ListProfileSkillsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileSkillsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileSkills ...
func (c ProfileSkillClient) ListProfileSkills(ctx context.Context, id UserId) (result ListProfileSkillsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileSkillsCustomPager{},
		Path:       fmt.Sprintf("%s/profile/skills", id.ID()),
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
		Values *[]beta.SkillProficiency `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileSkillsComplete retrieves all the results into a single object
func (c ProfileSkillClient) ListProfileSkillsComplete(ctx context.Context, id UserId) (ListProfileSkillsCompleteResult, error) {
	return c.ListProfileSkillsCompleteMatchingPredicate(ctx, id, SkillProficiencyOperationPredicate{})
}

// ListProfileSkillsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileSkillClient) ListProfileSkillsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate SkillProficiencyOperationPredicate) (result ListProfileSkillsCompleteResult, err error) {
	items := make([]beta.SkillProficiency, 0)

	resp, err := c.ListProfileSkills(ctx, id)
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

	result = ListProfileSkillsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
