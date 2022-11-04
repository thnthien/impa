package querybuilder

import "fmt"

func StandardizeQuery(driverName, query string, identifiers ...string) string {
	switch driverName {
	case "postgres":
		return postgresStandardizeQuery(query, identifiers...)
	case "mysql":
		return mysqlStandardizeQuery(query, identifiers...)
	default:
		return postgresStandardizeQuery(query, identifiers...)
	}
}

func StandardizeIdentifiers(driverName string, identifiers ...string) []string {
	ids := make([]string, 0, len(identifiers))
	switch driverName {
	case "mysql":
		for _, id := range identifiers {
			ids = append(ids, fmt.Sprintf("`%s`", id))
		}
	case "postgres":
		for _, id := range identifiers {
			ids = append(ids, fmt.Sprintf(`"%s"`, id))
		}
	default:
		for _, id := range identifiers {
			ids = append(ids, fmt.Sprintf(`"%s"`, id))
		}
	}
	return ids
}

func StandardizeIdentifier(driverName, identifier string) string {
	switch driverName {
	case "mysql":
		return "`" + identifier + "`"
	case "postgres":
		return `"` + identifier + `"`
	default:
		return `"` + identifier + `"`
	}
}

func NamedColumns(columns []string) []string {
	cols := make([]string, 0, len(columns))
	for _, col := range columns {
		cols = append(cols, ":"+col)
	}
	return cols
}

func postgresStandardizeQuery(query string, identifiers ...string) string {
	cols := make([]any, 0, len(identifiers))
	for i := range identifiers {
		cols = append(cols, fmt.Sprintf(`"%s"`, identifiers[i]))
	}

	return fmt.Sprintf(query, cols...)
}

func mysqlStandardizeQuery(query string, identifiers ...string) string {
	idts := make([]any, 0, len(identifiers))
	for i := range identifiers {
		idts = append(idts, fmt.Sprintf("`%s`", identifiers[i]))
	}

	return fmt.Sprintf(query, idts...)
}
