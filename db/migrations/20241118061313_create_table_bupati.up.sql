-- Membuat tabel Bupati
CREATE TABLE bupati (
    id_bupati SERIAL PRIMARY KEY,         -- Primary key dengan auto-increment
    nama VARCHAR(255) NOT NULL,           -- Nama Bupati (wajib diisi)
    periode_jabatan VARCHAR(100) NOT NULL, -- Periode Jabatan (wajib diisi)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Tanggal dibuat (default: waktu saat ini)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Tanggal diperbarui (otomatis diperbarui lewat trigger)
);

-- Membuat fungsi untuk memperbarui kolom `updated_at`
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Membuat trigger yang menggunakan fungsi di atas
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON bupati
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
