package repository

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/srospan"
	"go.opentelemetry.io/otel/trace"
)

func tagCharacter(ctx context.Context, characterId string) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		srospan.TargetCharacterId(characterId),
	)
}
