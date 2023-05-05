CREATE TABLE users (
  id INT PRIMARY KEY,
  nama VARCHAR(50) NOT NULL,
  tgl_lahir DATETIME NOT NULL,
  alamat VARCHAR(50) NOT NULL,
  disabilitas VARCHAR(50) NOT NULL,
  kelamin CHAR(1) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(32) NOT NULL
);

CREATE TABLE perusahaans (
  id INT PRIMARY KEY,
  nama VARCHAR(50) NOT NULL,
  sektor VARCHAR(50) NOT NULL,
  alamat VARCHAR(50) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(32) NOT NULL
);

CREATE TABLE jobs (
  id INT PRIMARY KEY,
  deskripsi VARCHAR(50) NOT NULL,
  alamat VARCHAR(50) NOT NULL,
  expire DATETIME NOT NULL,
  status VARCHAR(50) NOT NULL,
  gaji INT NOT NULL,
  perusahaan_id INT NOT NULL,
  FOREIGN KEY (perusahaan_id) REFERENCES perusahaans(id)
);

CREATE TABLE lamarans (
  id INT PRIMARY KEY,
  lamaran_status VARCHAR(32) NOT NULL,
  user_id INT NOT NULL,
  job_id INT NOT NULL,
  perusahaan_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (job_id) REFERENCES jobs(id),
  FOREIGN KEY (perusahaan_id) REFERENCES perusahaans(id)
);

CREATE TABLE lampirans (
  id INT PRIMARY KEY,
  lampiran_tipe INT NOT NULL,
  lampiran_content TEXT NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE admins (
  id INT PRIMARY KEY,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(32) NOT NULL
);
