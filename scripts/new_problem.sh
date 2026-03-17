#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 2 ]]; then
  echo "Usage: $0 <id> <slug>"
  echo "Example: $0 1 two-sum"
  exit 1
fi

id_raw="$1"
slug_raw="$2"

if ! [[ "$id_raw" =~ ^[0-9]+$ ]]; then
  echo "Error: id must be numeric"
  exit 1
fi

id=$(printf "%04d" "$id_raw")
slug=$(echo "$slug_raw" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/_/g; s/^_+//; s/_+$//')

if [[ -z "$slug" ]]; then
  echo "Error: slug is empty after normalization"
  exit 1
fi

pkg="p${id}_${slug}"
dir="problems/${id}_${slug}"

if [[ -e "$dir" ]]; then
  echo "Error: ${dir} already exists"
  exit 1
fi

mkdir -p "$dir"

func_name=$(echo "$slug" | awk -F'_' '{out=""; for(i=1;i<=NF;i++){ out=out toupper(substr($i,1,1)) substr($i,2)}; print out}')

cat > "${dir}/solution.go" <<EOF
package ${pkg}

func ${func_name}() {
	// TODO: implement solution
}
EOF

cat > "${dir}/solution_test.go" <<EOF
package ${pkg}

import "testing"

func Test${func_name}(t *testing.T) {
	t.Skip("add test cases")
}
EOF

echo "Created ${dir}"