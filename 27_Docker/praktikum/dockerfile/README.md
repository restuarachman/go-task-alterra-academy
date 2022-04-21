# Membuat EC2 instance
1. Login ke console.aws.com
2. Pilih region: `ap-southeast-1` (Singapore), atau `ap-southeast-3` (Jakarta), atau `ap-southeast-2` (Sydney)
3. Akses EC2 Dashboard
4. Buat dan download keypair
    * format `.pem`
    * download
5. Launch EC2 instance
    * pilih yang free tier
    * pilih keypair yang sebelumnya sudah dibuat
6. Catat public IP dari instance EC2 yang sudah terbuat

# Melakukan koneksi ke EC2 instance via SSH

```bash
chmod 700 file-permission.pem

ssh -i file-permission.pem ubuntu@ip-ec2-instance
# ssh -i laptop.pem ubuntu@54.255.63.96

# Install golang
bash
# di dalam ec2
sudo apt-get install golang
go version

# Copy source code ke ec2

bash
# di local computer
scp -r -i laptop.pem helloworld ubuntu@54.255.63.96:/home/ubuntu
# rsync -r -e "ssh -i ~/laptop.pem" helloworld ubuntu@54.179.163.193:/home/ubuntu

# Buka port (outbound rule)
Ada dua macam rule pada security group:

* Inbound: Port untuk data masuk
* Outbound: Port untuk data keluar

Cara untuk atur inbound/outbound rule:

1. Klik instance ec2 (di web)
2. Klik tab security
3. Klik security group sg-<nama-security-group>
4. Klik Edit Inbound Rule

# Terminate EC2

1. Masuk ke dashboard EC2
2. Pilih running instance
3. Klik centang
4. Klik instance state | terminate