package sql

const (
	queryGetUsersByCountry = "SELECT * FROM users WHERE country IN(%s);"
)

func GetUsers(countries []string) {
	// params is going to hold different placeholders
	params := ""

	// args is going to hold each argument required by stmt.Query()
	args := make([]interface{}, len(countries))

	// populate both params & args using a single for loop:
	for i := 0; i < len(countries); i++ {
		args[i] = countries[i]
		if i < len(countries)-1 {
			params += "?,"
		} else {
			params += "?"
		}
	}

	//query := fmt.Sprintf(queryGetUsersByCountry, params)

	// Once processed, the query should have as many placeholders as countries received in the input array.
	/*
		stmt, err := db.Prepare(query)
		if err != nil {
			return
		}
		defer stmt.Close()
	*/
	// Now that we have the built statement, we can now use the args slice created before to query as needed.
	/*
		rows, err := stmt.Query(args...)
		if err != nil {
			return
		}
		defer rows.Close()

		for rows.Next() {
			// Process each result in here.
		}
	*/
}
