package gobotNeurosky

import (
	"github.com/hybridgroup/gobot"
)

type NeuroskyAdaptor struct {
	gobot.Adaptor
	sp gobot.Port
}

func (me *NeuroskyAdaptor) Connect() bool {
	me.sp = gobot.ConnectToSerial(me.Adaptor.Port, 57600)
	me.Connected = true
	return true
}

func (me *NeuroskyAdaptor) Reconnect() bool {
	if me.Connected == true {
		me.Disconnect()
	}
	return me.Connect()
}

func (me *NeuroskyAdaptor) Disconnect() bool {
	me.sp.Close()
	me.Connected = false
	return true
}

func (me *NeuroskyAdaptor) Finalize() bool {
	return true
}
