# MELI backend challenge / mutantes

[Challenge found here](https://github.com/mauricionrgarcia/examen-mercadolibre-mutante)
(see README).

## IsMutant method

### Efficiency notes

- Simplest possible implementation that works for arbitrary matrix size and match length.
- We don't do an initial pass to collect all possible 4-tuples.
- All checks are done in place.
- We always return early as soon as possible.
- Not currently parallelized, as I didn't want to make it too fancy or Golang-specific,
  but would be trivial to run the 4 directional passes at the same time.

## REST API

- Go http server running on ECS.
  - Currently hosted at:
    - `http://mutant.emivespa.com/stats`
    - `http://mutant.emivespa.com/mutant`

### Assumptions

- We should return 200/403 regardless of whether the mutant candidate is already in the DB.

### Building and running

The Makefile contains convenience recipes for building and running the container with all the right flags,
so you can run:

- `make build` and then
- `make run`

There are also convenience recipes for hitting container endpoints:

- `make 200` for a mutant
- `make 403` for a non-mutant
- `make random` for random 6x6 dna
- `make healthcheck` for `/`
- `make stats`

Won't work without a running DB.

## DB

- Using Planetscale because setting up RDS would be a distraction.

## Known issues

- should have gone with Gorm as opposed to Prisma
- <80% test coverage.
  - ~No DB mocking means we can test the http response but code coverage tools will panic.~
    - (solved)
