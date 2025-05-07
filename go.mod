module github.com/SamW94/GoGator

go 1.23.5

replace github.com/SamW94/GoGator/internal/config => ./internal/config
replace github.com/SamW94/GoGator/internal/database => ./internal/database
replace github.com/SamW94/GoGator/internal/rss => ./internal/rss

require github.com/SamW94/GoGator/internal/config v0.0.0
require github.com/SamW94/GoGator/internal/database v0.0.0
require github.com/SamW94/GoGator/internal/rss v0.0.0

require github.com/lib/pq v1.10.9
require "github.com/google/uuid" v1.6.0
