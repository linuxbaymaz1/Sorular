#!/bin/bash

# Setup Go module properly
cd /home/kali/calisma/projects/go-app
go mod tidy
cd /home/kali/calisma

# Setup Python requirements
echo "Flask==3.0.0" > /home/kali/calisma/projects/python-app/requirements.txt
echo "sqlite3" >> /home/kali/calisma/projects/python-app/requirements.txt

# Run Go Benchmarks
echo "--- GO BENCHMARKS ---" > results.txt
./benchmark_analysis.sh projects/go-app golangci-lint "golangci-lint run /home/kali/calisma/projects/go-app/..." >> results.txt
./benchmark_analysis.sh projects/go-app gosec "gosec /home/kali/calisma/projects/go-app/..." >> results.txt
./benchmark_analysis.sh projects/go-app govulncheck "govulncheck -C /home/kali/calisma/projects/go-app ." >> results.txt
# go-licenses requires the module to be built/downloaded
cd /home/kali/calisma/projects/go-app && go build . && cd /home/kali/calisma
./benchmark_analysis.sh projects/go-app go-licenses "go-licenses check github.com/gin-gonic/gin" >> results.txt

# Run Python Benchmarks
echo "--- PYTHON BENCHMARKS ---" >> results.txt
source /home/kali/calisma/venv/bin/activate
./benchmark_analysis.sh projects/python-app pylint "pylint /home/kali/calisma/projects/python-app/app.py" >> results.txt
./benchmark_analysis.sh projects/python-app bandit "bandit -r /home/kali/calisma/projects/python-app/" >> results.txt
./benchmark_analysis.sh projects/python-app safety "safety check -r /home/kali/calisma/projects/python-app/requirements.txt" >> results.txt
./benchmark_analysis.sh projects/python-app semgrep-py "semgrep --config auto /home/kali/calisma/projects/python-app" >> results.txt
deactivate

# Run Vue/Node Benchmarks
echo "--- VUE BENCHMARKS ---" >> results.txt
./benchmark_analysis.sh projects/vue-app semgrep-vue "semgrep --config auto /home/kali/calisma/projects/vue-app" >> results.txt
# eslint requires local installation usually, but we have global. 
# We'll just run eslint on the file
./benchmark_analysis.sh projects/vue-app eslint "eslint --no-eslintrc --env browser --parser-options=ecmaVersion:2020 /home/kali/calisma/projects/vue-app/App.vue" >> results.txt

echo "ALL DONE"
