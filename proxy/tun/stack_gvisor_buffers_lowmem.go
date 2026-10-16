//go:build uniray_lowmem

package tun

// uniray_lowmem: 1 MiB each way, gVisor's default, instead of 8 and 6 MiB. The
// TUN's peer is an app on the same device (or a LAN client), a short path that
// needs no more, and many fast connections stay within a small memory limit
// (iOS gives a packet tunnel about 50 MB).
const (
	tcpRXBufMaxSize = 1 << 20
	tcpTXBufMaxSize = 1 << 20
)
