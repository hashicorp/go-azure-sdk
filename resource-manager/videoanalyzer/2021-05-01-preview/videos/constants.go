package videos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyEccAlgo string

const (
	AccessPolicyEccAlgoESFiveOneTwo     AccessPolicyEccAlgo = "ES512"
	AccessPolicyEccAlgoESThreeEightFour AccessPolicyEccAlgo = "ES384"
	AccessPolicyEccAlgoESTwoFiveSix     AccessPolicyEccAlgo = "ES256"
)

func PossibleValuesForAccessPolicyEccAlgo() []string {
	return []string{
		string(AccessPolicyEccAlgoESFiveOneTwo),
		string(AccessPolicyEccAlgoESThreeEightFour),
		string(AccessPolicyEccAlgoESTwoFiveSix),
	}
}

type AccessPolicyRole string

const (
	AccessPolicyRoleReader AccessPolicyRole = "Reader"
)

func PossibleValuesForAccessPolicyRole() []string {
	return []string{
		string(AccessPolicyRoleReader),
	}
}

type AccessPolicyRsaAlgo string

const (
	AccessPolicyRsaAlgoRSFiveOneTwo     AccessPolicyRsaAlgo = "RS512"
	AccessPolicyRsaAlgoRSThreeEightFour AccessPolicyRsaAlgo = "RS384"
	AccessPolicyRsaAlgoRSTwoFiveSix     AccessPolicyRsaAlgo = "RS256"
)

func PossibleValuesForAccessPolicyRsaAlgo() []string {
	return []string{
		string(AccessPolicyRsaAlgoRSFiveOneTwo),
		string(AccessPolicyRsaAlgoRSThreeEightFour),
		string(AccessPolicyRsaAlgoRSTwoFiveSix),
	}
}

type VideoType string

const (
	VideoTypeArchive VideoType = "Archive"
)

func PossibleValuesForVideoType() []string {
	return []string{
		string(VideoTypeArchive),
	}
}
