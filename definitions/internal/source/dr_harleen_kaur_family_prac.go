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

// https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73026/metadata
func GetSourceDrHarleenKaurFamilyPrac(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/authorize?cust=73026"
	sourceDef.TokenEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/token?cust=73026"

	sourceDef.Audience = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73026"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/73026"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDrHarleenKaurFamilyPrac]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Dr Harleen Kaur Family Prac"
	sourceDef.SourceType = pkg.SourceTypeDrHarleenKaurFamilyPrac
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}