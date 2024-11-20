CREATE TABLE jenis_acara (
    id_jenis_acara SERIAL PRIMARY KEY,         -- Primary Key dengan auto-increment
    nama_jenis_acara VARCHAR(255) NOT NULL,    -- Nama jenis acara, wajib diisi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu pembuatan otomatis
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Waktu pembaruan otomatis
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON jenis_acara
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
