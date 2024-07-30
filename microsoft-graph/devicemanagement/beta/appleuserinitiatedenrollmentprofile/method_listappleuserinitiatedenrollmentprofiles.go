package appleuserinitiatedenrollmentprofile

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

type ListAppleUserInitiatedEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppleUserInitiatedEnrollmentProfile
}

type ListAppleUserInitiatedEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppleUserInitiatedEnrollmentProfile
}

type ListAppleUserInitiatedEnrollmentProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppleUserInitiatedEnrollmentProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppleUserInitiatedEnrollmentProfiles ...
func (c AppleUserInitiatedEnrollmentProfileClient) ListAppleUserInitiatedEnrollmentProfiles(ctx context.Context) (result ListAppleUserInitiatedEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppleUserInitiatedEnrollmentProfilesCustomPager{},
		Path:       "/deviceManagement/appleUserInitiatedEnrollmentProfiles",
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
		Values *[]beta.AppleUserInitiatedEnrollmentProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppleUserInitiatedEnrollmentProfilesComplete retrieves all the results into a single object
func (c AppleUserInitiatedEnrollmentProfileClient) ListAppleUserInitiatedEnrollmentProfilesComplete(ctx context.Context) (ListAppleUserInitiatedEnrollmentProfilesCompleteResult, error) {
	return c.ListAppleUserInitiatedEnrollmentProfilesCompleteMatchingPredicate(ctx, AppleUserInitiatedEnrollmentProfileOperationPredicate{})
}

// ListAppleUserInitiatedEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppleUserInitiatedEnrollmentProfileClient) ListAppleUserInitiatedEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, predicate AppleUserInitiatedEnrollmentProfileOperationPredicate) (result ListAppleUserInitiatedEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.AppleUserInitiatedEnrollmentProfile, 0)

	resp, err := c.ListAppleUserInitiatedEnrollmentProfiles(ctx)
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

	result = ListAppleUserInitiatedEnrollmentProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
