#!/usr/bin/env zsh

fifo() {
	local -r temp="$(mktemp)"
	rm "$temp"
	mkfifo "$temp"
	exec 3<>"${temp}"
	rm -f "$temp"
}

serve() {
	fifo
	nc -ul 5050 <&3 | while read -r line; do
		hexdump -C <<< $line
		cat "data" | xxd -r -p >&3
	done

	exec 3>&-
}

serve
