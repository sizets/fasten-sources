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

// https://fhirprod.montefiore.org/FHIR/metadata
func GetSourceMontefioreNewRochelleHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhirprod.montefiore.org/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://fhirprod.montefiore.org/authorization/connect/token"

	sourceDef.Audience = "https://fhirprod.montefiore.org/FHIR"

	sourceDef.ApiEndpointBaseUrl = "https://fhirprod.montefiore.org/FHIR"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMontefioreNewRochelleHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Montefiore New Rochelle Hospital"
	sourceDef.SourceType = pkg.SourceTypeMontefioreNewRochelleHospital
	sourceDef.Category = []string{"208D00000X", "282N00000X", "207RX0202X"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1790108439", "1477976264", "1245654656", "1932523479", "1033534268", "1720414154", "1497179303", "1093138885", "1295163244"}}
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}