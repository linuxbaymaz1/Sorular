1. İşlem Hızı (Throughput) AnaliziGrafik: 
QuestDB: Turuncu çizgi diğerlerinden tamamen kopmuş durumda. Isındıkça hızı artıyor ve $14,000$ ops/s seviyesinin üzerinde, 600 saniye boyunca hiç yorulmadan devam ediyor.
TimescaleDB: Başlangıçta fena değil ($4,500$ ops/s) ancak süre uzadıkça performans kaybı yaşıyor. 600. saniyede hızın neredeyse yarıya düştüğünü net bir şekilde görebiliyoruz.
ClickHouse: En altta ama kaya gibi sağlam bir çizgi çiziyor. %40 update yükü altında hızı düşük olsa da hiç dalgalanma yaşamıyor.
2. Bellek (Memory) Yönetimi: 2GB Sınırı: 
Grafikteki barlara bakıldığında, QuestDB ve TimescaleDB'nin 300 ve 600 saniyelik testlerde $2048$ MB (2GB) limitine tam olarak yaslandığını görüyoruz
.QuestDB Farkı: QuestDB bu limite dayanmasına rağmen hızını korurken (off-heap başarısı), TimescaleDB aynı sınırda performans kaybı yaşamaya başlıyor.
3. Depolama (Disk) Verimliliği:
 (Logaritmik Ölçek)
ClickHouse Mucizesi: Mavi çizgi en dipte. QuestDB'nin $11$ GB yer kapladığı aynı veri setini ClickHouse sadece $\approx 1.5$ GB ile saklıyor.
 Bu, devasa log dosyaları için ClickHouse'un neden tercih edildiğinin kanıtı.
QuestDB: Hızın bedelini diskle ödüyor. Ham veriyi diske olduğu gibi yığdığı için disk kullanımı doğrusal olarak tırmanıyor.
Rapor İçin Teknik Özet:"Yapılan stres testleri sonucunda QuestDB yüksek yoğunluklu veri akışları (real-time monitoring) için en performanslı çözüm olarak öne çıkarken, ClickHouse depolama maliyetlerini %85 oranında optimize ederek uzun süreli log saklama (cold storage/forensics) senaryolarında üstünlüğünü kanıtlamıştır. TimescaleDB ise ilişkisel veri yapısıyla uyumluluğu korumasına rağmen, kısıtlı bellek kaynakları altında sürdürülebilir hız performansında geride kalmıştır."
