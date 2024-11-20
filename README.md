# Rombak-Projek-Magang_Go_Version

# ERD Version (1.0.0) :

## ENTITAS:

1. Bupati
   - ID Bupati (Primary Key)
   - Nama
   - Periode Jabatan
   - Created_at
   - Updated_at

2. Agenda
   - ID Agenda (Primary Key)
   - ID Bupati (Foreign Key)
   - Nama Agenda
   - Deskripsi
   - ID Lokasi (Foreign Key)
   - ID Jenis Acara (Foreign Key)
   - Tanggal Mulai
   - Tanggal Selesai

3. Lokasi
   - ID Lokasi (Primary Key)
   - Nama Lokasi
   - Alamat
   - Created_at
   - Updated_at

4. User
   - ID User (Primary Key)
   - Nama
   - Jabatan
   - username
   - email
   - password
   - role
   - is_email_verifed
   - created at
   - update at 
   - delete at

5. Jenis Acara
   - ID Jenis Acara (Primary Key)
   - Nama Jenis Acara
   - created at
   - update at 

RELASI:

1. Bupati dapat memiliki banyak agenda.
2. Setiap agenda berlangsung di satu lokasi, tetapi satu lokasi dapat digunakan untuk banyak agenda.
3. Setiap agenda memiliki satu jenis acara, tetapi satu jenis acara dapat digunakan untuk banyak agenda.

## Feature

### User Fitur 
1. Auth
2. Register
3. verif by email
4. management user
5. forget pass 
