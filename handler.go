package blivedm

type HandlerFunc func(*Context)

type DisconnectCallback func(*BLiveWsClient)
