package app

import "braces.dev/errtrace"

func StartServer() error {
	srv := newWsServer()
	if err := srv.Start("0.0.0.0:8080"); err != nil {
		return errtrace.Wrap(err)
	}

	return nil
}
