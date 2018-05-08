// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package repo_indexer

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
)

func NewRepoIndexerClient(ctx context.Context) (pb.RepoIndexerClient, error) {
	conn, err := manager.NewClient(ctx, constants.RepoIndexerHost, constants.RepoIndexerPort)
	if err != nil {
		return nil, err
	}
	return pb.NewRepoIndexerClient(conn), err
}
