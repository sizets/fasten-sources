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
func GetSourceRussellvilleHoldingsLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRussellvilleHoldingsLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Russellville Holdings LLC"
	sourceDef.SourceType = pkg.SourceTypeRussellvilleHoldingsLlc
	sourceDef.Category = []string{"104100000X", "225100000X", "225X00000X", "235Z00000X", "251E00000X", "261QA1903X", "261QR1300X", "273Y00000X", "282N00000X", "314000000X", "363LF0000X"}
	sourceDef.Aliases = []string{"SAINT MARY'S REGIONAL MEDICAL CENTER REHAB", "ST . MARY'S REGIONAL MEDICAL CENTER", "ST. MARY'S REGIONAL MEDICAL CENTER", "ST. MARY'S REGIONAL MEDICAL CENTER HOME HEALTH", "VALLEY HEALTH SERVICES OF HECTOR"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1033160049", "1376595132", "1538111398", "1881646644", "1932150943"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}