# MELI backend challenge / mutantes

[Challenge found here](https://github.com/mauricionrgarcia/examen-mercadolibre-mutante)
(see README).

## Part 1 - isMutant method

- [ ] is there a non log(n) way to do it?

### Current assumptions

- DNA strands are 6 characters long
- a match is 4 identical characters in a row
- we're being passed 6 strands exactly
- we're looking for 2 matches

### Efficiency notes

- go is plenty fast
- we return early whenever possible
  - tuple check returns early as soon as it knows it's not a match
  - passes return early once we've reached the mutant threshhold
  - isMutant returns early whenever a pass returns 0 matches left
  - **doesn't actually matter all that much for such a small input,
    the worst case still tests at 0.00s**

### Another possible approach

- generate list of all possible 4-tuples
  - transpose the array to get the vertical ones(?)
  - not sure whether there's an elegant way to do the diagonals,
    but it's certainly possible
- pass them all to the same function
- why not
  - would technically be slower - adds a pass to do pretty much the same
  - program logic would not be all that different

## Part 2 - rest API

Currently hosted at `http://mutant.emivespa.com/stats`

Wasn't sure whether to go with a container or a lambda. Went with a container.

<!-- Heard at "[¿Qué tal es trabajar en MERCADO LIBRE?](https://www.youtube.com/watch?v=6DpwMKNqoPk)" -->

### Building and running

The Makefile contains convenience recipes for building and running the container with all the right flags,
so you can run
`make build`, and then
`make run`.
You can also run
`make healthcheck`,
`make 200` and
`make 403`
to test the running container.

## Part 3 - database

- [ ] part 3 - database
