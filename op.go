package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type SysInfo struct {
	OS       string
	Init     string
	Kernel   string
	Uptime   string
	CPU      string
	GPU      string
	Memory   string
	Disk     string
	WM       string
	Terminal string
	Packages string
}

func main() {
	info := getSysInfo()

	logo := []string{
		``,
		`         \/` + "`" + `-.,`,
		`          \/` + "`" + `-.,`,
		`           \/` + "`" + `-.,`,
		`    _     _ \/` + "`" + `-.,`,
		`   ('>   ('>     \/` + "`" + `-.`,
		`  /\"(   /\"(      \/` + "`" + `-.`,
		`  \_)` + "`" + `   \_)` + "`",
		`  mrf    mrf`,
		` unicorn is hungry`,
	}

	const (
		reset  = "\033[0m"
		bold   = "\033[1m"
		dim    = "\033[2m"
		blue   = "\033[34m"
		cyan   = "\033[36m"
		purple = "\033[35m"
		white  = "\033[37m"
		green  = "\033[32m"
		yellow = "\033[33m"
		red    = "\033[31m"
	)

	logoColors := []string{reset, red, yellow, green, cyan, blue, purple, red, yellow, green}

	lines := []struct {
		label string
		value string
		color string
	}{
		{"OS", info.OS, dim},           // ← значение теперь серое, как метка
		{"Kernel", info.Kernel, cyan},
		{"Uptime", info.Uptime, green},
		{"Packages", info.Packages, purple},
		{"Init", info.Init, yellow},
		{"WM", info.WM, cyan},
		{"CPU", info.CPU, blue},
		{"GPU", info.GPU, red},
		{"Memory", info.Memory, purple},
		{"Disk", info.Disk, yellow},
		{"Terminal", info.Terminal, cyan},
	}

	maxRows := len(logo)
	if len(lines) > maxRows {
		maxRows = len(lines)
	}

	logoWidth := 0
	for _, l := range logo {
		if w := len(l); w > logoWidth {
			logoWidth = w
		}
	}

	for i := 0; i < maxRows; i++ {
		logoLine := ""
		if i < len(logo) {
			logoLine = logo[i]
		}
		lc := reset
		if i < len(logoColors) {
			lc = logoColors[i]
		}
		pad := strings.Repeat(" ", logoWidth-len(logoLine))
		fmt.Printf("%s%s%s    ", lc, logoLine+pad, reset)

		if i < len(lines) {
			item := lines[i]
			label := fmt.Sprintf("%-9s", item.label)
			fmt.Printf("%s%s  %s%s%s%s\n", dim, label, reset, item.color, item.value, reset)
		} else {
			fmt.Println()
		}
	}
	fmt.Println()
}

func getSysInfo() SysInfo {
	info := SysInfo{
		OS:       runtime.GOOS,
		Init:     "Unknown",
		Kernel:   "Unknown",
		Uptime:   "Unknown",
		CPU:      "Unknown",
		GPU:      "Unknown",
		Memory:   "Unknown",
		Disk:     "Unknown",
		WM:       "Unknown",
		Terminal: "Unknown",
		Packages: "Unknown",
	}

	if f, err := os.Open("/etc/os-release"); err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				val := strings.SplitN(line, "=", 2)[1]
				val = strings.Trim(val, `"' `)
				info.OS = val
				break
			}
		}
		f.Close()
	}

	if _, err := os.Stat("/run/systemd/system"); err == nil {
		info.Init = "systemd"
	} else if _, err := os.Stat("/run/openrc"); err == nil {
		info.Init = "OpenRC"
	} else if pid1, err := os.Readlink("/proc/1/exe"); err == nil {
		parts := strings.Split(pid1, "/")
		info.Init = parts[len(parts)-1]
	}

	if f, err := os.Open("/proc/sys/kernel/osrelease"); err == nil {
		sc := bufio.NewScanner(f)
		if sc.Scan() {
			info.Kernel = strings.TrimSpace(sc.Text())
		}
		f.Close()
	}

	if f, err := os.Open("/proc/uptime"); err == nil {
		var sec float64
		fmt.Fscanf(f, "%f", &sec)
		f.Close()
		days := int(sec) / 86400
		hours := (int(sec) % 86400) / 3600
		mins := (int(sec) % 3600) / 60
		parts := []string{}
		if days > 0 {
			parts = append(parts, fmt.Sprintf("%dd", days))
		}
		if hours > 0 {
			parts = append(parts, fmt.Sprintf("%dh", hours))
		}
		parts = append(parts, fmt.Sprintf("%dm", mins))
		info.Uptime = strings.Join(parts, " ")
	}

	if f, err := os.Open("/proc/cpuinfo"); err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			if strings.HasPrefix(line, "model name") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					cpu := strings.TrimSpace(parts[1])
					cpu = strings.ReplaceAll(cpu, "AMD ", "")
					cpu = strings.ReplaceAll(cpu, "Intel(R) Core(TM) ", "")
					cpu = strings.ReplaceAll(cpu, " CPU", "")
					info.CPU = cpu
				}
				break
			}
		}
		f.Close()
	}

	info.GPU = getGPU()

	if f, err := os.Open("/proc/meminfo"); err == nil {
		sc := bufio.NewScanner(f)
		var total, avail float64
		for sc.Scan() {
			line := sc.Text()
			if strings.HasPrefix(line, "MemTotal:") {
				fmt.Sscanf(line, "MemTotal: %f", &total)
			} else if strings.HasPrefix(line, "MemAvailable:") {
				fmt.Sscanf(line, "MemAvailable: %f", &avail)
			}
		}
		f.Close()
		if total > 0 {
			used := total - avail
			info.Memory = fmt.Sprintf("%.1f / %.1f GiB", used/1024/1024, total/1024/1024)
		}
	}

	if out, err := exec.Command("df", "-h", "/").Output(); err == nil {
		lines := strings.Split(string(out), "\n")
		if len(lines) > 1 {
			fields := strings.Fields(lines[1])
			if len(fields) >= 5 {
				info.Disk = fmt.Sprintf("%s / %s (%s)", fields[2], fields[1], fields[4])
			}
		}
	}

	info.WM = detectWM()

	if term := os.Getenv("TERM_PROGRAM"); term != "" {
		info.Terminal = term
	} else if term := os.Getenv("TERM"); term != "" && term != "linux" {
		info.Terminal = term
	} else {
		info.Terminal = "unknown"
	}

	info.Packages = getPackageCount()

	return info
}

