# MELI backend challenge / mutantes

[Challenge found here](https://github.com/mauricionrgarcia/examen-mercadolibre-mutante)
(see README).

## mutant.IsMutant method

### Efficiency notes

- We don't do an initial pass to collect al possible n-tuples.
- ~All checks are done in place.~
  - Excuse me, all checks _could_ be done in place,
    if we were ok with replacing the `checkLine` function with a fixed-length check like this:
    ```go
    if [i][j] == [i][j+1] && [i][j] == [i][j+2] && [i][j] == [i][j+2]
    ```
- We always return early as soon as possible.
  - Not that it matters all that much, the human test cases, where we go through all possible
    tuples, are still pretty fast.

## REST API

- Go http server running on ECS
  - Might convert it to a Lambda soon.

Currently hosted at

- `http://mutant.emivespa.com/`
- `http://mutant.emivespa.com/stats`
- `http://mutant.emivespa.com/mutant`

<!-- Heard at "[¿Qué tal es trabajar en MERCADO LIBRE?](https://www.youtube.com/watch?v=6DpwMKNqoPk)" -->

### Building and running

The Makefile contains convenience recipes for building and running the container with all the right flags,
so you can run:

- `make build` and then
- `make run`

There are also convenience recipes for hitting container endpoints:

- `make 200`
- `make 403`
- `make healthcheck`
- `make stats`

For now, won't work without a running DB.

## DB

- Using Planetscale because setting up RDS would be a distraction.
- Might be the shoddiest part of this, it's my first time using Prisma with
  Golang.
