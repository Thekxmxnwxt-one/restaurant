# 🍽️ ระบบร้านอาหาร (Restaurant System)

ระบบนี้เป็นส่วนหนึ่งของโปรเจกร้านอาหาร โดยเน้นในส่วนของ **ระบบ Login สำหรับตำแหน่ง admin**  
เมื่อเข้าสู่ระบบสำเร็จ ผู้ดูแลระบบจะถูกพาไปยังหน้า **จัดการเมนูอาหาร (/manage-menu)**

---

## 🛠️ เทคโนโลยีที่ใช้

- ✅ **Frontend**: Angular  
- ✅ **Backend**: Go (Gin)  
- ✅ **Database**: PostgreSQL  
- ✅ **Containerization**: Docker & Docker Compose

---

## 🌐 การเข้าถึงระบบ

📍 Frontend: [`http://localhost:4200`](http://localhost:4200)  
📍 Backend: [`http://localhost:8082`](http://localhost:8082)

#### 🔐 บัญชีตัวอย่างสำหรับทดลองใช้งาน:

| ตำแหน่ง | Username | Password |
|---------|----------|----------|
| admin   | admin    | pass123  |

> 🧭 เมื่อเข้าสู่ระบบสำเร็จ จะถูกนำไปที่หน้า `/management-menu`

---

## ⚙️ การเริ่มต้นใช้งาน

### 1. 🐳 รัน Backend ด้วย Docker

```bash
docker-compose up -d

###❗ หากมีปัญหาเรื่อง network ให้รันคำสั่งนี้ก่อน:

```bash
docker network create restaurant_network
docker-compose up -d

### 2. 🐳 รัน Frontend

```bash
ng serve

---