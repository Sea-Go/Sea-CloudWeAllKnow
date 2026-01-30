package schema

import "github.com/milvus-io/milvus/client/v2/entity"

func RecllCandidateTableName() *entity.Schema {
	chunkId := entity.NewField().
		WithName("id").
		WithDataType(entity.FieldTypeString).
		WithIsPrimaryKey(true).
		WithIsAutoID(false)

	ar

	vec := entity.NewField().
		WithName("vector").
		WithDataType(entity.FieldTypeFloatVector).
		WithDim(2048)

	tag := entity.NewField().
		WithName("tag").
		WithDataType(entity.FieldTypeVarChar).
		WithTypeParams(entity.TypeParamMaxLength, "256")

	return entity.NewSchema().
		WithName("RecallCandidateCollection").
		WithDescription("coarse recall vectors").
		WithAutoID(false).
		WithDynamicFieldEnabled(true).
		WithField(chunkId).
		WithField(vec).
		WithField(tag)
}
