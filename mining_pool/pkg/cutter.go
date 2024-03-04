package pkg


func CutIPAddress(ipAddr string) string {
	var result []byte
	for x := range ipAddr {
		if string(ipAddr[x]) == ":" {
			break
		}
		result = append(result, ipAddr[x])
	}
	return string(result)
}
