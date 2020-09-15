package rgnets

// Account ...
type Account struct {
	ID                 int                        `json:"id"`
	Username           string                     `json:"login"`
	Firstname          string                     `json:"first_name"`
	Lastname           string                     `json:"last_name"`
	Email              string                     `json:"email"`
	PSK                string                     `json:"pre_shared_key"`
	MaxDevices         int                        `json:"max_devices"`
	Group              AccountGroup               `json:"account_group"`
	LoginSessions      []AccountLoginSession      `json:"login_sessions"`
	VlanTagAssignments []AccountVlanTagAssignment `json:"vlan_tag_assignments"`
}

// AccountGroup ...
type AccountGroup struct {
	ID       int    `json:"id"`
	PolicyID int    `json:"policy_id"`
	Name     string `json:"name"`
}

// AccountLoginSession ...
type AccountLoginSession struct {
	ID       int    `json:"id"`
	Username string `json:"login"`
	IP       string `json:"ip"`
	MacAddr  string `json:"mac"`
	Online   bool   `json:"online"`
	Hostname string `json:"hostname"`
}

// AccountVlanTagAssignment ...
type AccountVlanTagAssignment struct {
	ID               int    `json:"id"`
	ClusterNodeID    int    `json:"cluster_node_id"`
	RadiusServerID   int    `json:"radius_server_id"`
	VlanID           int    `json:"vlan_id"`
	Tag              int    `json:"tag"`
	MacAddr          string `json:"mac"`
	NasIP            string `json:"nas_ip"` // Wireless AP ... What else?
	AccountSessionID string `json:"account_session_id"`
}

// GetAccountOptions options used to query for an Account
type GetAccountOptions struct {
	ID       int
	Username string
}

// Device ...
type Device struct {
	ID        int     `json:"id"`
	AccountID int     `json:"account_id"`
	Name      string  `json:"name"` // Defaults to MAC if not changed
	MacAddr   string  `json:"mac"`
	Account   Account `json:"account"`
}

// GetDeviceOptions options used to query for a user's Device(s)
type GetDeviceOptions struct {
	AccountID int
	DeviceID  int
	MacAddr   string
}

// DhcpLease ...
type DhcpLease struct {
	ID       int64  `json:"id"`
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	MacAddr  string `json:"mac"`
}

// DhcpLeaseOptions ...
type DhcpLeaseOptions struct {
	Hostname string
	IP       string
}
