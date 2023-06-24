dbup:
	docker-compose up -d

build:
  go build -o purl ./cmd/purl

