package rgnets

// Account ...
type Account struct {
	ID         int    `json:"id"`
	Username   string `json:"login"`
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Email      string `json:"email"`
	PSK        string `json:"pre_shared_key"`
	MaxDevices int    `json:"max_devices"`
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
