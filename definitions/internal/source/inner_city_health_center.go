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
func GetSourceInnerCityHealthCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeInnerCityHealthCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Inner City Health Center"
	sourceDef.SourceType = pkg.SourceTypeInnerCityHealthCenter
	sourceDef.Category = []string{"122300000X", "1223D0001X", "124Q00000X", "261QD0000X", "261QF0400X", "261QH0100X"}
	sourceDef.Aliases = []string{"INNER CITY HEALTH CENTER AT WHEATRIDGE", "NEW HOPE", "WHEAT RIDGE AT INNER CITY HEALTH CENTER", "WHEAT RIDGE FAMILY HEALTH CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003078742", "1013491208", "1558964734", "1578539979", "1811159551"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}