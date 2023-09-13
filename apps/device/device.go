package device

func NewDeviceSet() *DeviceSet {
	return &DeviceSet{
		Items: []*Device{},
	}
}

func (s *DeviceSet) Add(item *Device) {
	s.Items = append(s.Items, item)
}

func NewDeviceByNamespaceSet() *DeviceByNamespaceSet {
	return &DeviceByNamespaceSet{
		Items: []string{},
	}
}

func (s *DeviceByNamespaceSet) Add(item string) {
	s.Items = append(s.Items, item)
}
