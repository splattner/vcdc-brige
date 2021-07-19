package vdcdapi

import "log"

func (e *Device) NewDevice(client Client, uniqueID string) {
	e.UniqueID = uniqueID
	e.Tag = uniqueID
	e.Output = BasicOutput

	e.client = client
	e.ModelName = e.client.modelName
	e.VendorName = e.client.vendorName
}

func (e *Device) NewLightDevice(client Client, uniqueID string, dimmable bool) {
	e.NewDevice(client, uniqueID)

	if dimmable {
		e.Output = LightOutput
	}

	e.Group = YellowLightGroup
	e.ColorClass = YellowColorClassT
}

func (e *Device) SetName(name string) {
	e.Name = name
}

func (e *Device) SetTag(tag string) {
	e.Tag = tag
}

func (e *Device) AddButton(button Button) {

	e.Buttons = append(e.Buttons, button)

}

func (e *Device) AddSensor(sensor Sensor) {

	e.Sensors = append(e.Sensors, sensor)

}

func (e *Device) AddInput(input Input) {

	e.Inputs = append(e.Inputs, input)

}

func (e *Device) UpdateValue(newValue float32) {

	// only update when changed
	if newValue != e.value {
		e.value = newValue
		e.client.UpdateValue(*e)
	}

}

func (e *Device) SetInitDone() {
	e.InitDone = true
	log.Printf("Init for Device %s done: %t\n", e.UniqueID, e.InitDone)
}

func (e *Device) SetChannelMessageCB(cb func(message *GenericVDCDMessage, device *Device)) {
	e.channel_cb = cb
}
