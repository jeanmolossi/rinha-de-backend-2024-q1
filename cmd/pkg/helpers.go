package pkg

func JSONErr(msg string, err error) string {
	if err == nil {
		return `{ "message": "` + msg + `", "error": "null" }`
	}

	return `{ "message": "` + msg + `", "error": "` + err.Error() + `" }`
}
