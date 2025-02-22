// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package models

type Proxy struct {
	Id        int    `json:"id,omitempty"`
	Token     string `json:"token"`
	StartTime string `json:"start_time"`
	AdminAddr string `json:"admin_addr"`

	ProtoType string `json:"proto_type"`
	ProxyAddr string `json:"proxy_addr"`
	ProxyTLS  bool   `json:"proxy_tls"`
	ProxyTLSCert string `json:"proxy_tls_cert"`
	ProxyTLSKey string `json:"proxy_tls_key"`

	JodisPath string `json:"jodis_path,omitempty"`

	ProductName string `json:"product_name"`

	Pid int    `json:"pid"`
	Pwd string `json:"pwd"`
	Sys string `json:"sys"`

	MaxSlotNum int `toml:"max_slot_num" json:"max_slot_num"`

	Hostname   string `json:"hostname"`
	DataCenter string `json:"datacenter"`
}

func (p *Proxy) Encode() []byte {
	return jsonEncode(p)
}
