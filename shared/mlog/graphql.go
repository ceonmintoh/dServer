// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package mlog

import (
	"context"
)

// GraphQLLogger is used to log panics that occur during query execution.
type GraphQLLogger struct{}

// LogPanic satisfies the graphql/log.Logger interface.
// It just panics, which moves it up the layer and gets captured
// in the runServer function in cmd/mattermost/commands/server.go
// and gets logged appropriately.
func (l *GraphQLLogger) LogPanic(_ context.Context, value interface{}) {
	panic(value)
}
