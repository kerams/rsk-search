package config

import (
	"github.com/spf13/pflag"
	"github.com/warmans/rsk-search/pkg/flag"
)

type SearchServiceConfig struct {
	BleveIndexPath  string
	BlugeIndexPath  string
	Scheme          string
	Hostname        string
	RewardsDisabled bool
	FilesBasePath   string
	AudioUriPattern string
}

func (c *SearchServiceConfig) RegisterFlags(fs *pflag.FlagSet, prefix string) {
	flag.StringVarEnv(fs, &c.Scheme, prefix, "scheme", "http://", "scheme to use for absolute links")
	flag.StringVarEnv(fs, &c.Hostname, prefix, "hostname", "localhost", "hostname to use for absolute links")
	flag.StringVarEnv(fs, &c.BleveIndexPath, prefix, "bleve-index-path", "./var/rsk.bleve", "location of bleve search index")
	flag.StringVarEnv(fs, &c.BlugeIndexPath, prefix, "bluge-index-path", "./var/rsk.bluge", "location of bluge search index")
	flag.StringVarEnv(fs, &c.FilesBasePath, prefix, "files-base-path", "./var", "location files that can be downloaded")
	flag.BoolVarEnv(fs, &c.RewardsDisabled, prefix, "rewards-disabled", false, "Disable claiming rewards (but sill calculate them)")
	flag.StringVarEnv(fs, &c.AudioUriPattern, prefix, "audio-uri-pattern", "https://storage.googleapis.com/scrimpton-raw-audio/%s.mp3", "episode ID e.g. xfm-S1E01 will be interpolated into this string")
}
