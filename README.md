# Demo_mogoDB
วิธีเขียน GO เชื่อมต่อกับ Database mogoDB
สามารถอ่านการอธิบายการทำงานของ Code ได้ที่

[สร้าง API GO เชื่อมต่อกับ MongoDB](https://bit.ly/2DzWRCj)

## How To Ues
```
docker pull mongo:3.6
docker run --name smalldogStore -p 27017:27017 -d mongo:3.6
go run main.go
```
สามารถใช้งานได้ที่ http://localhost:3000/api/v1/product
