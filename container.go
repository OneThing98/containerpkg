package container

type Container struct {
	ID           string       `json:"id,omitempty"`
	NsPid        int          `json:"namespace_pid,omitempty"`
	Command      *Command     `json:"command,omitempty"`
	RootFs       string       `json:"rootfs,omitempty"`
	ReadonlyFs   bool         `json:"readonly_fs,omitempty"`
	NetNsFd      uintptr      `json:"network_namespace_fd,omitempty"`
	User         string       `json:"user,omitempty"`
	WorkingDir   string       `json:"working_dir,omitempty"`
	Namespaces   Namespaces   `json:"namespaces,omitempty"`
	Capabilities Capabilities `json:"capabilities,omitempty"`
}

type Command struct {
	Args []string `json:"args,omitempty"`
	Env  []string `json:"environment,omitempty"`
}

type Network struct {
	TempVethName string `json:"temp_veth,omitempty"`
	IP           string `json:"ip,omitempty"`
	Gateway      string `json:"gateway,omitempty"`
	Bridge       string `json:"bridge,omitempty"`
	Mtu          int    `json:"mtu,omitempty"`
}

type Namespace string
type Namespaces []Namespace

func (n Namespaces) Contains(ns Namespace) bool {
	for _, nns := range n {
		if nns == ns {
			return true
		}
	}
	return false
}

type Capability string
type Capabilities []Capability

func (c Capabilities) Contains(capp Capability) bool {
	for _, cc := range c {
		if cc == capp {
			return true
		}
	}
	return false
}

const (
	CAP_SETPCAP        Capability = "SETPCAP"
	CAP_SYS_MODULE     Capability = "SYS_MODULE"
	CAP_SYS_RAWIO      Capability = "SYS_RAWIO"
	CAP_SYS_PACCT      Capability = "SYS_PACCT"
	CAP_SYS_ADMIN      Capability = "SYS_ADMIN"
	CAP_SYS_NICE       Capability = "SYS_NICE"
	CAP_SYS_RESOURCE   Capability = "SYS_RESOURCE"
	CAP_SYS_TIME       Capability = "SYS_TIME"
	CAP_SYS_TTY_CONFIG Capability = "SYS_TTY_CONFIG"
	CAP_MKNOD          Capability = "MKNOD"
	CAP_AUDIT_WRITE    Capability = "AUDIT_WRITE"
	CAP_AUDIT_CONTROL  Capability = "AUDIT_CONTROL"
	CAP_MAC_OVERRIDE   Capability = "MAC_OVERRIDE"
	CAP_MAC_ADMIN      Capability = "MAC_ADMIN"

	CLONE_NEWNS   Namespace = "NEWNS"   // mount
	CLONE_NEWUTS  Namespace = "NEWUTS"  // utsname
	CLONE_NEWIPC  Namespace = "NEWIPC"  // ipc
	CLONE_NEWUSER Namespace = "NEWUSER" // user
	CLONE_NEWPID  Namespace = "NEWPID"  // pid
	CLONE_NEWNET  Namespace = "NEWNET"  // network
)
