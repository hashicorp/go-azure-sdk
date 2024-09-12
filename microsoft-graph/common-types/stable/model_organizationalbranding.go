package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OrganizationalBrandingProperties = OrganizationalBranding{}

type OrganizationalBranding struct {
	// Add different branding based on a locale.
	Localizations *[]OrganizationalBrandingLocalization `json:"localizations,omitempty"`

	// Fields inherited from OrganizationalBrandingProperties

	// Color that appears in place of the background image in low-bandwidth connections. We recommend that you use the
	// primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white
	// is #FFFFFF.
	BackgroundColor nullable.Type[string] `json:"backgroundColor,omitempty"`

	// Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB
	// and not larger than 1920 × 1080 pixels. A smaller image reduces bandwidth requirements and make the page load
	// faster.
	BackgroundImage nullable.Type[string] `json:"backgroundImage,omitempty"`

	// A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the
	// version served by a CDN. Read-only.
	BackgroundImageRelativeUrl nullable.Type[string] `json:"backgroundImageRelativeUrl,omitempty"`

	// A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger
	// than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
	BannerLogo nullable.Type[string] `json:"bannerLogo,omitempty"`

	// A relative URL for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the
	// read-only version served by a CDN. Read-only.
	BannerLogoRelativeUrl nullable.Type[string] `json:"bannerLogoRelativeUrl,omitempty"`

	// A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN
	// providers are used at the same time for high availability of read requests. Read-only.
	CdnList *[]string `json:"cdnList,omitempty"`

	// Represents the content options to be customized throughout the authentication flow for a tenant. NOTE: Supported by
	// Microsoft Entra External ID in external tenants only.
	ContentCustomization *ContentCustomization `json:"contentCustomization,omitempty"`

	// A custom URL for resetting account credentials. This URL must be in ASCII format or non-ASCII characters must be URL
	// encoded, and not exceed 128 characters.
	CustomAccountResetCredentialsUrl nullable.Type[string] `json:"customAccountResetCredentialsUrl,omitempty"`

	// CSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
	CustomCSS nullable.Type[string] `json:"customCSS,omitempty"`

	// A relative URL for the customCSS property that is combined with a CDN base URL from the cdnList to provide the
	// version served by a CDN. Read-only.
	CustomCSSRelativeUrl nullable.Type[string] `json:"customCSSRelativeUrl,omitempty"`

	// A string to replace the default 'Can't access your account?' self-service password reset (SSPR) hyperlink text on the
	// sign-in page. This text must be in Unicode format and not exceed 256 characters.
	CustomCannotAccessYourAccountText nullable.Type[string] `json:"customCannotAccessYourAccountText,omitempty"`

	// A custom URL to replace the default URL of the self-service password reset (SSPR) 'Can't access your account?'
	// hyperlink on the sign-in page. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not
	// exceed 128 characters. DO NOT USE. Use customAccountResetCredentialsUrl instead.
	CustomCannotAccessYourAccountUrl nullable.Type[string] `json:"customCannotAccessYourAccountUrl,omitempty"`

	// A string to replace the default 'Forgot my password' hyperlink text on the sign-in form. This text must be in Unicode
	// format and not exceed 256 characters.
	CustomForgotMyPasswordText nullable.Type[string] `json:"customForgotMyPasswordText,omitempty"`

	// A string to replace the default 'Privacy and Cookies' hyperlink text in the footer. This text must be in Unicode
	// format and not exceed 256 characters.
	CustomPrivacyAndCookiesText nullable.Type[string] `json:"customPrivacyAndCookiesText,omitempty"`

	// A custom URL to replace the default URL of the 'Privacy and Cookies' hyperlink in the footer. This URL must be in
	// ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
	CustomPrivacyAndCookiesUrl nullable.Type[string] `json:"customPrivacyAndCookiesUrl,omitempty"`

	// A string to replace the default 'reset it now' hyperlink text on the sign-in form. This text must be in Unicode
	// format and not exceed 256 characters. DO NOT USE: Customization of the 'reset it now' hyperlink text is currently not
	// supported.
	CustomResetItNowText nullable.Type[string] `json:"customResetItNowText,omitempty"`

	// A string to replace the the default 'Terms of Use' hyperlink text in the footer. This text must be in Unicode format
	// and not exceed 256 characters.
	CustomTermsOfUseText nullable.Type[string] `json:"customTermsOfUseText,omitempty"`

	// A custom URL to replace the default URL of the 'Terms of Use' hyperlink in the footer. This URL must be in ASCII
	// format or non-ASCII characters must be URL encoded, and not exceed 128characters.
	CustomTermsOfUseUrl nullable.Type[string] `json:"customTermsOfUseUrl,omitempty"`

	// A custom icon (favicon) to replace a default Microsoft product favicon on a Microsoft Entra tenant.
	Favicon nullable.Type[string] `json:"favicon,omitempty"`

	// A relative url for the favicon above that is combined with a CDN base URL from the cdnList to provide the version
	// served by a CDN. Read-only.
	FaviconRelativeUrl nullable.Type[string] `json:"faviconRelativeUrl,omitempty"`

	// The RGB color to apply to customize the color of the header.
	HeaderBackgroundColor nullable.Type[string] `json:"headerBackgroundColor,omitempty"`

	// A company logo that appears in the header of the sign-in page. The allowed types are PNG or JPEG not larger than 36
	// × 245 pixels. We recommend using a transparent image with no padding around the logo.
	HeaderLogo nullable.Type[string] `json:"headerLogo,omitempty"`

	// A relative URL for the headerLogo property that is combined with a CDN base URL from the cdnList to provide the
	// read-only version served by a CDN. Read-only.
	HeaderLogoRelativeUrl nullable.Type[string] `json:"headerLogoRelativeUrl,omitempty"`

	// Represents the layout configuration to be displayed on the login page for a tenant.
	LoginPageLayoutConfiguration *LoginPageLayoutConfiguration `json:"loginPageLayoutConfiguration,omitempty"`

	// Represents the various texts that can be hidden on the login page for a tenant.
	LoginPageTextVisibilitySettings *LoginPageTextVisibilitySettings `json:"loginPageTextVisibilitySettings,omitempty"`

	// Text that appears at the bottom of the sign-in box. Use this to communicate additional information, such as the phone
	// number to your help desk or a legal statement. This text must be in Unicode format and not exceed 1024 characters.
	SignInPageText nullable.Type[string] `json:"signInPageText,omitempty"`

	// A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows
	// Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than
	// 10 KB in size. We recommend using a transparent image with no padding around the logo.
	SquareLogo nullable.Type[string] `json:"squareLogo,omitempty"`

	// A square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows
	// Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than
	// 10 KB in size. We recommend using a transparent image with no padding around the logo.
	SquareLogoDark nullable.Type[string] `json:"squareLogoDark,omitempty"`

	// A relative URL for the squareLogoDark property that is combined with a CDN base URL from the cdnList to provide the
	// version served by a CDN. Read-only.
	SquareLogoDarkRelativeUrl nullable.Type[string] `json:"squareLogoDarkRelativeUrl,omitempty"`

	// A relative URL for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the
	// version served by a CDN. Read-only.
	SquareLogoRelativeUrl nullable.Type[string] `json:"squareLogoRelativeUrl,omitempty"`

	// A string that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without
	// links or code, and can't exceed 64 characters.
	UsernameHintText nullable.Type[string] `json:"usernameHintText,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s OrganizationalBranding) OrganizationalBrandingProperties() BaseOrganizationalBrandingPropertiesImpl {
	return BaseOrganizationalBrandingPropertiesImpl{
		BackgroundColor:                   s.BackgroundColor,
		BackgroundImage:                   s.BackgroundImage,
		BackgroundImageRelativeUrl:        s.BackgroundImageRelativeUrl,
		BannerLogo:                        s.BannerLogo,
		BannerLogoRelativeUrl:             s.BannerLogoRelativeUrl,
		CdnList:                           s.CdnList,
		ContentCustomization:              s.ContentCustomization,
		CustomAccountResetCredentialsUrl:  s.CustomAccountResetCredentialsUrl,
		CustomCSS:                         s.CustomCSS,
		CustomCSSRelativeUrl:              s.CustomCSSRelativeUrl,
		CustomCannotAccessYourAccountText: s.CustomCannotAccessYourAccountText,
		CustomCannotAccessYourAccountUrl:  s.CustomCannotAccessYourAccountUrl,
		CustomForgotMyPasswordText:        s.CustomForgotMyPasswordText,
		CustomPrivacyAndCookiesText:       s.CustomPrivacyAndCookiesText,
		CustomPrivacyAndCookiesUrl:        s.CustomPrivacyAndCookiesUrl,
		CustomResetItNowText:              s.CustomResetItNowText,
		CustomTermsOfUseText:              s.CustomTermsOfUseText,
		CustomTermsOfUseUrl:               s.CustomTermsOfUseUrl,
		Favicon:                           s.Favicon,
		FaviconRelativeUrl:                s.FaviconRelativeUrl,
		HeaderBackgroundColor:             s.HeaderBackgroundColor,
		HeaderLogo:                        s.HeaderLogo,
		HeaderLogoRelativeUrl:             s.HeaderLogoRelativeUrl,
		LoginPageLayoutConfiguration:      s.LoginPageLayoutConfiguration,
		LoginPageTextVisibilitySettings:   s.LoginPageTextVisibilitySettings,
		SignInPageText:                    s.SignInPageText,
		SquareLogo:                        s.SquareLogo,
		SquareLogoDark:                    s.SquareLogoDark,
		SquareLogoDarkRelativeUrl:         s.SquareLogoDarkRelativeUrl,
		SquareLogoRelativeUrl:             s.SquareLogoRelativeUrl,
		UsernameHintText:                  s.UsernameHintText,
		Id:                                s.Id,
		ODataId:                           s.ODataId,
		ODataType:                         s.ODataType,
	}
}

func (s OrganizationalBranding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OrganizationalBranding{}

func (s OrganizationalBranding) MarshalJSON() ([]byte, error) {
	type wrapper OrganizationalBranding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrganizationalBranding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrganizationalBranding: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.organizationalBranding"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrganizationalBranding: %+v", err)
	}

	return encoded, nil
}