func extractQuotedFields(line string) []string {
	var fields []string
	inQuote := false
	current := ""
	for _, r := range line {
		if r == '"' {
			inQuote = !inQuote
			if !inQuote && current != "" {
				fields = append(fields, current)
				current = ""
			}
		} else if inQuote {
			current += string(r)
		}
	}
	return fields
}

func cleanGPUName(raw string) string {
	trash := map[string]bool{
		"corporation": true, "inc.": true, "inc": true, "ltd": true,
		"limited": true, "co.": true, "gaming": true, "super": true,
		"ultra": true, "prime": true, "speedster": true, "merc": true,
		"black": true, "white": true, "edition": true, "series": true,
		"radeon": true, "geforce": true, "graphics": true, "card": true,
		"vga": true, "controller": true, "compatible": true,
	}

	vendors := map[string]string{
		"nvidia": "NVIDIA", "amd": "AMD", "intel": "Intel",
		"asus": "ASUS", "msi": "MSI", "gigabyte": "Gigabyte",
		"evga": "EVGA", "sapphire": "Sapphire", "xfx": "XFX",
		"zotac": "ZOTAC", "palit": "Palit", "galax": "GALAX",
	}

	words := strings.Fields(raw)
	if len(words) == 0 {
		return raw
	}

	// Ищем производителя, но не сохраняем – убираем его из слов
	for i, w := range words {
		lower := strings.ToLower(w)
		if _, ok := vendors[lower]; ok {
			words = append(words[:i], words[i+1:]...)
			break
		}
	}
	// Если производитель не распознан – пробуем убрать первый элемент,
	// если он похож на название вендора (например, "XFX")
	if len(words) > 0 {
		first := words[0]
		firstLower := strings.ToLower(first)
		if strings.HasSuffix(firstLower, "corporation") ||
			strings.HasSuffix(firstLower, "inc.") ||
			strings.HasSuffix(firstLower, "inc") ||
			strings.HasSuffix(firstLower, "ltd") ||
			strings.HasSuffix(firstLower, "limited") ||
			strings.HasSuffix(firstLower, "co.") ||
			(firstLower == "xfx" || firstLower == "msi" || firstLower == "asus" ||
				firstLower == "evga" || firstLower == "zotac" || firstLower == "gigabyte" ||
				firstLower == "palit" || firstLower == "galax" || firstLower == "sapphire") {
			words = words[1:]
		}
	}

	// Убираем мусорные слова
	badPrefixes := map[string]bool{"merc": true, "speedster": true, "series": true, "edition": true}
	filtered := []string{}
	skipNext := false
	for i, w := range words {
		if skipNext {
			skipNext = false
			continue
		}
		lower := strings.ToLower(w)
		if badPrefixes[lower] && i+1 < len(words) && isNumeric(words[i+1]) {
			skipNext = true
			continue
		}
		if trash[lower] {
			continue
		}
		filtered = append(filtered, w)
	}

	model := strings.Join(filtered, " ")
	if model == "" {
		return raw
	}
	return model
}

func isNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return len(s) > 0
}

