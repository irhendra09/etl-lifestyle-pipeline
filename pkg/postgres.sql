-- Menghapus tabel jika sudah ada (berguna jika Anda mematikan dan menghidupkan ulang container untuk testing)
DROP TABLE IF EXISTS fact_content_metrics CASCADE;
DROP TABLE IF EXISTS bridge_content_tag CASCADE;
DROP TABLE IF EXISTS dim_tag CASCADE;
DROP TABLE IF EXISTS dim_content CASCADE;
DROP TABLE IF EXISTS dim_user CASCADE;

-- ==========================================
-- TABEL DIMENSI (Menyimpan Teks & Deskripsi)
-- ==========================================

-- 1. Tabel Dimensi: Pengguna / Penulis Ulasan
CREATE TABLE dim_user (
                          user_sk SERIAL PRIMARY KEY,              -- Surrogate Key (Kunci internal DW)
                          original_user_id INT UNIQUE NOT NULL,    -- ID asli dari API eksternal
                          username VARCHAR(255),
                          inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Tabel Dimensi: Konten (Judul & Teks Ulasan Gaya Hidup)
CREATE TABLE dim_content (
                             content_sk SERIAL PRIMARY KEY,
                             original_post_id INT UNIQUE NOT NULL,    -- UNIQUE sangat penting untuk fitur UPSERT (ON CONFLICT) nanti
                             title VARCHAR(500),
                             body_text TEXT,
                             inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Tabel Dimensi: Tag / Kategori
CREATE TABLE dim_tag (
                         tag_sk SERIAL PRIMARY KEY,
                         tag_name VARCHAR(100) UNIQUE NOT NULL
);

-- 4. Tabel Relasi (Bridge): Menyelesaikan Many-to-Many antara Konten dan Tag
CREATE TABLE bridge_content_tag (
                                    content_sk INT REFERENCES dim_content(content_sk) ON DELETE CASCADE,
                                    tag_sk INT REFERENCES dim_tag(tag_sk) ON DELETE CASCADE,
                                    PRIMARY KEY (content_sk, tag_sk)
);

-- ==========================================
-- TABEL FAKTA (Menyimpan Angka & Metrik)
-- ==========================================

-- 5. Tabel Fakta: Kinerja Konten Harian
CREATE TABLE fact_content_metrics (
                                      metric_id SERIAL PRIMARY KEY,
                                      content_sk INT REFERENCES dim_content(content_sk) ON DELETE CASCADE,
                                      user_sk INT REFERENCES dim_user(user_sk) ON DELETE CASCADE,
                                      total_views INT DEFAULT 0,
                                      total_likes INT DEFAULT 0,
                                      total_dislikes INT DEFAULT 0,
                                      extraction_date DATE NOT NULL,           -- Mencatat kapan metrik ini ditarik
                                      inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Memastikan satu konten hanya punya satu baris data metrik per harinya
                                      UNIQUE (content_sk, extraction_date)
);

-- ==========================================
-- INDEXING (Untuk mempercepat Query AI)
-- ==========================================
CREATE INDEX idx_fact_extraction_date ON fact_content_metrics(extraction_date);
CREATE INDEX idx_dim_tag_name ON dim_tag(tag_name);