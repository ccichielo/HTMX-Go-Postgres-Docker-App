module example.com/main

go 1.21.3

replace example.com/lib => ./lib

require example.com/lib v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.9 // indirect