func getGPU() string {
	cmd := exec.Command("lspci", "-mm", "-d", "::0300")
	out, err := cmd.Output()
	if err == nil {
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			fields := extractQuotedFields(line)
			var vendor, device string

			if len(fields) >= 5 {
				possibleVendor := fields[2]
				if strings.Contains(possibleVendor, "Corporation") ||
					strings.Contains(possibleVendor, "Inc.") ||
					len(possibleVendor) < 20 {
					vendor = possibleVendor
					device = fields[3]
				} else {
					vendor = fields[3]
					if len(fields) > 4 {
						device = fields[4]
					}
				}
			} else if len(fields) >= 4 {
				device = fields[3]
			}

			if start := strings.Index(device, "["); start != -1 {
				if end := strings.Index(device, "]"); end != -1 && end > start {
					device = device[start+1 : end]
				}
			}
			device = strings.TrimSuffix(device, " [VGA controller]")
			device = strings.TrimSuffix(device, " [Display controller]")
			device = strings.TrimSpace(device)

			full := vendor + " " + device
			if vendor == "" {
				full = device
			}
			if full != "" {
				cleaned := cleanGPUName(full)
				if cleaned != "" && cleaned != "Unknown" {
					return cleaned
				}
				return full
			}
		}
	}

	if cards, err := os.ReadDir("/sys/class/drm"); err == nil {
		for _, card := range cards {
			if strings.HasPrefix(card.Name(), "card") {
				uevent, err := os.ReadFile("/sys/class/drm/" + card.Name() + "/device/uevent")
				if err == nil {
					for _, ln := range strings.Split(string(uevent), "\n") {
						if strings.HasPrefix(ln, "PCI_ID=") {
							return strings.TrimPrefix(ln, "PCI_ID=")
						}
					}
				}
			}
		}
	}
	return "Unknown"
}

func detectWM() string {
	if wm := os.Getenv("XDG_CURRENT_DESKTOP"); wm != "" {
		return wm
	}
	if wm := os.Getenv("DESKTOP_SESSION"); wm != "" {
		return wm
	}
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		switch {
		case os.Getenv("SWAYSOCK") != "":
			return "Sway"
		case os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "":
			return "Hyprland"
		case os.Getenv("NIRI_SOCKET") != "":
			return "Niri"
		case os.Getenv("RIVER_SOCKET") != "":
			return "River"
		default:
			return "Wayland"
		}
	}
	return "Unknown"
}

func getPackageCount() string {
	distro := ""
	if f, err := os.Open("/etc/os-release"); err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			if strings.HasPrefix(line, "ID=") {
				distro = strings.Trim(strings.TrimPrefix(line, "ID="), `"'`)
				break
			}
		}
		f.Close()
	}

	type pm struct {
		cmd       string
		args      []string
		parseFunc func(string) string
	}
	pms := map[string]pm{
		"arch":     {"/usr/bin/pacman", []string{"-Q"}, func(out string) string { return strings.Split(out, "\n")[0] }},
		"debian":   {"/usr/bin/dpkg", []string{"-l"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")-5) }},
		"ubuntu":   {"/usr/bin/dpkg", []string{"-l"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")-5) }},
		"fedora":   {"/usr/bin/rpm", []string{"-qa"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"rhel":     {"/usr/bin/rpm", []string{"-qa"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"opensuse": {"/usr/bin/rpm", []string{"-qa"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"gentoo":   {"/usr/bin/qlist", []string{"-I"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"alpine":   {"/sbin/apk", []string{"list", "--installed"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"void":     {"/usr/bin/xbps-query", []string{"-l"}, func(out string) string { return fmt.Sprintf("%d", strings.Count(out, "\n")) }},
		"nixos":    {"/run/current-system/sw/bin/nix", []string{"list-generations"}, func(out string) string { return "nix" }},
	}

	var pmInfo pm
	for id, p := range pms {
		if strings.Contains(distro, id) {
			pmInfo = p
			break
		}
	}
	if pmInfo.cmd == "" {
		for _, cmd := range [][]string{
			{"pacman", "-Q"},
			{"dpkg", "-l"},
			{"rpm", "-qa"},
			{"qlist", "-I"},
			{"apk", "list", "--installed"},
			{"xbps-query", "-l"},
		} {
			if path, err := exec.LookPath(cmd[0]); err == nil {
				pmInfo.cmd = path
				pmInfo.args = cmd[1:]
				pmInfo.parseFunc = func(out string) string {
					return fmt.Sprintf("%d", strings.Count(out, "\n"))
				}
				break
			}
		}
	}

	if pmInfo.cmd != "" {
		out, err := exec.Command(pmInfo.cmd, pmInfo.args...).Output()
		if err == nil {
			return pmInfo.parseFunc(string(out))
		}
	}
	return "Unknown"
}
