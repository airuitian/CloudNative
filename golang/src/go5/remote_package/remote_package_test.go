package remote_package

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestRemote(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
