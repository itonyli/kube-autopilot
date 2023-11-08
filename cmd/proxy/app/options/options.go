package options

import (
	flag "autopilot/pkg"
	"autopilot/pkg/etcd"
)

type ServerRunOptions struct {
	Etcd        *etcd.EtcdOptions
	MasterCount int
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		MasterCount: 3,
	}

	return &s
}

// Flags returns flags for a specific APIServer by section name
func (s *ServerRunOptions) Flags() (fss flag.NamedFlagSets) {
	fs := fss.FlagSet("misc")
	fs.IntVar(&s.MasterCount, "apiserver-count", s.MasterCount,
		"The number of apiservers running in the cluster, must be a positive number. (In use when --endpoint-reconciler-type=master-count is enabled.)")
	return fss
}
