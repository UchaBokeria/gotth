run:
	@rm -rf build && mkdir -p ./build/view && chmod -R 777 ./public
	make -j3 tailwind-watch templ-watch runAir

.PHONY: build
build:
	make templ-build tailwind-build vet staticcheck test
	go build -ldflags "-X main.Environment=production" -o ./bin/app ./cmd/app/main.go

migrate:
	go run ./cmd/migrate/main.go

seed:
	go run ./cmd/seed/main.go

oneTime:
	templ generate && make templ-clean && go build -o build/app ./cmd/app && ./build/app

tailwind-build:
	npx tailwindcss-cli -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css

tailwind-watch:
	npx tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --watch


# Templ
templ-build:
	templ generate -keep-orphaned-files
	make templ-clean

templ-watch:
	@templ generate -keep-orphaned-files --watch

templ-clean:
	@find ./server/view/ -type f \( -name "*_templ.go" -o -name "*_templ.txt" \) -exec mv -t ./build/view {} +

reload:
	go run ./cmd/utils/reload.go

# Testing
vet:
	go vet ./...

staticcheck: 
	staticcheck ./...

test:
	go test -race -v -timeout 30s ./...
    
runAir:
	air