package defaultplugins

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/ligato/cn-infra/logging/logrus"
	intf "github.com/ligato/vpp-agent/defaultplugins/ifplugin/model/interfaces"
)

const kafkaIfStateTopic = "if_state" // Kafka topic where interface state changes are published.

// Resync deletes obsolete operation status of network interfaces in DB
// Obsolete state is one that is not part of SwIfIndex
func (plugin *Plugin) resyncIfStateEvents(keys []string) error {
	for _, key := range keys {
		ifaceName, err := intf.ParseNameFromKey(key)
		if err != nil {
			return err
		}

		_, _, found := plugin.swIfIndexes.LookupIdx(ifaceName)
		if !found {
			log.Debug("deleting obsolete status begin ", key)
			err := plugin.Transport.PublishData(key, nil /*means delete*/)
			log.Debug("deleting obsolete status end ", key, err)
		} else {
			log.WithField("ifaceName", ifaceName).Debug("interface status is needed")
		}
	}

	return nil
}

// publishIfState goroutine is used to watch interface state notifications that are propagated to Kafka topic
func (plugin *Plugin) publishIfStateEvents(ctx context.Context) {
	plugin.wg.Add(1)
	defer plugin.wg.Done()

	for {
		select {
		case ifState := <-plugin.ifStateChan:
			plugin.Transport.PublishData(intf.InterfaceStateKey(ifState.State.Name), ifState.State)

			// marshall data into JSON & send kafka message
			if plugin.kafkaConn != nil && ifState.Type == intf.UPDOWN {
				json, err := json.Marshal(ifState.State)
				if err != nil {
					log.Error(err)
				} else {

					// send kafka message
					_, err = plugin.kafkaConn.SendSyncString(kafkaIfStateTopic,
						fmt.Sprintf("%s", ifState.State.Name), string(json))
					if err != nil {
						log.Error(err)
					} else {
						log.Debug("Sending Kafka notification")
					}
				}
			}

		case <-ctx.Done():
			// stop watching for state data updates
			return
		}
	}
}
