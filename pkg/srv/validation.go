package srv

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	"github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/WilSimpson/gocloak/v13"
)

func (s *chatServiceServer) validateRole(ctx context.Context, role *gocloak.Role) (*auth.SROClaims, error) {
	claims, ok := auth.RetrieveClaims(ctx)
	if !ok {
		return nil, srv.ErrPermissionDenied
	}
	if !claims.HasResourceRole(role, s.Context.Config.Keycloak.ClientId) {
		return nil, srv.ErrPermissionDenied
	}
	return claims, nil

}
