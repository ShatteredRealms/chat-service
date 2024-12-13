package repository

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/srospan"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

func tagCharacter(ctx context.Context, characterId *uuid.UUID) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		srospan.TargetCharacterId(characterId.String()),
	)
}
