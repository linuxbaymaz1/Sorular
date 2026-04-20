#!/bin/bash

# Setup Go module properly
cd /home/kali/calisma/real_projects/gin
go mod tidy
cd /home/kali/calisma

# Run Go Benchmarks
echo "--- GO BENCHMARKS ---" > real_results.txt
./benchmark_analysis.sh real_projects/gin golangci-lint "golangci-lint run /home/kali/calisma/real_projects/gin/..." >> real_results.txt
./benchmark_analysis.sh real_projects/gin gosec "gosec /home/kali/calisma/real_projects/gin/..." >> real_results.txt
./benchmark_analysis.sh real_projects/gin govulncheck "govulncheck -C /home/kali/calisma/real_projects/gin ." >> real_results.txt
./benchmark_analysis.sh real_projects/gin go-licenses "go-licenses check github.com/gin-gonic/gin" >> real_results.txt

# Run Python Benchmarks
echo "--- PYTHON BENCHMARKS ---" >> real_results.txt
source /home/kali/calisma/venv/bin/activate
./benchmark_analysis.sh real_projects/flask pylint "pylint /home/kali/calisma/real_projects/flask/src/flask" >> real_results.txt
./benchmark_analysis.sh real_projects/flask bandit "bandit -r /home/kali/calisma/real_projects/flask/src" >> real_results.txt
./benchmark_analysis.sh real_projects/flask safety "safety check" >> real_results.txt
./benchmark_analysis.sh real_projects/flask semgrep-py "semgrep --config auto /home/kali/calisma/real_projects/flask/src" >> real_results.txt
deactivate

# Run Vue/Node Benchmarks
echo "--- VUE BENCHMARKS ---" >> real_results.txt
./benchmark_analysis.sh real_projects/vue-element-admin semgrep-vue "semgrep --config auto /home/kali/calisma/real_projects/vue-element-admin/src" >> real_results.txt
./benchmark_analysis.sh real_projects/vue-element-admin eslint "eslint --no-eslintrc --env browser --parser-options=ecmaVersion:2020 /home/kali/calisma/real_projects/vue-element-admin/src" >> real_results.txt

echo "ALL DONE"
