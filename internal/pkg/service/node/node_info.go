package node

import (
	"github.com/gavril-s/attic-master/api"
	"github.com/gavril-s/attic-node/internal/config"
)

func GetNodeInfo(cfg config.Config) (api.NodeInfo, error) {
	nodeInfo := api.NodeInfo{
		IsPersistent:      cfg.IsPersistent,
		AcceptedChunkSize: cfg.AcceptedChunkSize.Bytes(),
		StorageInfo:       make([]api.StorageInfo, len(cfg.Storages)),
	}

	for storageIndex, storage := range cfg.Storages {
		freeCapacity, err := storage.FreeCapacity()
		if err != nil {
			return nodeInfo, err
		}
		nodeInfo.StorageInfo[storageIndex] = api.StorageInfo{
			StorageIndex:  uint64(storageIndex),
			TotalCapacity: storage.Capacity.Bytes(),
			FreeCapacity:  freeCapacity.Bytes(),
		}
	}

	return nodeInfo, nil
}
