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

// https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/75336/metadata
func GetSourceChildrenFirstPediatrics(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/authorize?cust=75336"
	sourceDef.TokenEndpoint = "https://fhir.prosuite.allscriptscloud.com/fhirroute/authorization/token?cust=75336"

	sourceDef.Audience = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/75336"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.prosuite.allscriptscloud.com/fhirroute/fhir/75336"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeChildrenFirstPediatrics]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Children First Pediatrics"
	sourceDef.SourceType = pkg.SourceTypeChildrenFirstPediatrics
	sourceDef.Category = []string{"1041C0700X", "208000000X"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1144739178", "1831131556"}}
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}