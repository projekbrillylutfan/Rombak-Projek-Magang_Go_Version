CREATE TABLE agenda (
    id_agenda SERIAL PRIMARY KEY,                 -- Primary Key dengan auto-increment
    id_bupati INT NOT NULL,                       -- Foreign Key ke tabel Bupati
    nama_agenda VARCHAR(255) NOT NULL,           -- Nama agenda, wajib diisi
    deskripsi TEXT,                               -- Deskripsi, opsional
    id_lokasi INT NOT NULL,                       -- Foreign Key ke tabel Lokasi
    id_jenis_acara INT NOT NULL,                  -- Foreign Key ke tabel Jenis Acara
    tanggal_mulai TIMESTAMP NOT NULL,             -- Tanggal mulai acara
    tanggal_selesai TIMESTAMP NOT NULL,           -- Tanggal selesai acara
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu pembuatan otomatis
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu pembaruan otomatis

    -- Definisi Foreign Keys
    CONSTRAINT fk_bupati FOREIGN KEY (id_bupati) REFERENCES bupati (id_bupati),
    CONSTRAINT fk_lokasi FOREIGN KEY (id_lokasi) REFERENCES lokasi (id_lokasi),
    CONSTRAINT fk_jenis_acara FOREIGN KEY (id_jenis_acara) REFERENCES jenis_acara (id_jenis_acara)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON agenda
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
