package main

import "flag"

var ToSeed = false

func init() {
	flag.BoolVar(&ToSeed, "seed", false, "To create initial database tables as defined in the seed.sql file")
}
