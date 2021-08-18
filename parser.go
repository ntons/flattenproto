package flattenproto

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type parser struct {
	*options
	result Values
}

func (ps *parser) flatten(m proto.Message) Values {
	ps.walkThroughMessage("", m.ProtoReflect())
	return ps.result
}

func (ps *parser) add(k string, v interface{}) {
	ps.result = append(ps.result, Field{k, v})
}

func (ps *parser) join(a ...interface{}) string {
	sb := strings.Builder{}
	for _, v := range a {
		if sb.Len() > 0 {
			sb.WriteByte(ps.FieldJoinCharacter)
		}
		switch v := v.(type) {
		case string:
			sb.WriteString(v)
		case protoreflect.FieldDescriptor:
			if ps.UseProtoNames {
				sb.WriteString(string(v.Name()))
			} else {
				sb.WriteString(v.JSONName())
			}
		default:
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}
	return sb.String()
}

func (ps *parser) walkThroughMessage(prefix string, msg protoreflect.Message) {
	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.IsMap() {
			ps.walkThroughMap(prefix, fd, v)
		} else if fd.IsList() {
			ps.walkThroughList(prefix, fd, v)
		} else {
			ps.parseField(ps.join(prefix, fd), fd, v)
		}
		return true
	})
}

func (ps *parser) walkThroughMap(prefix string, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	v.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		ps.parseField(ps.join(prefix, fd, k), fd.MapValue(), v)
		return true
	})
	return
}

func (ps *parser) walkThroughList(prefix string, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	for i := 0; i < v.List().Len(); i++ {
		ps.parseField(ps.join(prefix, fd, i), fd, v.List().Get(i))
	}
}

func (ps *parser) parseField(fullPath string, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	fmt.Println("parsing: ", fullPath)
	switch fd.Kind() {
	case protoreflect.EnumKind:
		if ps.UseEnumNumbers {
			ps.add(fullPath, v.Enum())
		} else {
			ps.add(fullPath, fd.Enum().Values().ByNumber(v.Enum()).Name())
		}
	case protoreflect.MessageKind:
		ps.walkThroughMessage(fullPath, v.Message())
	case protoreflect.GroupKind:
		// Note that this feature is deprecated
		// and should not be used when creating new message types
	default:
		ps.add(fullPath, v.Interface())
	}
	return
}
