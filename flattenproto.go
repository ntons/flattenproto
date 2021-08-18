package flattenproto

import (
	"google.golang.org/protobuf/proto"
)

func Flatten(m proto.Message, opts ...Option) Values {
	o := &options{
		FieldJoinCharacter: '.',
	}
	for _, opt := range opts {
		opt.apply(o)
	}
	return (&parser{options: o}).flatten(m)
}

type Field struct {
	Key   string
	Value interface{}
}
type Values []Field

type Option interface {
	apply(o *options)
}
type options struct {
	// Default: '.'
	FieldJoinCharacter byte
	// UseProtoNames uses proto field name instead of lowerCamelCase name in JSON
	// field names.
	UseProtoNames bool
	// UseEnumNumbers emits enum values as numbers.
	UseEnumNumbers bool
}
type funcOption struct {
	f func(*options)
}

func (fo funcOption) apply(o *options) { fo.f(o) }

func WithFieldJoinCharacter(b byte) Option {
	return funcOption{func(o *options) { o.FieldJoinCharacter = b }}
}
func UseProtoNames() Option {
	return funcOption{func(o *options) { o.UseProtoNames = true }}
}
func UseEnumNumbers() Option {
	return funcOption{func(o *options) { o.UseEnumNumbers = true }}
}
