module github.com/SamW94/blogo-aggregator

go 1.23.5

replace github.com/SamW94/blogo-aggregator/internal/config => ./internal/config
replace github.com/SamW94/blogo-aggregator/internal/database => ./internal/database

require github.com/SamW94/blogo-aggregator/internal/config v0.0.0
require github.com/SamW94/blogo-aggregator/internal/database v0.0.0

require github.com/lib/pq v1.10.9
require "github.com/google/uuid" v1.6.0
