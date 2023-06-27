# .env: https://www.robg3d.com/2020/05/2288/
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

default: ;

# Docker convenience recipes:
.PHONY: build
build:
	docker build \
		-t mutant:$$(date -u +"%Y%m%d%H%M%S") \
		-t mutant:latest \
		-- .
.PHONY: run
run:
	docker run \
		--rm \
		-e DATABASE_URL="$$DATABASE_URL" \
		-p 3000:8080 \
		mutant:latest

.PHONY: push
push:
	# This is really a job for CI ofc, just here for convenience.
	docker tag mutant:latest $$(terraform output --raw mutant_ecr_uri):latest
	docker push $$(terraform output --raw mutant_ecr_uri):latest

# Docker convenience recipes for testing the running container:
.PHONY: healthcheck
healthcheck:
	curl -w "%{http_code}\n" -- localhost:3000/
.PHONY: stats
stats:
	time curl -w "%{http_code}\n" -- localhost:3000/stats
.PHONY: 200
200:
	time curl \
		-H "Content-Type: application/json" \
		-X POST \
		-d '{"dna":["AAAAAA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}' \
		-w "%{http_code}\n" \
		-- localhost:3000/mutant
.PHONY: 403
403:
	time curl \
		-H "Content-Type: application/json" \
		-X POST \
		-d '{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}' \
		-w "%{http_code}\n" \
		-- localhost:3000/mutant
.PHONY: 500
500:
	time curl \
		-H "Content-Type: application/json" \
		-X POST \
		-d '' \
		-w "%{http_code}\n" \
		-- localhost:3000/mutant
.PHONY: custom
custom:
	command -v vipe
	time curl \
		-H "Content-Type: application/json" \
		-X POST \
		-d "$$(echo '{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}' | vipe)" \
		-w "%{http_code}\n" \
		-- localhost:3000/mutant
