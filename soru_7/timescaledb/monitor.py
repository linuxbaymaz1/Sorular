import subprocess
import time
import docker
import csv
import os

client = docker.from_env()

# DB Ayarları
databases = [
    {"name": "TimescaleDB", "drv": "pgx", "dsn": "postgres://postgres:bench_pass@localhost:5433/postgres?sslmode=disable", "typ": "pg", "container": "tsdb_bench", "path": "/var/lib/postgresql/data"},
    ]

durations = [60, 300, 600]
repeats = 5

def get_db_size(container_name, path):
    try:
        out = client.containers.get(container_name).exec_run(f"du -sm {path}")
        return out.output.decode().split()[0]
    except: return 0

def get_mem_usage(container_name):
    try:
        stats = client.containers.get(container_name).stats(stream=False)
        return stats['memory_stats']['usage'] / (1024 * 1024)
    except: return 0

with open('performance_report.csv', 'w') as f:
    writer = csv.writer(f)
    writer.writerow(["DB", "Duration", "Run", "Total_Inserts", "Total_Ops", "Throughput_Ops_s", "Mem_MB", "Disk_MB"])

    for db in databases:
        for dur in durations:
            for r in range(1, repeats + 1):
                print(f"Testing {db['name']} - {dur}s - Run {r}...")
                
                # Go Benchmark Başlat
                cmd = f"go run bench.go -db {db['name']} -drv {db['drv']} -dsn '{db['dsn']}' -typ {db['typ']} -dur {dur}"
                proc = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE)
                
                # Test sırasında metrikleri topla
                time.sleep(dur / 2)
                mem = get_mem_usage(db['container'])
                
                stdout, _ = proc.communicate()
                i, u, s = map(int, stdout.decode().split(','))
                
                total_ops = i + u + s
                disk = get_db_size(db['container'], db['path'])
                
                writer.writerow([db['name'], dur, r, i, total_ops, total_ops/dur, mem, disk])
                f.flush()

print("Benchmark Tamamlandı! Sonuçlar performance_report.csv dosyasında.")
