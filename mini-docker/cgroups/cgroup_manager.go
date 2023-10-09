package cgroups

import (
	"basicGo/mini-docker/cgroups/subsystems"

	log "github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	res := &CgroupManager{
		Path: path,
	}
	log.Info("create NewCgroupManager success, ", *res)
	return res
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		subSysIns.Apply(c.Path, pid)
		log.Infof("Apply is executed, %s, %s", c.Path, pid)
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		subSysIns.Set(c.Path, res)
		log.Infof("Set is executed, %s, %s", c.Path, res)
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("fail to remove cgroup %v", err)
		}
	}
	return nil
}
