package schema

import "github.com/milvus-io/milvus/client/v2/entity"

func RecallPreciseTableName() *entity.Schema {
	id := entity.NewField().
		WithName("id").
		WithDataType(entity.FieldTypeString).
		WithIsPrimaryKey(true).
		WithIsAutoID(false)

	vec := entity.NewField().
		WithName("vector").
		WithDataType(entity.FieldTypeFloatVector).
		WithDim(2048)

	tag := entity.NewField().
		WithName("tag").
		WithDataType(entity.FieldTypeVarChar).
		WithTypeParams(entity.TypeParamMaxLength, "256")

	return entity.NewSchema().
		WithName("RecallPreciseCollection").
		WithDescription("coarse recall vectors").
		WithAutoID(false).
		WithDynamicFieldEnabled(true).
		WithField(id).
		WithField(vec).
		WithField(tag)
}
