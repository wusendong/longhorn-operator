func volumeequal(old, cur *lv1.Volume) bool {
	if !baseequal(old.Spec.Volume, old.Spec.Volume) {
		return false
	}
	if old.Status == nil && cur.Status == nil {
		return true
	}
	if old.Status != nil && cur.Status != nil {
		if !equalController(old.Status.Controller, cur.Status.Controller) {
			return false
		}
		if len(old.Status.Replicas) != len(cur.Status.Replicas) {
			return false
		}
		if len(old.Status.Replicas) == 0 && len(cur.Status.Replicas) == 0 {
			return true
		}

		sort.Sort(Replicas(old.Status.Replicas))
		sort.Sort(Replicas(cur.Status.Replicas))
		for i := 0; i < len(old.Status.Replicas); i++ {
			if !equalReplica(old.Status.Replicas[i], cur.Status.Replicas[i]) {
				return false
			}
		}
		return true

	}
	return false

}

type Replicas []*types.ReplicaInfo

// Len is the number of elements in the collection.
func (r Replicas) Len() int {
	return len(r)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (r Replicas) Less(i, j int) bool {
	return r[i].Name < r[j].Name
}

// Swap swaps the elements with indexes i and j.
func (r Replicas) Swap(i, j int) {
	tmp := r[i]
	r[j] = r[i]
	r[i] = tmp
}

func baseequal(old, cur *types.VolumeInfo) bool {
	if old == nil && cur == nil {
		return true
	}
	if old != nil && cur != nil {
		if old.Name != cur.Name {
			return false
		}
		if old.Size != cur.Size {
			return false
		}
		if old.BaseImage != cur.BaseImage {
			return false
		}
		if old.FromBackup != cur.FromBackup {
			return false
		}
		if old.NumberOfReplicas != cur.NumberOfReplicas {
			return false
		}
		if old.StaleReplicaTimeout != cur.StaleReplicaTimeout {
			return false
		}
		if old.TargetNodeID != cur.TargetNodeID {
			return false
		}
		if old.DesireState != cur.DesireState {
			return false
		}
		if old.Created != cur.Created {
			return false
		}
		if old.NodeID != cur.NodeID {
			return false
		}
		if old.State != cur.State {
			return false
		}
		if old.Endpoint != cur.Endpoint {
			return false
		}
		return true
	}
	return false
}

func equalReplica(old, cur *types.ReplicaInfo) bool {
	if old == nil && cur == nil {
		return true
	}
	if old != nil && cur != nil {
		if old.ID != cur.ID {
			return false
		}
		if old.Type != cur.Type {
			return false
		}
		if old.Name != cur.Name {
			return false
		}
		if old.NodeID != cur.NodeID {
			return false
		}
		if old.IP != cur.IP {
			return false
		}
		if old.Running != cur.Running {
			return false
		}
		if old.VolumeName != cur.VolumeName {
			return false
		}
		return true
	}
	return false
}

func equalController(old, cur *types.ControllerInfo) bool {
	if old == nil && cur == nil {
		return true
	}

	if old != nil && cur != nil {
		if old.ID != cur.ID {
			return false
		}
		if old.Type != cur.Type {
			return false
		}
		if old.Name != cur.Name {
			return false
		}
		if old.NodeID != cur.NodeID {
			return false
		}
		if old.IP != cur.IP {
			return false
		}
		if old.Running != cur.Running {
			return false
		}
		if old.VolumeName != cur.VolumeName {
			return false
		}
		return true
	}
	return false
}