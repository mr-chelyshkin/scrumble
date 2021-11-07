package version

import "go.uber.org/zap"

func LogFields() []zap.Field {
	v := Get()
	return []zap.Field{
		zap.String("git.version", v.GitVersion),
		zap.String("git.commit_hash", v.GitCommit),
		zap.String("go.version", v.GoVersion),
		zap.String("build.date", v.BuildDate),
		zap.Bool("build.race", v.Race),
		zap.String("platform", v.Platform),
	}
}
