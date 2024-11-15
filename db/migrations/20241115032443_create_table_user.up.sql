CREATE TYPE user_role AS ENUM ('ADMIN', 'PASIEN');

CREATE TABLE users (
    id_user SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    jabatan VARCHAR(50),
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL  -- Role menggunakan enum user_role
);
