package resolver

import (
	"context"

	"github.com/1t2t3t4t/my_journal_api/service"
)

func guardLoggedInUser(ctx context.Context) (service.AuthClaim, bool) {
	claim, ok := ctx.Value(service.UserClaim).(service.AuthClaim)
	return claim, ok
}
