package profilecertification

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

type ListProfileCertificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonCertification
}

type ListProfileCertificationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonCertification
}

type ListProfileCertificationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileCertificationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileCertifications ...
func (c ProfileCertificationClient) ListProfileCertifications(ctx context.Context) (result ListProfileCertificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileCertificationsCustomPager{},
		Path:       "/me/profile/certifications",
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
		Values *[]beta.PersonCertification `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileCertificationsComplete retrieves all the results into a single object
func (c ProfileCertificationClient) ListProfileCertificationsComplete(ctx context.Context) (ListProfileCertificationsCompleteResult, error) {
	return c.ListProfileCertificationsCompleteMatchingPredicate(ctx, PersonCertificationOperationPredicate{})
}

// ListProfileCertificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileCertificationClient) ListProfileCertificationsCompleteMatchingPredicate(ctx context.Context, predicate PersonCertificationOperationPredicate) (result ListProfileCertificationsCompleteResult, err error) {
	items := make([]beta.PersonCertification, 0)

	resp, err := c.ListProfileCertifications(ctx)
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

	result = ListProfileCertificationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
