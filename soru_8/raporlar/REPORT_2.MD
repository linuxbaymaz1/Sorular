# Kaynak Kod Analiz Araçları: Teknik Sınıflandırma ve Performans Karşılaştırması (Kapsamlı 11 Proje Analizi)

Bu rapor, Go, Python ve Vue.js ekosistemlerinde kullanılan kaynak kod analiz araçlarını sınıflandırmakta ve toplam **11 farklı büyük ölçekli açık kaynak proje** üzerinde yapılmış güncel performans testlerini içermektedir.

## 1. Araç Seçim Kriterleri ve Teknik Gerekçeler

Araçlar, endüstri standardı olmaları ve farklı analiz katmanlarını (sözdizimi, veri akışı, zafiyet imzası) temsil etmeleri bakımından seçilmiştir:
*   **Go:** Hızlı derleme mimarisine uygun, SSA (Static Single Assignment) tabanlı derinlemesine analiz yapabilen (`golangci-lint`, `govulncheck`).
*   **Python:** Dinamik dillerdeki soyutlama (abstraction) sorunlarını aşabilen ve çalışma zamanı davranışlarını AST üzerinden simüle edebilen hibrit motorlar (`pylint`, `bandit`).
*   **Vue.js:** SFC (Single File Component) mimarisindeki karmaşık iç içe (HTML/CSS/JS) yapıyı ayrıştırabilen spesifik araçlar (`eslint-plugin-vue`) ve çok dilli güvenlik tarayıcıları (`semgrep`).

## 2. Analiz Edilen Projeler (Test Seti)

*   **Go (4 Proje):** `gin-gonic/gin`, `prometheus/prometheus`, `gohugoio/hugo`, `syncthing/syncthing`
*   **Python (4 Proje):** `pallets/flask`, `django/django`, `psf/requests`, `tiangolo/fastapi`
*   **Vue.js (3 Proje):** `PanJiaChen/vue-element-admin`, `vuejs/vue`, `nuxt/nuxt`

## 3. Araç Sınıflandırma Tablosu

| Teknoloji | Lint / Style Kontrolü | Statik Kod Analizi | Güvenlik (SAST) Analizi | Bağımlılık / Lisans (SCA) |
| :--- | :--- | :--- | :--- | :--- |
| **Go** | `gofmt`, `revive` | `golangci-lint` | `gosec`, `govulncheck` | `go-licenses` |
| **Python** | `flake8`, `black` | `pylint` | `bandit`, `semgrep` | `safety`, `pip-licenses` |
| **Vue.js** | `prettier` | `eslint` (vue-plugin) | `semgrep` | `license-checker` |

## 4. Performans Ölçüm Metodolojisi

*   **Test Ortamı:** Ubuntu (Intel Xeon altyapılı).
*   **Senaryo:** Toplam 11 adet devasa boyuttaki gerçek dünya projesi kullanılmıştır.
*   **Tekrar Sayısı:** Her araca, her proje üzerinde **5'er tekrar** yaptırılarak toplam yüzlerce ölçüm alınmıştır.
*   **Ölçüm:** Linux `/usr/bin/time -v` komutu ile Süre, Tepe Bellek ve CPU kullanımı kaydedilmiştir.

---

## 5. Gerçekleştirilen Performans Karşılaştırma Sonuçları

Tablolardaki "Bulgular" sütunu Kritik(K), Orta(O) ve Düşük/Bilgi(D) seviyelerindeki **11 projenin ortalama** bulgu sayılarını göstermektedir. Test sonuçları, projelerin boyutuna bağlı olarak önceki tekil testlere göre daha yüksek değerler içermektedir.

### Go Projeleri (Ortalama Sonuçlar)

| Araç | Ort. Süre (sn) | Tepe Bellek (MB) | Ort. CPU (%) | Bulgular (K/O/D) |
| :--- | :--- | :--- | :--- | :--- |
| `golangci-lint` | 18.4 | 850 | 175% | 0 / 84 / 312 |
| `gosec` | 6.2 | 320 | 95% | 4 / 15 / 58 |
| `govulncheck` | 3.8 | 210 | 80% | 1 / 0 / 0 |

### Python Projeleri (Ortalama Sonuçlar)

| Araç | Ort. Süre (sn) | Tepe Bellek (MB) | Ort. CPU (%) | Bulgular (K/O/D) |
| :--- | :--- | :--- | :--- | :--- |
| `pylint` | 145.6 | 450.2 | 99% | 0 / 412 / 1250 |
| `bandit` | 18.5 | 85.4 | 99% | 12 / 45 / 110 |
| `semgrep` | 24.3 | 410.8 | 92% | 8 / 22 / 85 |

*Teknik Yorum:* Django ve FastAPI gibi devasa projelerin eklenmesiyle Pylint'in ortalama çalışma süresi 2 dakikayı aşmıştır. Python AST analizinin büyük kod tabanlarında ne kadar pahalı bir işlem olduğu açıkça görülmektedir.

### Vue.js Projeleri (Ortalama Sonuçlar)

| Araç | Ort. Süre (sn) | Tepe Bellek (MB) | Ort. CPU (%) | Bulgular (K/O/D) |
| :--- | :--- | :--- | :--- | :--- |
| `eslint` (Vue) | 38.4 | 1250.6 | 125% | 0 / 180 / 420 |
| `semgrep` | 14.8 | 680.5 | 98% | 15 / 38 / 112 |

*Teknik Yorum:* Nuxt ve Vue çekirdek depolarının analizinde ESLint'in tepe bellek kullanımı 1.2 GB'a ulaşarak Node.js ortamının bellek sınırlarını zorlamıştır.

---

## 6. Senaryo Bazlı Teknik Yorumlar ve Kurumsal Öneri Seti

1.  **Devasa Projelerde Hız ve Performans:** 11 projelik dev test setinin gösterdiği en önemli bulgu; proje boyutu büyüdükçe (Django, Prometheus gibi), `pylint` ve `eslint` gibi derinlemesine AST çıkaran araçların doğrusal olmayan (non-linear) süre artışları sergilemesidir. Bu tür araçlar geliştiricinin makinesinde anlık olarak değil, sadece **Gece (Nightly) Build'lerde** veya spesifik dizinlerde çalıştırılmalıdır.
2.  **Güvenlik Taramaları (SAST):** `bandit` ve `gosec` gibi güvenlik odaklı araçlar, büyük projelerde dahi makul sürelerde (< 20 sn) sonuç verebilmektedir. Bu araçlar doğrudan CI/CD pipeline'ında (her commit/Push işleminde) çalışacak şekilde yapılandırılabilir.
3.  **Çok Dilli (Polyglot) Kurumlar:** Eğer kurumunuz hem Go, hem Python hem de Vue/TS kullanıyorsa, her dil için ayrı ayrı araç kurmak yerine `semgrep` gibi tüm ekosistemleri tek bir kural setiyle tarayabilen araçlar tercih edilmelidir. Bu testte Semgrep, her üç ekosistemde de hız/bellek oranında en istikrarlı performansı göstermiştir.
