tests:
	gotestsum -- . -failfast -race -coverprofile ./coverage.out
coverage:
	make tests
	go tool cover -html=coverage.out