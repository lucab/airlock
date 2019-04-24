package cli

import (
	"fmt"
	"net/http"

	"github.com/coreos/airlock/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdServe = &cobra.Command{
		Use:  "serve",
		RunE: runServe,
	}
)

func runServe(cmd *cobra.Command, cmdArgs []string) error {
	logrus.WithFields(logrus.Fields{
		"groups": runSettings.LockGroups,
	}).Debug("lock groups")

	logrus.WithFields(logrus.Fields{
		"address": runSettings.ServiceAddress,
		"port":    runSettings.ServicePort,
	}).Info("starting service")

	airlock := server.Airlock{*runSettings}
	http.Handle(server.PreRebootEndpoint, airlock.PreReboot())
	http.Handle(server.SteadyStateEndpoint, airlock.SteadyState())

	listenAddr := fmt.Sprintf("%s:%d", runSettings.ServiceAddress, runSettings.ServicePort)
	return http.ListenAndServe(listenAddr, nil)

	return nil
}
