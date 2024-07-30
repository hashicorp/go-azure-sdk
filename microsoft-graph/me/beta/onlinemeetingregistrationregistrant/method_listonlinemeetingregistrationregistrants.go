package onlinemeetingregistrationregistrant

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

type ListOnlineMeetingRegistrationRegistrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MeetingRegistrantBase
}

type ListOnlineMeetingRegistrationRegistrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MeetingRegistrantBase
}

type ListOnlineMeetingRegistrationRegistrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingRegistrationRegistrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingRegistrationRegistrants ...
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrants(ctx context.Context, id MeOnlineMeetingId) (result ListOnlineMeetingRegistrationRegistrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOnlineMeetingRegistrationRegistrantsCustomPager{},
		Path:       fmt.Sprintf("%s/registration/registrants", id.ID()),
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
		Values *[]beta.MeetingRegistrantBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnlineMeetingRegistrationRegistrantsComplete retrieves all the results into a single object
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrantsComplete(ctx context.Context, id MeOnlineMeetingId) (ListOnlineMeetingRegistrationRegistrantsCompleteResult, error) {
	return c.ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate(ctx, id, MeetingRegistrantBaseOperationPredicate{})
}

// ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate MeetingRegistrantBaseOperationPredicate) (result ListOnlineMeetingRegistrationRegistrantsCompleteResult, err error) {
	items := make([]beta.MeetingRegistrantBase, 0)

	resp, err := c.ListOnlineMeetingRegistrationRegistrants(ctx, id)
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

	result = ListOnlineMeetingRegistrationRegistrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
