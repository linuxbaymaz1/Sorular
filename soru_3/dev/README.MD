# SSH Network Scanner

Bu proje 172.29.0.0/16 ağ bloğunu tarar, SSH açık cihazları bulur ve bu cihazlara bağlanarak verilen komutu çalıştırır. Çıktılar terminale basılır.

Çalıştırma 
git clone https://github.com/kullanici/proje.git
cd proje
go mod init scanner
go mod tidy
go get golang.org/x/crypto/ssh
go run g.go
go run g.go "komut"



8. Kodun genel mantığı
-172.29.0.0/16 taranır
-SSH açık makineler bulunur
-Her makineye paralel bağlanılır
-Komut çalıştırılır
-Çıktı terminale basılır

