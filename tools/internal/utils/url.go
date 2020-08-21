package utils

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetRTMPURLFlag(cmd *cobra.Command, flagname string) (url.URL, error) {
	urlstr, err := cmd.Flags().GetString(flagname)
	if err != nil {
		return url.URL{}, errors.Wrap(err, "could not read flag")
	}
	u, err := url.Parse(urlstr)
	if err != nil {
		return url.URL{}, errors.Wrap(err, "could not parse url")
	}
	if u.Scheme != "rtmp" && u.Scheme != "rtmps" {
		return url.URL{}, errors.New("invalid url scheme. it must be rtmp:// or rmtps://")
	}
	return *u, nil
}
