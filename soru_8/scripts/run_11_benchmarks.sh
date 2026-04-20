#!/bin/bash

RESULTS_FILE="/home/kali/calisma/massive_results.txt"
echo "--- 11 PROJELİK DEVASA BENCHMARK BAŞLADI ---" > $RESULTS_FILE

# GO PROJECTS
GO_PROJECTS=("gin" "prometheus" "hugo" "syncthing")
for PROJ in "${GO_PROJECTS[@]}"; do
    echo "======================================" >> $RESULTS_FILE
    echo "GO PROJESİ: $PROJ" >> $RESULTS_FILE
    echo "======================================" >> $RESULTS_FILE
    
    cd "/home/kali/calisma/real_projects/$PROJ"
    go mod tidy >/dev/null 2>&1
    cd /home/kali/calisma
    
    ./benchmark_analysis.sh real_projects/$PROJ golangci-lint "golangci-lint run /home/kali/calisma/real_projects/$PROJ/..." >> $RESULTS_FILE
    ./benchmark_analysis.sh real_projects/$PROJ gosec "gosec /home/kali/calisma/real_projects/$PROJ/..." >> $RESULTS_FILE
    ./benchmark_analysis.sh real_projects/$PROJ govulncheck "govulncheck -C /home/kali/calisma/real_projects/$PROJ ." >> $RESULTS_FILE
done

# PYTHON PROJECTS
PYTHON_PROJECTS=("flask" "django" "requests" "fastapi")
source /home/kali/calisma/venv/bin/activate
for PROJ in "${PYTHON_PROJECTS[@]}"; do
    echo "======================================" >> $RESULTS_FILE
    echo "PYTHON PROJESİ: $PROJ" >> $RESULTS_FILE
    echo "======================================" >> $RESULTS_FILE
    
    # Her projenin kök dizinini tara
    ./benchmark_analysis.sh real_projects/$PROJ pylint "pylint /home/kali/calisma/real_projects/$PROJ" >> $RESULTS_FILE
    ./benchmark_analysis.sh real_projects/$PROJ bandit "bandit -r /home/kali/calisma/real_projects/$PROJ" >> $RESULTS_FILE
    ./benchmark_analysis.sh real_projects/$PROJ semgrep-py "semgrep --config auto /home/kali/calisma/real_projects/$PROJ" >> $RESULTS_FILE
done
deactivate

# VUE/NODE PROJECTS
VUE_PROJECTS=("vue-element-admin" "vue" "nuxt")
for PROJ in "${VUE_PROJECTS[@]}"; do
    echo "======================================" >> $RESULTS_FILE
    echo "VUE/NODE PROJESİ: $PROJ" >> $RESULTS_FILE
    echo "======================================" >> $RESULTS_FILE
    
    ./benchmark_analysis.sh real_projects/$PROJ semgrep-vue "semgrep --config auto /home/kali/calisma/real_projects/$PROJ" >> $RESULTS_FILE
    ./benchmark_analysis.sh real_projects/$PROJ eslint "eslint --no-eslintrc --env browser --parser-options=ecmaVersion:2020 /home/kali/calisma/real_projects/$PROJ" >> $RESULTS_FILE
done

echo "ALL MASSIVE BENCHMARKS DONE" >> $RESULTS_FILE
echo "ALL DONE"
