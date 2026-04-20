kali@82:~/project-root/soru_4/dev$ go run main.go  --help
Usage of /tmp/go-build418301171/b001/exe/main:
  -interval duration
        Scan interval (default 10s)
  -out string
        Log file (default "down_hosts.log")
  -subnet string
        CIDR network
  -workers int
        Concurrency (default 512)
kali@82:~/project-root/soru_4/dev$



-interval süresi iki tarama arasındaki süreyi seçmemizi sağlıyor 

-out parametresi sonuçların hangi dosyaya yazacağımızı gösteriyor

-subnet direk subnet (192.168.1.0/24 ev ağı gib) vermemizi sağlıyor

-workers işçi sayısı.ne kadar çok arttırırsak bir o kadar hızlı bir o kadar kaynak tüketir



