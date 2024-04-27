.PHONY: run
run:
	@rm -rf build && mkdir -p ./build/view && chmod -R 777 ./public
	make -j3 tailwind-watch templ-watch runAir

.PHONY: reload
reload:
	go run ./cmd/utils/reload.go

.PHONY: prod
prod:
	templ generate && make templ-clean && go build -o build/app ./cmd/app && ./build/app

.PHONY: build
build:
	make templ-build tailwind vet staticcheck test
	go build -ldflags "-X main.Environment=production" -o ./bin/app ./cmd/app/main.go



.PHONY: migrate
migrate:
	go run ./cmd/migrate/main.go

.PHONY: seed
seed:
	go run ./cmd/seed/main.go



.PHONY: tailwind
tailwind:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --minify

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --watch



.PHONY: templ
templ:
	templ generate -keep-orphaned-files
	make templ-clean

.PHONY: templ-watch
templ-watch:
	@templ generate -keep-orphaned-files --watch

.PHONY: templ-clean
templ-clean:
	@find ./server/view/ -type f \( -name "*_templ.go" -o -name "*_templ.txt" \) -exec mv -t ./build/view {} +



.PHONY: vet
vet:
	go vet ./...
    
.PHONY: runAir
runAir:
	air
	
.PHONY: test
test:
	go test -race -v -timeout 30s ./...

.PHONY: staticcheck
staticcheck: 
	staticcheck ./...