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

// https://twfhir.uasomh.uab.edu/FHIR/metadata
func GetSourceUabHuntsvilleRegionalCampus(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://twfhir.uasomh.uab.edu/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://twfhir.uasomh.uab.edu/authorization/connect/token"

	sourceDef.Audience = "https://twfhir.uasomh.uab.edu/FHIR"

	sourceDef.ApiEndpointBaseUrl = "https://twfhir.uasomh.uab.edu/FHIR"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUabHuntsvilleRegionalCampus]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "UAB-Huntsville Regional Campus"
	sourceDef.SourceType = pkg.SourceTypeUabHuntsvilleRegionalCampus
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}