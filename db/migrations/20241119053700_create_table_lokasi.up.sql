create Table lokasi(
    id_lokasi SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Membuat trigger yang menggunakan fungsi di atas
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON lokasi
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();