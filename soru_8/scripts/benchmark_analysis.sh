#!/bin/bash
# benchmark_analysis.sh

TARGET_PROJECT=$1
TOOL_NAME=$2
# Execute the rest of the arguments as the command
shift 2
COMMAND="$@"

echo "Analiz başlatılıyor: $TOOL_NAME projesi $TARGET_PROJECT"
echo "Komut: $COMMAND"
echo "--------------------------------------------------------"

for i in {1..5}
do
    echo "Tekrar $i..."
    # /usr/bin/time -v ile CPU, bellek ve zaman verilerini topla
    /usr/bin/time -v $COMMAND > "metrics_${TOOL_NAME}_${i}.out" 2> "metrics_${TOOL_NAME}_${i}.txt"
    sleep 1 # Sistem dinlenmesi için kısa ara
done

# Sonuçların özetlenmesi (Örnek: Ortalama zaman ve tepe bellek)
echo "Analiz Sonuç Özeti ($TOOL_NAME):"
grep "Elapsed (wall clock) time" metrics_${TOOL_NAME}_*.txt | awk '{print $NF}' | sed 's/0://'
grep "Maximum resident set size" metrics_${TOOL_NAME}_*.txt | awk '{print "Bellek (KB): " $NF}'
grep "Percent of CPU this job got" metrics_${TOOL_NAME}_*.txt | awk '{print "CPU (%): " $NF}'
