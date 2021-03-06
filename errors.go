// Package bloomfilter is face-meltingly fast, thread-safe,
// marshalable, unionable, probability- and
// optimal-size-calculating Bloom filter in go
//
// https://github.com/steakknife/bloomfilter
//
// Copyright © 2014, 2015, 2018 Barry Allard
//
// MIT license
//
package bloomfilter

import "fmt"

var errHash = fmt.Errorf(
	"Hash mismatch, the Bloom filter is probably corrupt")
var errK = fmt.Errorf(
	"keys must have length %d or greater", KMin)
var errM = fmt.Errorf(
	"m (number of bits in the Bloom filter) must be >= %d", MMin)
var errUniqueKeys = fmt.Errorf(
	"Bloom filter keys must be unique")
var errIncompatibleBloomFilters = fmt.Errorf(
	"Cannot perform union on two incompatible Bloom filters")
