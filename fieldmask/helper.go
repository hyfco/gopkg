package fieldmask

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Filter clears every field that not visibility.
func Filter(pb proto.Message, vf Visibility) {
	m := pb.ProtoReflect()
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		opts := fd.Options().(*descriptorpb.FieldOptions)
		vd := proto.GetExtension(opts, E_Fv).(Visibility)
		if vd&vf == 0 {
			m.Clear(fd)
			return true
		}

		if fd.IsList() {
			list := v.List()
			for i := 0; i < list.Len(); i++ {
				e := list.Get(i)
				Filter(e.Message().Interface(), vf)
			}
			return true
		}

		if fd.Kind() == protoreflect.MessageKind {
			Filter(v.Message().Interface(), vf)
		}

		return true
	})
}
