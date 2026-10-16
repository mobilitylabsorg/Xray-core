//go:build !uniray_lowmem

package tun

// The most a TUN connection's TCP may buffer each way. uniray_lowmem (UniRay's
// builds for a small memory limit: the iOS packet tunnel, a router) bounds
// them lower, see stack_gvisor_buffers_lowmem.go.
const (
	tcpRXBufMaxSize = 8 << 20 // 8MiB
	tcpTXBufMaxSize = 6 << 20 // 6MiB
)
