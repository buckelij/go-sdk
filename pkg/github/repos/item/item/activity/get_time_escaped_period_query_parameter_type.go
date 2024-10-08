package activity
type GetTime_periodQueryParameterType int

const (
    DAY_GETTIME_PERIODQUERYPARAMETERTYPE GetTime_periodQueryParameterType = iota
    WEEK_GETTIME_PERIODQUERYPARAMETERTYPE
    MONTH_GETTIME_PERIODQUERYPARAMETERTYPE
    QUARTER_GETTIME_PERIODQUERYPARAMETERTYPE
    YEAR_GETTIME_PERIODQUERYPARAMETERTYPE
)

func (i GetTime_periodQueryParameterType) String() string {
    return []string{"day", "week", "month", "quarter", "year"}[i]
}
func ParseGetTime_periodQueryParameterType(v string) (any, error) {
    result := DAY_GETTIME_PERIODQUERYPARAMETERTYPE
    switch v {
        case "day":
            result = DAY_GETTIME_PERIODQUERYPARAMETERTYPE
        case "week":
            result = WEEK_GETTIME_PERIODQUERYPARAMETERTYPE
        case "month":
            result = MONTH_GETTIME_PERIODQUERYPARAMETERTYPE
        case "quarter":
            result = QUARTER_GETTIME_PERIODQUERYPARAMETERTYPE
        case "year":
            result = YEAR_GETTIME_PERIODQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetTime_periodQueryParameterType(values []GetTime_periodQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTime_periodQueryParameterType) isMultiValue() bool {
    return false
}
