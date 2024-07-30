package onlinemeetingregistrationcustomquestion

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

type ListOnlineMeetingRegistrationCustomQuestionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MeetingRegistrationQuestion
}

type ListOnlineMeetingRegistrationCustomQuestionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MeetingRegistrationQuestion
}

type ListOnlineMeetingRegistrationCustomQuestionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingRegistrationCustomQuestionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingRegistrationCustomQuestions ...
func (c OnlineMeetingRegistrationCustomQuestionClient) ListOnlineMeetingRegistrationCustomQuestions(ctx context.Context, id MeOnlineMeetingId) (result ListOnlineMeetingRegistrationCustomQuestionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingRegistrationCustomQuestionsCustomPager{},
		Path:       fmt.Sprintf("%s/registration/customQuestions", id.ID()),
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
		Values *[]beta.MeetingRegistrationQuestion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingRegistrationCustomQuestionsComplete retrieves all the results into a single object
func (c OnlineMeetingRegistrationCustomQuestionClient) ListOnlineMeetingRegistrationCustomQuestionsComplete(ctx context.Context, id MeOnlineMeetingId) (ListOnlineMeetingRegistrationCustomQuestionsCompleteResult, error) {
	return c.ListOnlineMeetingRegistrationCustomQuestionsCompleteMatchingPredicate(ctx, id, MeetingRegistrationQuestionOperationPredicate{})
}

// ListOnlineMeetingRegistrationCustomQuestionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingRegistrationCustomQuestionClient) ListOnlineMeetingRegistrationCustomQuestionsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate MeetingRegistrationQuestionOperationPredicate) (result ListOnlineMeetingRegistrationCustomQuestionsCompleteResult, err error) {
	items := make([]beta.MeetingRegistrationQuestion, 0)

	resp, err := c.ListOnlineMeetingRegistrationCustomQuestions(ctx, id)
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

	result = ListOnlineMeetingRegistrationCustomQuestionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
