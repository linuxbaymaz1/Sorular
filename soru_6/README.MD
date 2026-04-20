# DHCP Log Stream Parser

Bu araç, sürekli akan (streaming) ağ logları içerisinden DHCPREQUEST (Type 3) mesajlarını anlık olarak tespit eder ve ilgili alanları ayrıştırarak raporlar.

## Yakalanan Alanlar

- Timestamp (Zaman Damgası)
- Client-identifier
- Request IP Address
- Vendor Class Identifier
- Host Name

## Kullanım

git clone [https://github.com/KULLANICI_ADIN/dhcp-parser.git](https://github.com/KULLANICI_ADIN/dhcp-parser.git)
cd dhcp-parser
go run main.go
