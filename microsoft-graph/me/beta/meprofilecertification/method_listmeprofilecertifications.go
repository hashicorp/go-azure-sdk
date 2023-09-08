package meprofilecertification

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeProfileCertificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonCertificationCollectionResponse
}

type ListMeProfileCertificationsCompleteResult struct {
	Items []models.PersonCertificationCollectionResponse
}

// ListMeProfileCertifications ...
func (c MeProfileCertificationClient) ListMeProfileCertifications(ctx context.Context) (result ListMeProfileCertificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.PersonCertificationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeProfileCertificationsComplete retrieves all the results into a single object
func (c MeProfileCertificationClient) ListMeProfileCertificationsComplete(ctx context.Context) (ListMeProfileCertificationsCompleteResult, error) {
	return c.ListMeProfileCertificationsCompleteMatchingPredicate(ctx, models.PersonCertificationCollectionResponseOperationPredicate{})
}

// ListMeProfileCertificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeProfileCertificationClient) ListMeProfileCertificationsCompleteMatchingPredicate(ctx context.Context, predicate models.PersonCertificationCollectionResponseOperationPredicate) (result ListMeProfileCertificationsCompleteResult, err error) {
	items := make([]models.PersonCertificationCollectionResponse, 0)

	resp, err := c.ListMeProfileCertifications(ctx)
	if err != nil {
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

	result = ListMeProfileCertificationsCompleteResult{
		Items: items,
	}
	return
}
