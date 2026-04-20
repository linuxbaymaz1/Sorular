2. EAP-TLS, EAP-TTLS,PEAPvebenzeriRADIUS/EAPtabanlı kimlik doğrulama yöntemleri arasındaki temel farklar nelerdir

                    EAP-TLS
Temel Özellik: Karşılıklı  Kimlik Doğrulama anlamına gelir.
Sertifika Gereksinimi: İki trarafta sertifika kullanır.
Güvenlik:Parola kullanılmaz.

                      PEAP
Temel Özellik: Korumalı Tünel.
Sertifika Gereksinimi: Sadece sunucu tarafında sertifika kullanılır.
Kimlik Doğrulama: TLS tüneli içerisinde genellikle kullanıcı adı ve parola kullanılır.
Güvenlik: EAP-TLS deki gibi kullanıcı tarafında sertifika olmadığı içib biraz daha az güvenlidir

                               EAP-TTLS  
Temel Özellik: Esnek 
Sertifika Gereksinimi: Sadece sunucu tarafında sertifika gereklidir.
Kimlik Doğrulama: PEAP gibidir, ama ek olarak eski protokolleri de destekler. 
Güvenlik: Yüksek güvenlik sağlar.
Kullanım Senaryosu: Eduroamda kullanılır. 
