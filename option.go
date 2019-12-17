package jq

type Option int

const OptionVarIndexAt = iota

type Options map[Option]interface{}
