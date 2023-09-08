package userprofilecertification

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

type ListUserByIdProfileCertificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PersonCertificationCollectionResponse
}

type ListUserByIdProfileCertificationsCompleteResult struct {
	Items []models.PersonCertificationCollectionResponse
}

// ListUserByIdProfileCertifications ...
func (c UserProfileCertificationClient) ListUserByIdProfileCertifications(ctx context.Context, id UserId) (result ListUserByIdProfileCertificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/profile/certifications", id.ID()),
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

// ListUserByIdProfileCertificationsComplete retrieves all the results into a single object
func (c UserProfileCertificationClient) ListUserByIdProfileCertificationsComplete(ctx context.Context, id UserId) (ListUserByIdProfileCertificationsCompleteResult, error) {
	return c.ListUserByIdProfileCertificationsCompleteMatchingPredicate(ctx, id, models.PersonCertificationCollectionResponseOperationPredicate{})
}

// ListUserByIdProfileCertificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserProfileCertificationClient) ListUserByIdProfileCertificationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PersonCertificationCollectionResponseOperationPredicate) (result ListUserByIdProfileCertificationsCompleteResult, err error) {
	items := make([]models.PersonCertificationCollectionResponse, 0)

	resp, err := c.ListUserByIdProfileCertifications(ctx, id)
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

	result = ListUserByIdProfileCertificationsCompleteResult{
		Items: items,
	}
	return
}
