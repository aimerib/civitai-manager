TAILWIND_CLI := ./bin/tailwindcss

.PHONY: dev build clean install

dev: install
		make -j 2 run-tailwind run-buffalo

run-tailwind:
		$(TAILWIND_CLI) -i ./assets/css/tailwind.css -o ./public/assets/main.css --watch

run-buffalo:
		buffalo dev

.PHONY: dev run-tailwind run-buffalo

build: install
				$(TAILWIND_CLI) -i ./assets/css/tailwind.css -o ./public/assets/main.css --minify
				buffalo build

clean:
		rm -rf tmp
		rm -rf bin
		rm public/assets/main.css

install:
		go mod tidy
		# Create bin directory if it doesn't exist
		mkdir -p bin
		# Download Tailwind CLI if not present
		if [ ! -f $(TAILWIND_CLI) ]; then \
				curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64; \
				mv tailwindcss-macos-arm64 $(TAILWIND_CLI); \
				chmod +x $(TAILWIND_CLI); \
		fi
