package policy

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Firefox does not allow you to supply a path to an enterprise policy at launch
// So we must delete the existing policy and create a new file at each launch
var macPoliciesPath = `Contents/Resources/distribution`
var policiesFilename = `policies.json`

var DefaultPolicy = PolicyRoot{
	Policies: Policies{
		DontCheckDefaultBrowser: &[]bool{true}[0],
		AppAutoUpdate:           &[]bool{false}[0],
		ManualAppUpdateOnly:     &[]bool{false}[0],
	},
}

func ClearPolicies(appPath string) error {
	err := os.Remove(filepath.Join(appPath, macPoliciesPath, policiesFilename))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return nil
}

// TODO: Unfortunately macos thinks we're being malicious and blocks this operation unless the user allows it.
// What are the alternatives? Construct a profile programmatically and use CLI args probably.
func (p *PolicyRoot) Save(appPath string) error {
	buf, err := json.Marshal(p)
	if err != nil {
		return err
	}
	py := filepath.Join(appPath, macPoliciesPath)
	err = os.MkdirAll(py, 0755)
	if err != nil {
		return err
	}
	err = ClearPolicies(appPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(py, policiesFilename), buf, 0644)
	if err != nil {
		return err
	}
	return nil
}

type PolicyRoot struct {
	Policies Policies `json:"policies"`
}

type Policies struct {
	AppAutoUpdate           *bool               `json:"AppAutoUpdate,omitempty"`
	DontCheckDefaultBrowser *bool               `json:"DontCheckDefaultBrowser,omitempty"`
	ManualAppUpdateOnly     *bool               `json:"ManualAppUpdateOnly,omitempty"`
	Certificates            *CertificatesPolicy `json:"Certificates,omitempty"`
	Proxy                   *ProxyPolicy        `json:"Proxy,omitempty"`
}

type CertificatesPolicy struct {
	Install []string `json:"Install,omitempty"`
}

type ProxyPolicy struct {
	Mode                        string `json:"Mode,omitempty"`
	Locked                      *bool  `json:"Locked,omitempty"`
	UseHTTPProxyForAllProtocols *bool  `json:"UseHTTPProxyForAllProtocols,omitempty"`
	SSLProxy                    string `json:"SSLProxy,omitempty"`
}
