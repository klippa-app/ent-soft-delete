package softdelete

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Mixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Optional(),
	}
}
