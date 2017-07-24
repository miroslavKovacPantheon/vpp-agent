package cacheutil

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ligato/cn-infra/db"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/vpp-agent/idxvpp"
)

// CacheHelper is a helper which implementation is reused among multiple typesafe Caches.
// Beware: index stored in cached mapping is not valid. The meaningful values are the name and metadata.
type CacheHelper struct {
	IDX           idxvpp.NameToIdxRW
	Prefix        string
	DataPrototype proto.Message
	ParseName     func(key string) (name string, err error)
}

const placeHolderIndex uint32 = 0

// DoWatching is supposed to be used as a go routine. It select the data from channels in arguments.
func (helper *CacheHelper) DoWatching(resyncName string, watcher datasync.Watcher) {
	changeChan := make(chan datasync.ChangeEvent, 100)
	resyncChan := make(chan datasync.ResyncEvent, 100)

	watcher.WatchData(resyncName, changeChan, resyncChan, helper.Prefix)

	for {
		select {
		case resyncEv := <-resyncChan:
			err := helper.DoResync(resyncEv)
			resyncEv.Done(err)
		case dataChng := <-changeChan:
			err := helper.DoChange(dataChng)
			dataChng.Done(err)
		}
	}
}

// DoChange calls:
// - RegisterName in case of db.Put
// - UnregisterName in case of data.Del
func (helper *CacheHelper) DoChange(dataChng datasync.ChangeEvent) error {
	var err error
	switch dataChng.GetChangeType() {
	case db.Put:
		current := proto.Clone(helper.DataPrototype)
		dataChng.GetValue(current)
		name, err := helper.ParseName(dataChng.GetKey())
		if err == nil {
			helper.IDX.RegisterName(name, placeHolderIndex, current)
		}
	case db.Delete:
		name, err := helper.ParseName(dataChng.GetKey())
		if err == nil {
			helper.IDX.UnregisterName(name)
		}
	}
	return err
}

// DoResync list keys&values in ResyncEvent and then:
// - RegisterName (for names that are part of ResyncEvent)
// - UnregisterName (for names that are not part of ResyncEvent)
func (helper *CacheHelper) DoResync(resyncEv datasync.ResyncEvent) error {
	var wasError error
	//idx.RegisterName()
	ifaces, found := resyncEv.GetValues()[helper.Prefix]
	if found {
		// Step 1: fill the existing things
		resyncNames := map[string]interface{}{}
		for {
			item, stop := ifaces.GetNext()
			if stop {
				break
			}
			ifaceName, err := helper.ParseName(item.GetKey())
			if err != nil {
				wasError = err
			} else {
				current := proto.Clone(helper.DataPrototype)
				item.GetValue(current)
				helper.IDX.RegisterName(ifaceName, placeHolderIndex, current)
				resyncNames[ifaceName] = nil
			}
		}

		// Step 2:
		existingNames := []string{} //TODO
		for _, existingName := range existingNames {
			if _, found := resyncNames[existingName]; !found {
				helper.IDX.UnregisterName(existingName)
			}
		}
	}
	return wasError
}

func (helper *CacheHelper) String() string {
	return helper.Prefix
}
