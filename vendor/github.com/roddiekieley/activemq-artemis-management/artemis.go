package artemis

import (
	"github.com/roddiekieley/activemq-artemis-management/jolokia"
	"strings"
)

type IArtemis interface {
	NewArtemis(_ip string, _jolokiaPort string, _name string, _userName string, _password string) *Artemis
	Uptime() (*jolokia.ReadData, error)
	CreateQueue(addressName string, queueName string) (*jolokia.ExecData, error)
	DeleteQueue(queueName string) (*jolokia.ExecData, error)
	ListBindingsForAddress(addressName string) (*jolokia.ExecData, error)
	DeleteAddress(addressName string) (*jolokia.ExecData, error)
}

type Artemis struct {
	ip          string
	jolokiaPort string
	name        string
	userName    string
	password    string
	jolokia     *jolokia.Jolokia
}

func NewArtemis(_ip string, _jolokiaPort string, _name string, _user string, _password string) *Artemis {

	artemis := Artemis{
		ip:          _ip,
		jolokiaPort: _jolokiaPort,
		name:        _name,
		jolokia:     jolokia.NewJolokia(_ip, _jolokiaPort, "/console/jolokia", _user, _password),
	}

	return &artemis
}

func (artemis *Artemis) Uptime() (*jolokia.ReadData, error) {

	uptimeURL := "org.apache.activemq.artemis:broker=\"" + artemis.name + "\"/Uptime"
	data, err := artemis.jolokia.Read(uptimeURL)

	return data, err
}

func (artemis *Artemis) CreateQueue(addressName string, queueName string, routingType string) (*jolokia.ExecData, error) {

	url := "org.apache.activemq.artemis:broker=\\\"" + artemis.name + "\\\""
	routingType = strings.ToUpper(routingType)
	parameters := `"` + addressName + `","` + queueName + `",` + `"` + routingType + `"`
	jsonStr := `{ "type":"EXEC","mbean":"` + url + `","operation":"createQueue(java.lang.String,java.lang.String,java.lang.String)","arguments":[` + parameters + `]` + ` }`
	data, err := artemis.jolokia.Exec(url, jsonStr)

	return data, err
}

func (artemis *Artemis) DeleteQueue(queueName string) (*jolokia.ExecData, error) {

	url := "org.apache.activemq.artemis:broker=\\\"" + artemis.name + "\\\""
	parameters := `"` + queueName + `"`
	jsonStr := `{ "type":"EXEC","mbean":"` + url + `","operation":"destroyQueue(java.lang.String)","arguments":[` + parameters + `]` + ` }`
	data, err := artemis.jolokia.Exec(url, jsonStr)

	return data, err
}

func (artemis *Artemis) ListBindingsForAddress(addressName string) (*jolokia.ExecData, error) {

	url := "org.apache.activemq.artemis:broker=\\\"" + artemis.name + "\\\""
	parameters := `"` + addressName + `"`
	jsonStr := `{ "type":"EXEC","mbean":"` + url + `","operation":"listBindingsForAddress(java.lang.String)","arguments":[` + parameters + `]` + ` }`
	data, err := artemis.jolokia.Exec(url, jsonStr)

	return data, err
}

func (artemis *Artemis) DeleteAddress(addressName string) (*jolokia.ExecData, error) {

	url := "org.apache.activemq.artemis:broker=\\\"" + artemis.name + "\\\""
	parameters := `"` + addressName + `"`
	jsonStr := `{ "type":"EXEC","mbean":"` + url + `","operation":"deleteAddress(java.lang.String)","arguments":[` + parameters + `]` + ` }`
	data, err := artemis.jolokia.Exec(url, jsonStr)

	return data, err
}
