// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://api.platform.athenahealth.com/fhir/r4/metadata
func GetSourceKibbeOrthopedicsPllc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeKibbeOrthopedicsPllc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "KIBBE ORTHOPEDICS PLLC"
	sourceDef.SourceType = pkg.SourceTypeKibbeOrthopedicsPllc
	sourceDef.Category = []string{"207X00000X", "207XX0005X", "332B00000X"}
	sourceDef.Aliases = []string{"KENTUCKY ORTHOPEDICS & SPORTS MEDICINE", "KIBBE ORTHOPEDICS, PLLC"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1063192847", "1073252375"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}