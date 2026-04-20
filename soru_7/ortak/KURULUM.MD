docker-compose up -d
go mod init db_bench
go mod tidy
docker exec -it tsdb_bench psql -U postgres -d postgres -c "CREATE TABLE IF NOT EXISTS bench (id BIGINT, ts TIMESTAMPTZ NOT NULL, val DOUBLE PRECISION, PRIMARY KEY (id, ts)); SELECT create_hypertable('bench', 'ts', if_not_exists => TRUE); CREATE INDEX IF NOT EXISTS idx_bench_id ON bench(id);"
docker exec -it tsdb_bench sh -c "PGPASSWORD=quest psql -h quest_bench -p 8812 -U admin -d qdb -c \"CREATE TABLE IF NOT EXISTS bench (id LONG, ts TIMESTAMP, val DOUBLE) TIMESTAMP(ts) PARTITION BY DAY WAL; ALTER TABLE bench DEDUP ENABLE UPSERT KEYS(ts, id);\""
docker exec -it ch_bench clickhouse-client --query "CREATE TABLE IF NOT EXISTS bench (id UInt64, ts DateTime, val Float64) ENGINE = ReplacingMergeTree() ORDER BY id"
python3 monitor.py

#test başlatılmadan önce istenen db'nin monitor.py'si bu klasöre taşınır.

#İşletim sistemi seviyesinde "sudo apt install docker.io docker-compose golang-go" kullanılır.
