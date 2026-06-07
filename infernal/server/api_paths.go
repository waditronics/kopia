package server

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/kopia/kopia/infernal/ospath"
	"github.com/kopia/kopia/infernal/serverapi"
	"github.com/kopia/kopia/snapshot"
)

func handlePathResolve(_ context.Context, rc requestContext) (any, *apiError) {
	var req serverapi.ResolvePathRequest

	if err := json.Unmarshal(rc.body, &req); err != nil {
		return nil, requestError(serverapi.ErrorMalformedRequest, "malformed request body")
	}

	return &serverapi.ResolvePathResponse{
		SourceInfo: snapshot.SourceInfo{
			Path:     filepath.Clean(ospath.ResolveUserFriendlyPath(req.Path, true)),
			Host:     rc.rep.ClientOptions().Hostname,
			UserName: rc.rep.ClientOptions().Username,
		},
	}, nil
}
