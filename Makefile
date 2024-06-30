dev:
	MODE=development go run main.go run
generate:
	go run main.go generate
preview:
	go run main.go preview
css:
	npm i -g sass
	sass --watch internal/styles:static/css --no-source-map
dashboard:
	MODE=development go run main.go dash



