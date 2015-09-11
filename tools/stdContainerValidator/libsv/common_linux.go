/*generated from:
linux/capability.h
asm-generic/resource.h
*/
package specsValidator

func capValid(capability string) bool {
	caps := map[string]int{
		"CAP_CHOWN":            0,
		"CAP_DAC_OVERRIDE":     1,
		"CAP_DAC_READ_SEARCH":  2,
		"CAP_FOWNER":           3,
		"CAP_FSETID":           4,
		"CAP_KILL":             5,
		"CAP_SETGID":           6,
		"CAP_SETUID":           7,
		"CAP_SETPCAP":          8,
		"CAP_LINUX_IMMUTABLE":  9,
		"CAP_NET_BIND_SERVICE": 10,
		"CAP_NET_BROADCAST":    11,
		"CAP_NET_ADMIN":        12,
		"CAP_NET_RAW":          13,
		"CAP_IPC_LOCK":         14,
		"CAP_IPC_OWNER":        15,
		"CAP_SYS_MODULE":       16,
		"CAP_SYS_RAWIO":        17,
		"CAP_SYS_CHROOT":       18,
		"CAP_SYS_PTRACE":       19,
		"CAP_SYS_PACCT":        20,
		"CAP_SYS_ADMIN":        21,
		"CAP_SYS_BOOT":         22,
		"CAP_SYS_NICE":         23,
		"CAP_SYS_RESOURCE":     24,
		"CAP_SYS_TIME":         25,
		"CAP_SYS_TTY_CONFIG":   26,
		"CAP_MKNOD":            27,
		"CAP_LEASE":            28,
		"CAP_AUDIT_WRITE":      29,
		"CAP_AUDIT_CONTROL":    30,
		"CAP_SETFCAP	": 31,
		"CAP_MAC_OVERRIDE":  32,
		"CAP_MAC_ADMIN":     33,
		"CAP_SYSLOG":        34,
		"CAP_WAKE_ALARM":    35,
		"CAP_BLOCK_SUSPEND": 36,
	}
	if _, ok := caps[capability]; ok {
		return true
	}
	return false
}

func rlimitValid(rlimit string) bool {
	rlimits := map[string]int{
		"RLIMIT_CPU":        0,
		"RLIMIT_FSIZE":      1,
		"RLIMIT_DATA":       2,
		"RLIMIT_STACK":      3,
		"RLIMIT_CORE":       4,
		"RLIMIT_RSS":        5,
		"RLIMIT_NPROC":      6,
		"RLIMIT_NOFILE":     7,
		"RLIMIT_MEMLOCK":    8,
		"RLIMIT_AS":         9,
		"RLIMIT_LOCKS":      10,
		"RLIMIT_SIGPENDING": 11,
		"RLIMIT_MSGQUEUE":   12,
		"RLIMIT_NICE":       13,
		"RLIMIT_RTPRIO":     14,
		"RLIMIT_RTTIME":     15,
	}
	if _, ok := rlimits[rlimit]; ok {
		return true
	}
	return false
}

func requiredPaths() []string {
	paths := []string{
		"/proc", //procfs
		"/sys",  //sysfs
	}
	return paths
}

func requiredDevices() []string {
	devices := []string{
		"/dev/null",    //device
		"/dev/zero",    //device
		"/dev/full",    //device
		"/dev/random",  //device
		"/dev/urandom", //device
		"/dev/tty",     //device
		"/dev/console", //device
		"/dev/pts",     //devpts
		"/dev/ptmx",    //device	Bind-mount or symlink of /dev/pts/ptmx
		"/dev/shm",     //tmpfs
	}
	return devices
}