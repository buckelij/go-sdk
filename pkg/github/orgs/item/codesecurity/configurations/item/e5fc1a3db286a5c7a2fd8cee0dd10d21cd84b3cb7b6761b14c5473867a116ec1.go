package item
// The enablement status of private vulnerability reporting
type WithConfiguration_PatchRequestBody_private_vulnerability_reporting int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING WithConfiguration_PatchRequestBody_private_vulnerability_reporting = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
)

func (i WithConfiguration_PatchRequestBody_private_vulnerability_reporting) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_private_vulnerability_reporting(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_PRIVATE_VULNERABILITY_REPORTING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_private_vulnerability_reporting(values []WithConfiguration_PatchRequestBody_private_vulnerability_reporting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_private_vulnerability_reporting) isMultiValue() bool {
    return false
}
