package termsandconditiongroupassignment

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

type ListTermsAndConditionGroupAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TermsAndConditionsGroupAssignment
}

type ListTermsAndConditionGroupAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TermsAndConditionsGroupAssignment
}

type ListTermsAndConditionGroupAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsAndConditionGroupAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsAndConditionGroupAssignments ...
func (c TermsAndConditionGroupAssignmentClient) ListTermsAndConditionGroupAssignments(ctx context.Context, id DeviceManagementTermsAndConditionId) (result ListTermsAndConditionGroupAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsAndConditionGroupAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/groupAssignments", id.ID()),
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
		Values *[]beta.TermsAndConditionsGroupAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsAndConditionGroupAssignmentsComplete retrieves all the results into a single object
func (c TermsAndConditionGroupAssignmentClient) ListTermsAndConditionGroupAssignmentsComplete(ctx context.Context, id DeviceManagementTermsAndConditionId) (ListTermsAndConditionGroupAssignmentsCompleteResult, error) {
	return c.ListTermsAndConditionGroupAssignmentsCompleteMatchingPredicate(ctx, id, TermsAndConditionsGroupAssignmentOperationPredicate{})
}

// ListTermsAndConditionGroupAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsAndConditionGroupAssignmentClient) ListTermsAndConditionGroupAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementTermsAndConditionId, predicate TermsAndConditionsGroupAssignmentOperationPredicate) (result ListTermsAndConditionGroupAssignmentsCompleteResult, err error) {
	items := make([]beta.TermsAndConditionsGroupAssignment, 0)

	resp, err := c.ListTermsAndConditionGroupAssignments(ctx, id)
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

	result = ListTermsAndConditionGroupAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
